# go-presentx

![presentx screenshot](./_assets/screen1.png)

Go's [present tool](https://github.com/golang/tools/tree/master/cmd/present) but with code syntax highlighting. Highlighted code is editable with [some caveats](#about-the-syntax-highlighting-changes)

This is a quick and dirty implementation of [flippeeer's reddit post](https://www.reddit.com/r/golang/comments/jpugtg/today_i_presented_go_to_my_team_the_screen/).

### Requirements and instructions

* Requires Go installed. Enable modules with `go env -w GO111MODULE="on"`

Start presentation on [127.0.0.1:3999](http://127.0.0.1:3999/) (default) by running:

```console
go run . -base .
```

## Installation

The following commands will generate an executable tool for your OS and architecture. Requires Go
```console
git clone https://github.com/soypat/go-presentx.git
cd go-presentx
go build
```

## About the syntax highlighting changes
* No auto-rendering while editing

* Uses the lightweight [prism.js](https://prismjs.com/) for syntax highlighting

* Change code style by replacing `static/prism.css` with one from prism.js' site and modifying it to fit your needs (see [css config section](#prism-css-config))

* Press <kbd>CTRL</kbd> while in editable code block to re-render all syntax

* Issue with the caret being invisible at times and sometimes appearing to be on the next line
## Prism CSS config 

* Ligatures off by default. You may enable them in `prism-github-dark.css` by setting `"calt"` to `on` (see comment).

Non-default settings listed:

```css
pre[class*="language-"] {
    caret-color: #ccc;
    font-size: 0.95em;
    border-radius: 7px;
    line-height: 1.2;
}
pre[class*="language-"] {
    padding: 0.6em;
}
```

##  Contributing, License and authors

PRs welcome. I'll look at issues, I'll try to be nice.

idk, [same license as Go](https://github.com/golang/go/blob/master/LICENSE) I guess. Credit to **[The Go Authors.](https://github.com/golang/go/blob/master/AUTHORS)**
