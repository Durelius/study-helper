"""Speak the audiobook script with Kokoro.

Run with the venv's python, which has kokoro and soundfile:

    ESPEAK_DATA_PATH=/opt/homebrew/share/espeak-ng-data \
      venv/bin/python tools/synth.py script.json out/ [--limit N]

Each lecture becomes one .m4a with track metadata, so a phone shows them in order.
"""
import json
import os
import pathlib
import subprocess
import sys
import time

import numpy as np
import soundfile as sf
from kokoro import KPipeline

SR = 24000
VOICE = os.environ.get('VOICE', 'am_michael')
ALBUM = 'Value Chain Excellence & Project Management — midterm'


def main():
    script = json.loads(pathlib.Path(sys.argv[1]).read_text())
    dest = pathlib.Path(sys.argv[2])
    dest.mkdir(parents=True, exist_ok=True)
    limit = int(sys.argv[sys.argv.index('--limit') + 1]) if '--limit' in sys.argv else None
    only = sys.argv[sys.argv.index('--only') + 1] if '--only' in sys.argv else None

    pipe = KPipeline(lang_code='a')
    total_audio = 0.0
    started = time.time()

    for topic, lecture in sorted(script.items(), key=lambda kv: kv[1]['index']):
        if only and topic != only:
            continue
        segments = lecture['segments'][:limit] if limit else lecture['segments']
        pieces = []
        for n, (text, pause) in enumerate(segments, 1):
            chunks = [a for _, _, a in pipe(text, voice=VOICE, speed=1.0)]
            if chunks:
                pieces.append(np.concatenate(chunks))
            if pause:
                pieces.append(np.zeros(int(SR * pause), dtype=np.float32))
            if n % 25 == 0 or n == len(segments):
                print(f'  {topic} {n}/{len(segments)}', flush=True)

        audio = np.concatenate(pieces) if pieces else np.zeros(1, dtype=np.float32)
        seconds = len(audio) / SR
        total_audio += seconds

        index = lecture['index']
        wav = dest / f'{index:02d}.wav'
        sf.write(wav, audio, SR)
        m4a = dest / f'{index:02d} {lecture["title"].replace("/", "-")}.m4a'
        subprocess.run([
            'ffmpeg', '-v', 'error', '-y', '-i', str(wav),
            '-c:a', 'aac', '-b:a', '96k',
            '-metadata', f'title={lecture["title"]}',
            '-metadata', f'album={ALBUM}',
            '-metadata', 'artist=Lecture notes',
            '-metadata', f'track={index}/9',
            str(m4a)], check=True)
        wav.unlink()
        print(f'{index}. {lecture["title"]} -> {seconds / 60:.1f} min '
              f'({(time.time() - started) / 60:.1f} min elapsed)', flush=True)

    print(f'DONE: {total_audio / 60:.1f} minutes of audio in '
          f'{(time.time() - started) / 60:.1f} minutes')


if __name__ == '__main__':
    main()
