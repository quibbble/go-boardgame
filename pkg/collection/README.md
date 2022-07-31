# Collection

[![GoDoc](https://godoc.org/github.com/quibbble/go-boardgame/pkg/collection?status.svg)](https://godoc.org/github.com/quibbble/go-boardgame/pkg/collection)

This collection package provides out of the box logic for things like decks, hands, etc. In the past there have been problems with creating repetitive code across different games, for example rewriting/copying deck logic, so this simple package has been created with [Go generics](https://go.dev/doc/tutorial/generics) to allow it to be used with any of your own custom types and preventing this unnecessary repetition.