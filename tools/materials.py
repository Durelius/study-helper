"""Render a course's source decks and infographics into images the app can show.

    python3 tools/materials.py courses/genai-literacy

Reads <course>/materials.json, which maps each deck name the content cites (the
`deck` in a question's source, or the words before "slide" in a lecture reference) to
the PDF or zipped images it came from. Writes one WebP per page into
<course>/materials/<slug>/ and a manifest.json the app reads to make every citation
clickable. Like the audiobook, the output is not in git; deploy.sh ships it with the
rest of the course directory.

Needs pdftoppm (poppler) and cwebp.
"""
import json
import pathlib
import re
import shutil
import subprocess
import sys
import tempfile
import zipfile

WIDTH = 1600     # slides are read on a laptop, full width
QUALITY = 80

course = pathlib.Path(sys.argv[1])
spec = json.loads((course / 'materials.json').read_text())
root = pathlib.Path(spec['root']).expanduser()
out = course / 'materials'
shutil.rmtree(out, ignore_errors=True)
out.mkdir()


def slug(name):
    return re.sub(r'[^a-z0-9]+', '-', name.lower()).strip('-')


def webp(src, dest, width=None):
    args = ['cwebp', '-quiet', '-q', str(QUALITY)]
    if width:
        args += ['-resize', str(width), '0']
    subprocess.run(args + [str(src), '-o', str(dest)], check=True)


manifest = {}
with tempfile.TemporaryDirectory() as tmp:
    tmp = pathlib.Path(tmp)
    for d in spec['decks']:
        name = slug(d['deck'])
        (out / name).mkdir()
        pages = {}
        if 'pdf' in d:
            png = tmp / name
            png.mkdir()
            subprocess.run(['pdftoppm', '-png', '-scale-to-x', str(WIDTH), '-scale-to-y', '-1',
                            str(root / d['pdf']), str(png / 'p')], check=True)
            for p in sorted(png.glob('p-*.png')):
                n = int(p.stem.split('-')[1])
                webp(p, out / name / f'{n}.webp')
                pages[str(n)] = [f'{name}/{n}.webp']
        else:
            # Infographics are numbered by the sheet table in exam-focus.md §11, not by
            # their order in the zip, so each sheet names its file(s) explicitly.
            with zipfile.ZipFile(root / d['zip']) as z:
                members = z.namelist()
                for n, files in d['pages'].items():
                    pages[n] = []
                    for i, f in enumerate(files):
                        member = next(m for m in members if m.endswith(f))
                        raw = tmp / f'{name}-{n}-{i}{pathlib.Path(f).suffix}'
                        raw.write_bytes(z.read(member))
                        rel = f'{name}/{n}' + (f'-{i + 1}' if len(files) > 1 else '') + '.webp'
                        webp(raw, out / rel)
                        pages[n].append(rel)
        manifest[d['deck']] = {'title': d['title'], 'label': d.get('label', 'slide'), 'pages': pages}
        print(f'{d["deck"]}: {len(pages)} pages')

(out / 'manifest.json').write_text(json.dumps(manifest, indent=1))
size = sum(p.stat().st_size for p in out.rglob('*.webp'))
print(f'{size / 1e6:.1f} MB in {out}')
