package chance

import (
	"math/rand"
	"time"
)

type Chance struct {
	Seed int64
	Rand *rand.Rand
}

func NewChance() *Chance {
	return NewChanceWithSeed(time.Now().UTC().UnixNano())
}

func NewChanceWithSeed(seed int64) *Chance {
	c := new(Chance)
	c.Seed = seed
	c.Rand = rand.New(rand.NewSource(c.Seed))
	return c
}
