<h1>Dumber</h1>

`dumber`, a (not so dumb) command line tool for **d**(igital n)**umber**(ing) Markdown document sections and create a table of contents accordingly on demand.

With `dumber` (you already use `git`, don't you?) you can *add* or *remove* to your Mardown files:

- section numbers to header sections. Works on hash sign only, so one can exclude HTML section tags as &lt;H1&gt;.
- a table of contents with links on entries.

## 2. Installation

Install `dumber` from Github:

```
go install github.com/MichelBoucey/dumber@latest
```

## 1. Usage

```
user@machine $ dumber -h
Usage: dumber [OPTION] FILE

  -h    Show help
  -r    Remove table of contents and section numbers from the .md file
  -t    Add a table of contents to the .md file (can not be combined with -r)
  -v    Show version
  -w    Write section numbers to the .md file (default to stdout)
```

