package chance

import (
	"strconv"
)

// Generate a random Brazilian tax id.
func (c *Chance) Cpf() string {
	// Generate 9 random digits
	ns := make([]int, 0)
	for i := 0; i < 9; i++ {
		n, _ := c.NaturalWithParams(0, 9)
		ns = append(ns, n)
	}

	d1 := ns[8]*2 + ns[7]*3 + ns[6]*4 + ns[5]*5 + ns[4]*6 + ns[3]*7 + ns[2]*8 + ns[1]*9 + ns[0]*10
	d1 = 11 - (d1 % 11)
	if d1 >= 10 {
		d1 = 0
	}

	d2 := d1*2 + ns[8]*3 + ns[7]*4 + ns[6]*5 + ns[5]*6 + ns[4]*7 + ns[3]*8 + ns[2]*9 + ns[1]*10 + ns[0]*11
	d2 = 11 - (d2 % 11)
	if d2 >= 10 {
		d2 = 0
	}

	output := strconv.Itoa(ns[0]+ns[1]+ns[2]) + "." +
		strconv.Itoa(ns[3]+ns[4]+ns[5]) + "." +
		strconv.Itoa(ns[6]+ns[7]+ns[8]) + "-" +
		strconv.Itoa(d1) + strconv.Itoa(d2)
	return output
}
