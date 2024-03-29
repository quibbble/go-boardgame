# Go-boardgame

[![GoDoc](https://godoc.org/github.com/quibbble/go-boardgame?status.svg)](https://godoc.org/github.com/quibbble/go-boardgame)

Go-boardgame is a simple [Go](https://golang.org) package that can be used as the scaffolding to write the game logic for any board or turn based game.

## Status

This package is now stable with new features in the works. 

## Installation

```
go get github.com/quibbble/go-boardgame
```

## Packages

Go-boardgame also contains a number of helpful packages that can be optionally used to speed up development:

- [bgerr](https://github.com/quibbble/go-boardgame/tree/main/pkg/bgerr) adds helpful error statuses and messages to board games.
- [bgn](https://github.com/quibbble/go-boardgame/tree/main/pkg/bgn) adds a standardized machine-readable notation to board games for easy storage and game recreation.
- [collection](https://github.com/quibbble/go-boardgame/tree/main/pkg/collection) adds logic that can be used for decks, hands, etc. This package makes use of [Go generics](https://go.dev/doc/tutorial/generics) so that it can be used with any of your custom types.

## Examples

[Tic-Tac-Toe](https://github.com/quibbble/go-tictactoe) provides a simple example implementation.

You can also view [all currently implemented games here](https://github.com/quibbble#currently-implemented-games).

## Future Plans

Additional resources common to many games will be added to the above packages list as time goes on to make writing the logic for these games a far faster process. Any ideas or PRs to improve or add additional features are always welcome.
