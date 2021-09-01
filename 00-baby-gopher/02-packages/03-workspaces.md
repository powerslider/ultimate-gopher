# Workspaces

- [Overview](#overview)
- [Creating a workspace](#creating-a-workspace)
- [Workspace Structure](#workspace-structure)

## Overview

- Your Go code is kept in a _workspace_.
- A workspace contains many source repositories.
- The `go` tool understands the layout of a workspace.
- You do not need a `Makefile`. The file layout is everything.
- Change the file layout, change the build.

```
$GOPATH/
    src/
        github.com/user/repo/
            mypkg/
                mysrc1.go
                mysrc2.go
            cmd/mycmd/
                main.go
    bin/
        mycmd
```

## Creating a workspace

- The traditional place for a workspace usually resides in `$HOME` or `$HOME/go`:

```bash
mkdir -p $HOME/go
```

- Set `GOPATH` environment variable to point to your workspace directory:

```bash
export GOPATH="$HOME/go"
```

---
__NOTE__
> The `GOPATH` environment variable tells the `go` tool where your workspace is located.
---

- The `go get` command fetches source repositories from the internet and places them in your workspace:

```bash
go get github.com/dsymonds/fixhub/cmd/fixhub
```

---
__NOTE__
> Package paths matter to the `go` tool. Using `github.com/...` means the tool knows how to fetch your repository.
---

- The `go install` command builds a binary and places it in `$GOPATH/bin/fixhub`:

```bash
go install github.com/dsymonds/fixhub/cmd/fixhub
```

## Workspace Structure

- After the execution of the previous `go get` and `go install` commands we got the following file structure:

```
$GOPATH/
    bin/fixhub                              # installed binary
    pkg/darwin_amd64/                       # compiled archives
        code.google.com/p/goauth2/oauth.a
        github.com/...
    src/                                    # source repositories
        code.google.com/p/goauth2/
            .hg
            oauth                           # used by package go-github
            ...
        github.com/
            golang/lint/...                 # used by package fixhub
                .git
            google/go-github/...            # used by package fixhub
                .git
            dsymonds/fixhub/
                .git
                client.go
                cmd/fixhub/fixhub.go        # package main
```

- `go get` fetched many repositories including the transitive dependencies of our current dependency.
- `go install` built a binary out of them.

[Next Section](04-package-organization.md)

[Previous Section](02-scope-and-visibility.md)

[Chapter Overview](README.md)