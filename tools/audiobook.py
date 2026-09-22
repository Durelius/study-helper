"""Turn the lecture notes into a listenable audiobook.

Reading notes aloud verbatim does not work: a markdown table becomes a wall of pipes,
a slide reference becomes noise, and the inline `?check` markers vanish entirely. So
this rewrites the notes for the ear —

  * tables become sentences, using the header to label each cell
  * `[slides 5-6]` references are dropped; they mean nothing out loud
  * bold and italics are stripped, but the sentence keeps its shape
  * acronyms the course uses as letters (W B S, E S, L F) get spelled out
  * symbols become words: -> becomes "leads to", >= becomes "at least"
  * a `?check` marker becomes real retrieval practice: the question is read, then a
    silence long enough to answer it in your head, then the answer

The output is a list of (text, pause_after) segments, so pacing is controlled here
rather than left to punctuation.
"""
import json
import pathlib
import re
import sys

ROOT = pathlib.Path(__file__).resolve().parent.parent
LECTURES = ROOT / 'content' / 'lectures'
QUESTIONS = ROOT / 'content' / 'questions'

# The order the course teaches them in, which is also the order to listen in.
ORDER = [
    'value-chain', 'operations', 'process-mapping', 'scm', 'logistics',
    'pm-foundations', 'initiation', 'schedule-cost', 'risk-comm-quality',
]

# Said as letters, not as words. Anything pronounceable (RACI, SWOT, PMBOK, ISO,
# SCOR, FMEA) is deliberately left alone.
SPELLED = [
    'WBS', 'RBS', 'PBS', 'OBS', 'KPI', 'CPM', 'SCM', 'BPM', 'BPMN',
    'ES', 'EF', 'LS', 'LF', 'FS', 'SS', 'FF', 'SF', 'QA', 'QC', 'CoQ',
    'PV', 'EV', 'AC', 'CV', 'SV', 'CPI', 'SPI', 'EAC', 'BAC', 'ETC', 'VAC',
    'AON', 'RACI', 'XOR', 'AND', 'OR', 'PO', '3PL', '4PL', 'MHESI', 'ERP',
]
# Of those, the ones that really are words keep their pronunciation.
SPELLED = [a for a in SPELLED if a not in {'RACI', 'AND', 'OR'}]

SYMBOLS = [
    (r'→', ' leads to '), (r'←', ' comes from '), (r'↔', ' pairs with '),
    (r'×', ' by '), (r'÷', ' divided by '), (r'−', ' minus '), (r'–', ', '),
    (r'≥', ' at least '), (r'≤', ' at most '), (r'≠', ' is not '),
    (r'\bvs\.?\b', ' versus '), (r'&', ' and '), (r'฿', ' baht '),
    (r'%', ' percent'), (r'\.\.\.', '. '), (r'(?<=[\w\s])=(?=[\w\s])', ' equals '),
    (r'≈', ' about '), (r'…', '. '), (r'·', ', '), (r'○', ' '), (r'\|', ', '),
    (r'↑', ' upward '), (r'↓', ' downward '), (r'_', ' '), (r'\+', ' plus '),
    (r'<(?=\s*\d)', ' under '), (r'>(?=\s*\d)', ' over '),
    (r'<', ' less than '), (r'>', ' greater than '),
    (r'[\[\]*]', ''),
]


def spell(text: str) -> str:
    for acronym in sorted(SPELLED, key=len, reverse=True):
        spaced = ' '.join(acronym.upper())

        def replace(m, spaced=spaced):
            # A trailing plural joins the last letter, so KPIs reads "K P Is".
            return spaced + m.group(1)

        text = re.sub(rf'(?<![A-Za-z]){re.escape(acronym)}(s?)(?![A-Za-z])', replace, text)
    return text


def clean(text: str) -> str:
    """Strip markdown down to something a voice can read."""
    text = re.sub(r'`\[[^\]]*slides?[^\]]*\]`', '', text)   # slide references
    text = re.sub(r'\[[^\]]*slides?[^\]]*\]', '', text)
    text = re.sub(r'!\[[^\]]*\]\([^)]*\)', '', text)     # images
    text = re.sub(r'\[([^\]]+)\]\([^)]*\)', r'\1', text)  # links keep their words
    text = text.replace('**', '').replace('__', '')
    text = re.sub(r'(?<!\w)[*_](\S[^*_]*\S)[*_](?!\w)', r'\1', text)
    text = text.replace('`', '')
    for pattern, replacement in SYMBOLS:
        text = re.sub(pattern, replacement, text)
    text = spell(text)
    text = re.sub(r'[ \t]+', ' ', text).strip()
    # A dangling em dash reads as a stumble; make it a comma.
    text = text.replace(' — ', ', ').replace('—', ', ')
    return text


def split_long(text: str, limit: int = 55) -> list[str]:
    """Break a long paragraph on sentence boundaries, so no single utterance runs on."""
    text = text.strip()
    if not text:
        return []
    if len(text.split()) <= limit:
        return [text]
    sentences = re.split(r'(?<=[.?!])\s+', text)
    out, current = [], []
    for sentence in sentences:
        current.append(sentence)
        if len(' '.join(current).split()) >= limit:
            out.append(' '.join(current))
            current = []
    if current:
        out.append(' '.join(current))
    return out


def table_to_speech(rows: list[str]) -> list[str]:
    """A markdown table, said out loud one row at a time."""
    cells = [[c.strip() for c in r.strip().strip('|').split('|')] for r in rows]
    if len(cells) < 3:
        return []
    header = [clean(h) for h in cells[0]]
    out = []
    for row in cells[2:]:                                  # row 1 is the --- divider
        row = [clean(c) for c in row]
        if not any(row):
            continue
        lead = row[0].rstrip('.')
        parts = []
        for name, value in zip(header[1:], row[1:]):
            if not value or value in {'-', '—'}:
                continue
            parts.append(f'{name}: {value}' if name else value)
        out.append(f'{lead}. ' + '. '.join(parts) + '.' if parts else f'{lead}.')
    return out


def question_bank() -> dict:
    bank = {}
    for path in QUESTIONS.glob('*.json'):
        for q in json.loads(path.read_text()):
            bank[q['id']] = q
    return bank


def check_segments(q: dict) -> list[tuple[str, float]]:
    """A question, a silence to think in, then the answer and why."""
    stem = clean(re.sub(r'\s+', ' ', q['stem']))
    letters = 'A B C D E'.split()
    if q['type'] == 'tf':
        prompt = f'True or false. {stem}'
    else:
        options = '. '.join(f'{letters[i]}. {clean(c)}' for i, c in enumerate(q['choices']))
        prompt = f'{stem} {options}.'
    answer = ', '.join(f'{letters[i]}, {clean(q["choices"][i])}' for i in q['answer'])
    return [
        ('Question.', 0.35),
        (prompt, 3.2),                                     # long enough to actually answer
        (f'The answer is {answer}.', 0.5),
        (clean(q['explanation']), 1.1),
    ]


def lecture_segments(path: pathlib.Path, index: int, bank: dict) -> tuple[str, list]:
    raw = path.read_text()
    meta, body = {}, raw
    if raw.startswith('---'):
        head, body = raw[3:].split('\n---', 1)
        for line in head.splitlines():
            if ':' in line:
                k, v = line.split(':', 1)
                meta[k.strip()] = v.strip().strip('"')

    title = meta.get('title', path.stem)
    segments = [
        (f'Lecture {index}. {clean(title)}.', 1.0),
        (clean(meta.get('summary', '')), 1.4),
    ]

    lines = body.splitlines()
    i, paragraph, table = 0, [], []

    def flush_paragraph():
        if paragraph:
            text = clean(' '.join(paragraph))
            for part in split_long(text):
                segments.append((part, 0.55))
            paragraph.clear()

    def flush_table():
        if table:
            for row in table_to_speech(table):
                segments.append((row, 0.35))
            table.clear()

    while i < len(lines):
        line = lines[i].rstrip()
        stripped = line.strip()

        if stripped.startswith('|'):
            table.append(stripped)
            i += 1
            continue
        flush_table()

        marker = re.match(r'^\?check id=(\S+)$', stripped)
        if marker:
            flush_paragraph()
            q = bank.get(marker.group(1))
            if q:
                segments.extend(check_segments(q))
            i += 1
            continue

        if stripped.startswith('## '):
            flush_paragraph()
            segments.append((f'Section. {clean(stripped[3:])}.', 0.9))
            i += 1
            continue
        if stripped.startswith('### '):
            flush_paragraph()
            segments.append((clean(stripped[4:]) + '.', 0.7))
            i += 1
            continue

        if stripped.startswith('>'):
            flush_paragraph()
            quote = [stripped.lstrip('> ').strip()]
            while i + 1 < len(lines) and lines[i + 1].strip().startswith('>'):
                i += 1
                quote.append(lines[i].strip().lstrip('> ').strip())
            text = clean(' '.join(quote))
            # The exam-focus callouts already start with "Exam focus."
            segments.append((text, 1.0))
            i += 1
            continue

        if re.match(r'^[-*] |^\d+\. ', stripped):
            flush_paragraph()
            item = clean(re.sub(r'^([-*] |\d+\. )', '', stripped))
            if item:
                segments.append((item + ('' if item.endswith('.') else '.'), 0.3))
            i += 1
            continue

        if not stripped:
            flush_paragraph()
            i += 1
            continue

        paragraph.append(stripped)
        i += 1

    flush_paragraph()
    flush_table()
    segments = [(t, p) for t, p in segments if t.strip()]
    return title, segments


def main():
    bank = question_bank()
    out = {}
    for n, topic in enumerate(ORDER, start=1):
        path = LECTURES / f'{topic}.md'
        if not path.exists():
            continue
        title, segments = lecture_segments(path, n, bank)
        out[topic] = {'index': n, 'title': title, 'segments': segments}
    if len(sys.argv) > 1 and sys.argv[1] == '--dump':
        topic = sys.argv[2]
        for text, pause in out[topic]['segments']:
            print(f'[{pause:.2f}] {text}')
        return
    dest = pathlib.Path(sys.argv[1]) if len(sys.argv) > 1 else ROOT / 'tools' / 'script.json'
    dest.write_text(json.dumps(out, indent=1))
    words = sum(len(t.split()) for v in out.values() for t, _ in v['segments'])
    print(f'{len(out)} lectures, {sum(len(v["segments"]) for v in out.values())} segments, '
          f'{words} words (~{words / 150:.0f} minutes of speech)')


if __name__ == '__main__':
    main()
