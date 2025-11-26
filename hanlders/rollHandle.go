package hanlders

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"

	cache "github.com/MikebangSfilya/diceRoller/Cache"
	"github.com/MikebangSfilya/diceRoller/diceroll"
	"github.com/MikebangSfilya/diceRoller/parser"
)

type Handlers struct {
	Dice  *diceroll.Dice
	Parse parser.ParserManager
	Cache *cache.InitiativeCache
}

func New(dice *diceroll.Dice, parse parser.ParserManager) *Handlers {
	return &Handlers{
		Dice:  dice,
		Parse: parse,
		Cache: cache.NewInitiativeCache(),
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

	CacheData := []cache.CacheRols{}
	personsRes := []parser.PersonResult{}

	for _, v := range persons {
		roll := h.Dice.Roll()
		dextPlusWits := v.Dext + v.Wits

		CacheData = append(CacheData, cache.New(v.Name, dextPlusWits))

		d := <-roll

		sum := d + dextPlusWits
		personsRes = append(personsRes, parser.PersonResult{Name: v.Name, Sum: sum})
	}

	h.Cache.Set(CacheData)

	sort.Slice(personsRes, func(i, j int) bool {
		return personsRes[i].Sum < personsRes[j].Sum
	})

	if err := json.NewEncoder(w).Encode(personsRes); err != nil {
		panic(err)
	}

}
