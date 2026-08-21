# LeetCode (Go)

My [LeetCode](https://leetcode.com) solutions to problems I have solved in Go.

**Profile:** [NabhagMotivaras](https://leetcode.com/u/NabhagMotivaras/)

[![LeetCode Stats](https://leetcard.jacoblin.cool/NabhagMotivaras?ext=heatmap)](https://leetcode.com/u/NabhagMotivaras/)

Accepted submissions are exported with [leetcode-export](https://github.com/NeverMendel/leetcode-export) via Docker. Problem statements are not stored (LeetCode content is their IP).

## Layout

```text
0001-two-sum/
  123456789.go
  123456790.go
0042-trapping-rain-water/
  111222333.go
```

- One directory per problem at the repo root: `{id}-{title-slug}/`
- IDs are zero-padded to 4 digits so folders sort in problem order
- Every accepted Go submission is kept; filename is the LeetCode submission id

## Re-export

Requires Docker. Copy the `cookie` request header from a logged-in leetcode.com request in DevTools.

```bash
docker pull nevermendel/leetcode-export

docker run --rm \
  -v "$(pwd):/usr/app/out" \
  nevermendel/leetcode-export \
  --cookies "$LEETCODE_COOKIES" \
  --only-accepted \
  --language=golang \
  --no-problem-statement \
  --problem-folder-name='${question_id}-${title_slug}' \
  --submission-filename='${id}.go'
```

Then zero-pad folder IDs (`1-two-sum` → `0001-two-sum`):

```bash
python3 - <<'PY'
from pathlib import Path

for path in list(Path(".").iterdir()):
    if not path.is_dir() or "-" not in path.name:
        continue
    prefix, rest = path.name.split("-", 1)
    if not prefix.isdigit():
        continue
    dest = path.with_name(f"{int(prefix):04d}-{rest}")
    if dest == path:
        continue
    if dest.exists():
        raise SystemExit(f"Cannot rename {path} → {dest}: destination exists")
    path.rename(dest)
PY
```
