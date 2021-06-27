/*
* Implemntation of MT19937-32 based on the pseudocode from https://en.wikipedia.org/wiki/Mersenne_Twister
 */

package mersenne

import (
	"math"
)

type Mersenne struct {
	state []uint64
	index int
}

const (
	w, n, m, r, f = 32, 624, 397, 31, 1812433253
	a, u, d, s, b = 0x9908B0DF, 11, 0xFFFFFFFF, 7, 0x9D2C5680
	t, c          = 15, 0xEFC60000
	l             = 18
	lower_mask    = (1 << r) - 1
	upper_mask    = ((1 << w) - 1) &^ lower_mask
)

func Init() Mersenne {
	return Mersenne{
		state: make([]uint64, n),
		index: n + 1,
	}
}

func (mt *Mersenne) Seed(seed int64) {
	//mt.index = n
	mt.state[0] = uint64(seed)
	for i := uint64(1); i < n; i++ {
		mt.state[i] = ((1 << w) - 1) & (f*(mt.state[i-1]^(mt.state[i-1]>>uint64(w-2))) + i)
	}
	mt.index = n
}

// TODO : Implement shuffle, choice etc. for fun

func (mt *Mersenne) Random() float64 {
	return float64(mt.extract_number()) / math.Pow(2, 32)
}

func (mt *Mersenne) RandInt(a, b uint64) uint64 {
	n := mt.Random()
	return uint64(n/(1/float64(b-a)) + float64(a))
}

func (mt *Mersenne) extract_number() uint64 {
	if mt.index >= n {
		if mt.index > n {
			mt.Seed(5489)
		}
		mt.twist()
	}
	y := mt.state[mt.index]
	y ^= ((y >> u) & d)
	y ^= ((y << s) & b)
	y ^= ((y << t) & c)
	y ^= (y >> l)
	mt.index++
	return y
}

func (mt *Mersenne) twist() {
	for i := 0; i < n; i++ {
		x := mt.state[i]&upper_mask + (mt.state[(i+1)%n] & lower_mask)
		xA := x >> 1
		if (x % 2) != 0 {
			xA = xA ^ a
		}
		mt.state[i] = mt.state[(i+m)%n] ^ xA
	}
	mt.index = 0
}
