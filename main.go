package main

import (
	"github.com/MikebangSfilya/diceRoller/diceroll"
	"github.com/MikebangSfilya/diceRoller/hanlders"
	"github.com/MikebangSfilya/diceRoller/parser"
	"github.com/MikebangSfilya/diceRoller/server"
)

func main() {
	dice := &diceroll.Dice{}

	parserManager := parser.NewParserManager()
	handler := hanlders.New(dice, *parserManager)
	srv := server.NewServer(handler)
	srv.Start()
}
