"""Rebuild manifest.json from an existing audiobook directory.

synth.py writes the manifest itself; this exists for a directory rendered before it
did, or one assembled by hand.
"""
import json
import pathlib
import re
import subprocess
import sys

dest = pathlib.Path(sys.argv[1])
script = json.loads(pathlib.Path(sys.argv[2]).read_text()) if len(sys.argv) > 2 else {}
by_index = {v['index']: k for k, v in script.items()}

out = []
for path in sorted(dest.glob('*.m4a')):
    m = re.match(r'^(\d+)\s+(.*)\.m4a$', path.name)
    if not m:
        continue
    index = int(m.group(1))
    seconds = float(subprocess.run(
        ['ffprobe', '-v', 'error', '-show_entries', 'format=duration', '-of', 'csv=p=0', str(path)],
        capture_output=True, text=True, check=True).stdout.strip())
    out.append({
        'topic': by_index.get(index, f'lecture-{index}'),
        'index': index,
        'title': m.group(2),
        'file': path.name,
        'seconds': round(seconds, 1),
        'bytes': path.stat().st_size,
    })

out.sort(key=lambda e: e['index'])
(dest / 'manifest.json').write_text(json.dumps(out, indent=1))
print(f'{len(out)} episodes, {sum(e["seconds"] for e in out) / 60:.0f} minutes')
