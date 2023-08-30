<h1>Dumber</h1>

`dumber`, a (not so dumb) command line tool for **d**(igital n)**umber**(ing) Markdown document sections and create a table of contents accordingly on demand.

With `dumber` (you already use `git`, don't you?) you can *add* or *remove* to your Mardown files:

- section numbers to header sections. This works on hash sign only, so one can exclude HTML section tags, as &lt;H1&gt;, to stay unnumenbered.
- a table of contents with links on entries.

See an [example](./example.md).

## 1. Installation

Install `dumber` from Github:

```
go install github.com/MichelBoucey/dumber@latest
```

## 2. Usage

```
user@machine $ dumber -h
Usage: dumber [OPTION] FILE

  -h    Show help
  -r    Remove section numbers and table of contents from the .md file
  -v    Show version
  -w    Write section numbers to the .md file (default to stdout)
```

To add a table of contents you have to add a line with the HTML comment below, where you want the table of contents to appear:

```
<!-- ToC -->
```

The table of contents will be written just after the HTML comment line. And you can add this comment line as many times as you want, if you are, like me, a big fan of tables of contents, or if the length of your document needs a second table of contents at its end.

