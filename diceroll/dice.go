package diceroll

import (
	"math/rand"
	"sync"
)

type Dice struct {
	DiceRes int8
	mu      sync.RWMutex
}

// Roll rolls a D10 dice, where n is the number of dice rolled.
func (d *Dice) Roll() <-chan int8 {
	out := make(chan int8)

	go func() {
		out <- int8(rand.Intn(10)) + 1
		close(out)
	}()

	return out
}

// func (d *Dice) Accumulate(in <-chan int8) []int8 {
// 	res := make([]int8, 0)
// 	for v := range in {
// 		res = append(res, v)
// 	}
// 	return res
// }

// func (d *Dice) RollPul(n int) []int8 {
// 	res := []int8{}

// 	var wg sync.WaitGroup

// 	for range n {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()

// 			ch := d.Roll()
// 			roll := d.Accumulate(ch)

// 			d.mu.Lock()
// 			defer d.mu.Unlock()
// 			res = append(res, roll...)
// 		}()
// 	}

// 	wg.Wait()

// 	return res
// }
