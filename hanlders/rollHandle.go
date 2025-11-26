package hanlders

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"

	"github.com/MikebangSfilya/diceRoller/diceroll"
	"github.com/MikebangSfilya/diceRoller/parser"
)

type Handlers struct {
	Dice  *diceroll.Dice
	Parse parser.ParserManager
}

func New(dice *diceroll.Dice, parse parser.ParserManager) *Handlers { // интерфейс
	return &Handlers{
		Dice:  dice,
		Parse: parse,
	}
}

func (h *Handlers) Roll(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")

	persons, err := h.Parse.Parse(contentType, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	personsRes := []parser.PersonResult{}

	for _, v := range persons {
		roll := h.Dice.Roll()
		d := <-roll
		sum := d + v.Wits + v.Dext
		personsRes = append(personsRes, parser.PersonResult{Name: v.Name, Sum: sum})
	}

	sort.Slice(personsRes, func(i, j int) bool {
		return personsRes[i].Sum < personsRes[j].Sum
	})

	if err := json.NewEncoder(w).Encode(personsRes); err != nil {
		panic(err)
	}

}
