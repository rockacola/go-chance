package chance

import (
	"errors"
	"fmt"

	"github.com/rockacola/go-chance/data"
)

// Generate a prime number between 1 and 10,000 inclusive
func (c *Chance) Prime() int {
	randomIndex := c.Rand.Intn(len(data.Primes) - 1)
	return data.Primes[randomIndex]
}

func (c *Chance) PrimeWithParams(min int, max int) (int, error) {
	if min < 1 || min > 10000 {
		return 0, errors.New("Min must be between 1 and 10000.")
	} else if max < 1 || max > 10000 {
		return 0, errors.New("Max must be between 1 and 10000.")
	} else if min > max {
		return 0, errors.New("Min must be smaller than Max.")
	}

	// find smallest acceptable prime and its index
	floorIndex := 0
	floorPrime := data.Primes[floorIndex]
	for floorPrime < min {
		floorIndex += 1
		floorPrime = data.Primes[floorIndex]
	}

	// find biggest acceptable prime and its index
	ceilingIndex := len(data.Primes) - 1
	ceilingPrime := data.Primes[ceilingIndex]
	for ceilingPrime > max {
		ceilingIndex -= 1
		ceilingPrime = data.Primes[ceilingIndex]
	}

	if floorIndex > ceilingIndex {
		return 0, errors.New(fmt.Sprintf("Unable to find a prime number between %d and %d", min, max))
	}

	if floorIndex == ceilingIndex {
		output := data.Primes[floorIndex]
		return output, nil
	}

	randomIndex := floorIndex + c.Rand.Intn(ceilingIndex-floorIndex)
	output := data.Primes[randomIndex]
	return output, nil
}
