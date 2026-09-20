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
python3 scripts/normalize_problem_folders.py
```

If both `1-two-sum/` and `0001-two-sum/` exist, the script matches them by the
numeric problem ID, moves new submission files into the existing padded folder,
and removes the unpadded folder. Matching by ID also handles a changed problem
slug. Identical files are safely skipped. If the same filename contains
different content, the script stops instead of overwriting either copy.

## Automatic sync

The [Sync LeetCode submissions](.github/workflows/sync-leetcode.yml) GitHub
Actions workflow checks for new accepted Go submissions every 15 minutes. If
the export changes the repository, the workflow updates an
`automation/leetcode-sync` branch and opens a pull request against the default
branch. After every export step succeeds, the workflow squash-merges its own
pull request and deletes the automation branch. The export is downloaded into
a temporary staging directory first, so duplicate or partial export folders
never appear in the repository checkout. It can also be started manually from
**Actions → Sync LeetCode
submissions → Run workflow**.

To enable it:

1. Log in to LeetCode and open your browser's developer tools.
2. In the Network tab, select a request to `leetcode.com` and copy its complete
   `cookie` request header value.
3. In this GitHub repository, open **Settings → Secrets and variables → Actions**.
4. Create a repository secret named `LEETCODE_COOKIES` and paste the cookie
   value into it.
5. Under **Settings → Actions → General → Workflow permissions**, allow
   **Read and write permissions** and enable **Allow GitHub Actions to create
   and approve pull requests**. The workflow creates and merges only its
   `automation/leetcode-sync` pull request.
6. Open the workflow in the Actions tab and run it once manually to verify the
   secret.

LeetCode login cookies expire. If the workflow starts failing authentication,
replace the `LEETCODE_COOKIES` secret with a fresh cookie header. Keep this
value only in GitHub Actions secrets; never commit it to the repository.
