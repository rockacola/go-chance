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

	// find biggest acceptable prime and its index
	capIndex := len(data.Primes) - 1
	capPrime := data.Primes[capIndex]
	for capPrime > max {
		capIndex -= 1
		capPrime = data.Primes[capIndex]
	}

	if min > capIndex {
		return 0, errors.New(fmt.Sprintf("Unable to find a prime number between %d and %d", min, max))
	}

	randomIndex := min + c.Rand.Intn(capIndex-min)
	output := data.Primes[randomIndex]
	return output, nil
}
