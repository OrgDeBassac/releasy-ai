Design: `release-cli create-rel`

CLI behavior
- Command: `release-cli create-rel [--from-tag <tag>] [--from-rel <REL-x.y>] [--push] [--force]`
- If both `--from-tag` and `--from-rel` are provided, prefer `--from-tag` with a warning.
- If `--from-rel` is provided, find the latest tag for that virtual group (using existing tag listing logic) and use it as base.
- Validate tag exists: `git rev-list -n 1 <tag>`
- Branch name: `REL-<MAJOR>.<MINOR>` derived from tag
- Creation steps:
  - If branch already exists locally: fail unless `--force` (then delete & recreate)
  - If branch exists on origin: fail unless `--force` (then force push)
  - Create local branch at tag commit: `git checkout -b REL-X.Y <tag>`
  - If `--push` set: `git push origin REL-X.Y` (use --force if requested)

Error handling & messages
- Tag missing: "Tag <tag> not found"
- REL exists: "Branch REL-X.Y already exists. Use --force to recreate."
- Push failure: propagate error with guidance

Testing
- Unit tests for parsing and deriving REL name
- Integration test: create an annotated tag in a temporary repo and run command with --push to a fork (manual)

