# go-presentx
golang's present tool but with code syntax highlighting

This is a quick and dirty implementation of [flippeeers reddit post](https://www.reddit.com/r/golang/comments/jpugtg/today_i_presented_go_to_my_team_the_screen/).

Requires Go installed.

Start playground by running

```
go run . -base .
```

PRs welcome.

### Prism CSS config
Non-default settings listed:

```css
pre[class*="language-"] {
    font-size: 0.9em;
    border-radius: 5px;
    line-height: 1.2;
}
pre[class*="language-"] {
    padding: 0.6em;
}
```