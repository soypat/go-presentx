# go-presentx
golang's [present tool](https://github.com/golang/tools/tree/master/cmd/present) but with code syntax highlighting

This is a quick and dirty implementation of [flippeeer's reddit post](https://www.reddit.com/r/golang/comments/jpugtg/today_i_presented_go_to_my_team_the_screen/).

Requires Go installed. Enable modules with `go env -w GO111MODULE="on"`

Start presentation on [127.0.0.1:3999](http://127.0.0.1:3999/) (default) by running:

```console
go run . -base .
```

PRs welcome.

## About the playground
* Uses the lightweight [prism.js](https://prismjs.com/) for syntax highlighting
* Change code style by downloading `prism.css` from prism.js' site and modifying it to fit your needs (see [css config section](#prism-css-config))
* Press <kbd>CTRL</kbd> while in editable code block to re-render all syntax

## Prism CSS config 

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