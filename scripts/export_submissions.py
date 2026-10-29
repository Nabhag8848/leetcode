#!/usr/bin/env python3
"""Export accepted LeetCode Go submissions through the GraphQL API.

The REST endpoint used by leetcode-export
(https://leetcode.com/api/submissions/) now returns a Cloudflare challenge
page, which is not JSON. The submissions UI itself reads GraphQL, and that
endpoint still returns JSON.
"""

import argparse
import json
import os
import sys
import time
import urllib.error
import urllib.request
from pathlib import Path

GRAPHQL_URL = "https://leetcode.com/graphql"
USER_AGENT = (
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "
    "AppleWebKit/537.36 (KHTML, like Gecko) "
    "Chrome/131.0.0.0 Safari/537.36"
)
PAGE_SIZE = 20

STATUS_QUERY = """
query {
  userStatus {
    isSignedIn
    username
  }
}
"""

LIST_QUERY = """
query submissionList($offset: Int!, $limit: Int!, $lastKey: String, $questionSlug: String) {
  submissionList(
    offset: $offset
    limit: $limit
    lastKey: $lastKey
    questionSlug: $questionSlug
  ) {
    lastKey
    hasNext
    submissions {
      id
      titleSlug
      statusDisplay
      lang
      langName
    }
  }
}
"""

DETAIL_QUERY = """
query submissionDetails($submissionId: Int!) {
  submissionDetails(submissionId: $submissionId) {
    code
    question {
      questionId
      questionFrontendId
      titleSlug
    }
  }
}
"""


def parse_cookie_header(raw: str) -> dict[str, str]:
    text = "".join(raw.replace("\r", "").split("\n")).strip().strip('"').strip("'")
    if text.lower().startswith("cookie:"):
        text = text.split(":", 1)[1].strip()

    cookies: dict[str, str] = {}
    for part in text.split(";"):
        if "=" not in part:
            continue
        name, value = part.split("=", 1)
        name = name.strip()
        value = value.strip()
        if name:
            cookies[name] = value
    return cookies


def existing_submission_ids(root: Path) -> set[str]:
    found: set[str] = set()
    if not root.is_dir():
        return found
    for problem_dir in root.iterdir():
        if not problem_dir.is_dir() or problem_dir.name.startswith("."):
            continue
        for submission in problem_dir.glob("*.go"):
            if submission.stem.isdigit():
                found.add(submission.stem)
    return found


def is_accepted_go(submission: dict) -> bool:
    status = (submission.get("statusDisplay") or "").lower()
    lang = (submission.get("lang") or "").lower()
    lang_name = (submission.get("langName") or "").lower()
    return status == "accepted" and (lang in {"golang", "go"} or lang_name == "go")


class GraphQLClient:
    def __init__(self, cookies: dict[str, str]) -> None:
        self.headers = {
            "Accept": "application/json",
            "Content-Type": "application/json",
            "Origin": "https://leetcode.com",
            "Referer": "https://leetcode.com/submissions/",
            "User-Agent": USER_AGENT,
            "Cookie": "; ".join(f"{name}={value}" for name, value in cookies.items()),
            "x-csrftoken": cookies["csrftoken"],
        }
        try:
            from curl_cffi import requests as curl_requests

            self._curl = curl_requests
            print("Using a Chrome TLS fingerprint for LeetCode.", flush=True)
        except ImportError:
            self._curl = None
            print("Using Python urllib for LeetCode.", flush=True)

    def call(self, query: str, variables: dict | None = None) -> dict:
        payload = {"query": query, "variables": variables or {}}
        delay = 1
        last_error: Exception | None = None

        for attempt in range(6):
            try:
                status, text = self._post(payload)
            except Exception as error:
                last_error = error
                time.sleep(delay)
                delay = min(delay * 2, 30)
                continue

            if text.lstrip().startswith("<"):
                raise SystemExit(
                    "LeetCode returned a Cloudflare challenge instead of JSON "
                    f"(HTTP {status})."
                )

            try:
                body = json.loads(text)
            except json.JSONDecodeError:
                last_error = RuntimeError(
                    f"response was not JSON (HTTP {status}): {text[:180]!r}"
                )
                time.sleep(delay)
                delay = min(delay * 2, 30)
                continue

            if body.get("errors") and not body.get("data"):
                message = "; ".join(
                    error.get("message", "unknown error") for error in body["errors"]
                )
                last_error = RuntimeError(message)
                time.sleep(delay)
                delay = min(delay * 2, 30)
                continue

            return body

        raise SystemExit(f"LeetCode request failed: {last_error}")

    def _post(self, payload: dict) -> tuple[int, str]:
        if self._curl is not None:
            response = self._curl.post(
                GRAPHQL_URL,
                json=payload,
                headers=self.headers,
                impersonate="chrome",
                timeout=30,
            )
            return response.status_code, response.text

        request = urllib.request.Request(
            GRAPHQL_URL,
            data=json.dumps(payload).encode(),
            headers=self.headers,
        )
        try:
            with urllib.request.urlopen(request, timeout=30) as response:
                return response.status, response.read().decode()
        except urllib.error.HTTPError as error:
            return error.code, error.read().decode()


def submission_page(client: GraphQLClient, offset: int, last_key: str | None) -> dict:
    delay = 1
    for _ in range(6):
        payload = client.call(
            LIST_QUERY,
            {
                "offset": offset,
                "limit": PAGE_SIZE,
                "lastKey": last_key,
                "questionSlug": None,
            },
        )
        page = (payload.get("data") or {}).get("submissionList")
        if page and page.get("submissions") is not None:
            return page
        time.sleep(delay)
        delay = min(delay * 2, 30)

    raise SystemExit(
        "LeetCode did not return a submission list. "
        "If this persists, replace LEETCODE_COOKIES with a fresh cookie header from a logged-in leetcode.com request."
    )


def submission_code(client: GraphQLClient, submission_id: str) -> dict:
    delay = 1
    for _ in range(6):
        payload = client.call(DETAIL_QUERY, {"submissionId": int(submission_id)})
        details = (payload.get("data") or {}).get("submissionDetails")
        if details and details.get("code") and details.get("question"):
            return details
        time.sleep(delay)
        delay = min(delay * 2, 30)

    raise SystemExit(
        f"LeetCode returned no code for submission {submission_id}. "
        "This usually means the account is rate limited; rerun the export later."
    )


def problem_number(question: dict) -> int:
    raw = question.get("questionFrontendId") or question.get("questionId")
    if raw is None or not str(raw).isdigit():
        raise SystemExit(f"Submission is missing a numeric problem id: {question}")
    return int(raw)


def write_submission(output: Path, submission_id: str, details: dict) -> Path:
    question = details["question"]
    slug = question.get("titleSlug")
    if not slug:
        raise SystemExit(f"Submission {submission_id} is missing a title slug.")

    code = details["code"]
    if not code.endswith("\n"):
        code += "\n"

    folder = output / f"{problem_number(question)}-{slug}"
    folder.mkdir(parents=True, exist_ok=True)
    path = folder / f"{submission_id}.go"
    path.write_text(code, encoding="utf-8")
    return path


def export(output: Path, existing_root: Path | None, cookies: dict[str, str]) -> None:
    client = GraphQLClient(cookies)
    status = (client.call(STATUS_QUERY).get("data") or {}).get("userStatus") or {}
    if not status.get("isSignedIn"):
        raise SystemExit(
            "LeetCode did not accept these cookies. "
            "Replace LEETCODE_COOKIES with a fresh cookie header from a logged-in leetcode.com request."
        )

    username = status.get("username")
    if username:
        print(f"Signed in as {username}.", flush=True)

    known = existing_submission_ids(existing_root) if existing_root else set()
    seen: set[str] = set()
    offset = 0
    last_key = None
    written = 0
    skipped = 0

    for _ in range(500):
        page = submission_page(client, offset, last_key)
        submissions = page["submissions"]
        fresh = []
        for submission in submissions:
            submission_id = str(submission.get("id"))
            if not submission_id.isdigit() or submission_id in seen:
                continue
            seen.add(submission_id)
            fresh.append(submission)

        if not fresh:
            break

        for submission in fresh:
            submission_id = str(submission["id"])
            if not is_accepted_go(submission):
                continue
            if submission_id in known:
                skipped += 1
                continue

            details = submission_code(client, submission_id)
            path = write_submission(output, submission_id, details)
            written += 1
            print(f"Saved {path.relative_to(output)}", flush=True)
            time.sleep(0.4)

        if page.get("hasNext") is False:
            break
        if page.get("hasNext") is None and len(submissions) < PAGE_SIZE:
            break

        offset += len(submissions)
        last_key = page.get("lastKey")
        time.sleep(0.3)

    print(
        f"Wrote {written} new accepted Go submission(s); "
        f"skipped {skipped} already in the repo.",
        flush=True,
    )


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Export accepted LeetCode Go submissions into problem folders."
    )
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument(
        "--existing",
        type=Path,
        help="repository checkout; submission ids already present here are skipped",
    )
    parser.add_argument(
        "--cookies",
        help="cookie request header (defaults to the LEETCODE_COOKIES environment variable)",
    )
    args = parser.parse_args()

    raw_cookies = args.cookies if args.cookies is not None else os.environ.get("LEETCODE_COOKIES", "")
    cookies = parse_cookie_header(raw_cookies)
    missing = [name for name in ("csrftoken", "LEETCODE_SESSION") if name not in cookies]
    if missing:
        joined = " and ".join(missing)
        raise SystemExit(
            f"Cookie header must include {joined}. "
            "Copy the entire cookie request header from a logged-in leetcode.com request."
        )

    output = args.output.resolve()
    output.mkdir(parents=True, exist_ok=True)
    existing = args.existing.resolve() if args.existing else None
    export(output, existing, cookies)


if __name__ == "__main__":
    try:
        main()
    except BrokenPipeError:
        sys.exit(0)
