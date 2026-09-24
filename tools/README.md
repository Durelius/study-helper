# tools

## audiobook.py + synth.py — the lecture notes, read aloud

`audiobook.py` rewrites `content/lectures/*.md` for the ear and `synth.py` speaks the
result with [Kokoro](https://github.com/hexgrad/kokoro), a local neural TTS model. No
account, no API, no per-character cost — it runs on the laptop.

Reading the notes verbatim does not work, so the rewrite:

- turns every markdown table into spoken sentences, labelled by the table's own header
- drops the `[slides 5-6]` references, which are noise out loud
- spells out the acronyms the course says as letters (`WBS` → "W B S", `KPIs` → "K P Is")
- says the symbols: `→` "leads to", `≥` "at least", `×` "by", `−` "minus"
- turns each inline `?check` marker into real retrieval practice — the question is read,
  then a 3.2-second silence to answer it in your head, then the answer and why

Pacing is set per segment rather than left to punctuation, which is what makes the
quiz pauses long enough to actually think in.

### Running it

```sh
python3.12 -m venv venv && ./venv/bin/pip install kokoro soundfile
brew install espeak-ng ffmpeg

python3.12 tools/audiobook.py script.json
ESPEAK_DATA_PATH=/opt/homebrew/share/espeak-ng-data \
  ./venv/bin/python tools/synth.py script.json ~/Desktop/valuechain-audiobook
```

`ESPEAK_DATA_PATH` is required: the phonemiser otherwise looks for its data at a path
baked in on someone else's build machine.

`VOICE` picks the narrator (default `am_michael`). `--only <topic>` and `--limit N`
render a short sample instead of all nine lectures.

Output is one `.m4a` per lecture with track metadata, so a phone lists them in order.
Roughly 150 minutes in total, and synthesis runs about 8x faster than real time.

Episodes appear one lecture at a time, so `deploy/sync-audio.sh` ships whatever has
rendered so far without rebuilding the app: it rewrites the manifest from what is on
disk, rsyncs only the new files and restarts the service, which reads the manifest at
startup.

To check a voice before committing to a full run, `audiobook.py --dump <topic>` prints
the speech text with its pauses, which is also the fastest way to catch a symbol that
would read badly.

## materials.py — the slides behind every citation

Renders each deck a course cites into one WebP per page, so a citation in the app opens
the slide it came from. `<course>/materials.json` maps each deck name the content uses
(`lec 1`, `GAI Lit 07`, `infographics`) to its PDF, or for zipped images, to the file
behind each sheet number.

```sh
brew install poppler webp
python3 tools/materials.py courses/genai-literacy
```

Output goes to `<course>/materials/` (about 19 MB for GenAI Literacy). It is not in git;
`deploy.sh` ships it with the rest of the course directory. A course without it shows
citations as plain text.
