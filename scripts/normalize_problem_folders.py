#!/usr/bin/env python3
"""Zero-pad LeetCode problem folders and merge duplicate folder variants."""

import argparse
import filecmp
import shutil
from pathlib import Path


def normalize(source_root: Path, destination_root: Path) -> None:
    same_root = source_root == destination_root

    for source in list(source_root.iterdir()):
        if not source.is_dir() or "-" not in source.name:
            continue

        prefix, slug = source.name.split("-", 1)
        if not prefix.isdigit():
            continue

        problem_id = int(prefix)
        padded_prefix = f"{problem_id:04d}-"

        # An existing folder is authoritative for its numeric problem ID. This
        # also handles a LeetCode slug change without creating a second folder.
        existing = [
            path
            for path in destination_root.iterdir()
            if path.is_dir()
            and (not same_root or path != source)
            and path.name.startswith(padded_prefix)
        ]
        if len(existing) > 1:
            names = ", ".join(sorted(path.name for path in existing))
            raise SystemExit(
                f"Multiple padded folders exist for problem {problem_id}: {names}"
            )

        destination = (
            existing[0]
            if existing
            else destination_root / f"{problem_id:04d}-{slug}"
        )
        if destination == source:
            continue

        if not destination.exists():
            shutil.move(str(source), destination)
            print(f"Imported {source.name} -> {destination.name}")
            continue

        for submission in source.iterdir():
            if not submission.is_file():
                raise SystemExit(
                    f"Cannot merge nested directory: {submission.relative_to(source_root)}"
                )

            target = destination / submission.name
            if not target.exists():
                shutil.move(str(submission), target)
                print(
                    f"Moved {submission.relative_to(source_root)} -> "
                    f"{target.relative_to(destination_root)}"
                )
            elif filecmp.cmp(submission, target, shallow=False):
                submission.unlink()
                print(
                    "Skipped identical file: "
                    f"{target.relative_to(destination_root)}"
                )
            else:
                raise SystemExit(
                    "Refusing to overwrite different files with the same name: "
                    f"{submission.relative_to(source_root)} and "
                    f"{target.relative_to(destination_root)}"
                )

        source.rmdir()
        print(f"Removed merged folder: {source.name}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Zero-pad LeetCode problem IDs and merge duplicate folders."
    )
    parser.add_argument(
        "source",
        nargs="?",
        type=Path,
        default=Path("."),
        help="export directory to normalize (default: current directory)",
    )
    parser.add_argument(
        "--destination",
        type=Path,
        help="repository directory to merge into (default: same as source)",
    )
    args = parser.parse_args()
    source = args.source.resolve()
    destination = (args.destination or source).resolve()
    normalize(source, destination)
