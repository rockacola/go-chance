package chance

import "strconv"

// Generate a random US area code.
func (c *Chance) AreaCode() string {
	output, _ := c.AreaCodeWithParams(true)
	return output
}

func (c *Chance) AreaCodeWithParams(parenthesis bool) (string, error) {
	// Don't want area codes to start with 1, or have a 9 as the second digit
	digit1, _ := c.NaturalWithParams(2, 9)
	digit2, _ := c.NaturalWithParams(0, 8)
	digit3, _ := c.NaturalWithParams(0, 9)
	output := strconv.Itoa(digit1) + strconv.Itoa(digit2) + strconv.Itoa(digit3)
	if parenthesis {
		output = "(" + output + ")"
	}
	return output, nil
}
