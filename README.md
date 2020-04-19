# breadth-walker

This is a tool to print directories walked in a BFS style from current location to standard out. 
I use it in conjunction with [fzf](https://github.com/junegunn/fzf) to make a quick fuzzy directory navigator.

```
alias f='breadthwalker 5 | fzf | xargs -I @ sh -c '[ -d @ ] && cd @'
```

## installation

`go get https://github.com/tomatosource/breadth-walker`

## usage

`breadthwalker 3 # where 3 is optional max depth, default is 50`

## TODOs

- use proper flag system
- ignores to flag rather than hardcoded

