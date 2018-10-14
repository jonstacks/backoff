# README

CLI commands for running a command that could fail multiple times with
different backoffs and number of retries.

## Install

To install run:

```
go get -u github.com/jonstacks/backoff/cmd/...
```

## Commands

### constant-backoff

Constant backoff has the following usage:

```
constant-backoff -retries=3 -wait=5 cmd args...
```

### exponential-backoff

Exponential backoff has the following usage:

```
exponential-backoff -retries=3 cmd args...
```

## Notes

Both commands call the user supplied `cmd` with any given `args`. If the command
exits with a 0 status code, both backoffs exit with a 0 as well. If the user
supplied command exists with a non-zero status code, then it will be retried
until it does return a 0 status code or we run out of available retries.
