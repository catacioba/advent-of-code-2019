package ch22

import (
	"aoc/util"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type deck []int

func newDeck(size int) deck {
	deck := make([]int, size)
	for i := 0; i < size; i++ {
		deck[i] = i
	}
	return deck
}

func (d deck) print() {
	for i := range d {
		fmt.Printf("%d ", d[i])
	}
	fmt.Println()
}

func dealIntoNewStack(d deck) deck {
	n := len(d)
	newDeck := make(deck, n)
	for i := range d {
		newDeck[n-i-1] = d[i]
	}
	return newDeck
}

func dealWithIncrement(d deck, c int) deck {
	n := len(d)
	newDeck := make(deck, n)
	for i := range d {
		newDeck[i*c%n] = d[i]
	}
	return newDeck
}

func cut(d deck, c int) deck {
	n := len(d)
	newDeck := make(deck, n)

	if c > 0 {
		for i := c; i < n; i++ {
			newDeck[i-c] = d[i]
		}
		for i := 0; i < c; i++ {
			newDeck[n-c+i] = d[i]
		}
	} else {
		for i := n + c; i < n; i++ {
			newDeck[i-n-c] = d[i]
		}
		for i := 0; i < n+c; i++ {
			newDeck[-c+i] = d[i]
		}
	}
	return newDeck
}

func PartOne() {
	lines := util.ReadLines("ch22/input.txt")

	d := newDeck(10007)
	for _, line := range lines {
		if line == "deal into new stack" {
			d = dealIntoNewStack(d)
		} else if strings.HasPrefix(line, "deal with increment") {
			c := strings.TrimPrefix(line, "deal with increment ")
			ci, _ := strconv.Atoi(c)
			d = dealWithIncrement(d, ci)
		} else {
			c := strings.TrimPrefix(line, "cut ")
			ci, _ := strconv.Atoi(c)
			d = cut(d, ci)
		}
	}

	for idx, v := range d {
		if v == 2019 {
			fmt.Println(idx)
			return
		}
	}
}

func trackNewStackPosition(p, n int64) int64 {
	return n - 1 - p
}

func trackInverseStackPosition(p, n int64) int64 {
	return trackNewStackPosition(p, n)
}

func trackIncrementPosition(p, n, c int64) int64 {
	return p * c % n
}

func trackInverseIncrementPosition(p, n, c int64) int64 {
	bp := big.NewInt(p)
	bi := big.NewInt(modularInverse(c, n))
	bn := big.NewInt(n)
	var res big.Int
	res.Mul(bp, bi)
	res.Mod(&res, bn)
	return res.Int64()
}

func trackCutPosition(p, n, c int64) int64 {
	return (n + p - c) % n
}

func trackInverseCutPosition(p, n, c int64) int64 {
	return trackCutPosition(p, n, -c)
}

func gcdExtended(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := gcdExtended(b, a%b)

	x := y1
	y := x1 - y1*(a/b)

	return gcd, x, y
}

func modularInverse(a, n int64) int64 {
	gcd, x, _ := gcdExtended(a, n)
	if gcd != 1 {
		panic(fmt.Sprintf("no modular inverse for %d and %d", a, n))
	}
	z := (x%n + n) % n
	if z < 0 {
		panic("invalid")
	}
	return z
}

func forward(p, n int64, lines []string) int64 {
	for _, line := range lines {
		if line == "deal into new stack" {
			p = trackNewStackPosition(p, n)
		} else if strings.HasPrefix(line, "deal with increment") {
			c := strings.TrimPrefix(line, "deal with increment ")
			ci, _ := strconv.ParseInt(c, 10, 64)
			p = trackIncrementPosition(p, n, ci)
		} else {
			c := strings.TrimPrefix(line, "cut ")
			ci, _ := strconv.ParseInt(c, 10, 64)
			p = trackCutPosition(p, n, ci)
		}
	}
	return p
}

func backwards(p, n int64, lines []string) int64 {
	for idx := len(lines) - 1; idx >= 0; idx-- {
		line := lines[idx]
		if line == "deal into new stack" {
			p = trackInverseStackPosition(p, n)
		} else if strings.HasPrefix(line, "deal with increment") {
			c := strings.TrimPrefix(line, "deal with increment ")
			ci, _ := strconv.ParseInt(c, 10, 64)
			p = trackInverseIncrementPosition(p, n, ci)
		} else {
			c := strings.TrimPrefix(line, "cut ")
			ci, _ := strconv.ParseInt(c, 10, 64)
			p = trackInverseCutPosition(p, n, ci)
		}
		if p < 0 {
			fmt.Printf("line: %s\n", line)
			panic("invalid")
		}
	}
	return p
}

func find_coefficients(n, p, pp int64, lines []string) (int64, int64) {
	cnt := 0
	b := int64(1)

	for _, line := range lines {
		if strings.HasPrefix(line, "deal with increment") {
			c := strings.TrimPrefix(line, "deal with increment ")
			ci, _ := strconv.ParseInt(c, 10, 64)
			b = b * ci % n
		} else {
			cnt += 1
		}
	}

	a := pp % n
	if cnt%2 == 1 {
		b = -b
	}
	fmt.Printf("b=%d a=%d\n", b, a)
	d := b * p % n
	a = (n + a - d) % n

	return b, a
}

func pow(a *big.Int, b, n int64) *big.Int {
	if b == 1 {
		return new(big.Int).Set(a)
	}
	if b == 0 {
		return big.NewInt(int64(1))
	}
	nb := big.NewInt(n)
	aa := new(big.Int).Mul(a, a)
	aa.Mod(aa, nb)
	x := pow(aa, b/2, n)
	if b%2 == 0 {
		return x
	} else {
		x.Mul(x, a)
		x.Mod(x, nb)
		return x
	}
}

func forward_brute(b, a, p, t, n int64) int64 {
	x := p
	for idx := int64(0); idx < t; idx++ {
		// x = (n + ((n+b)%n)*x + a) % n
		x = (n + (b*x%n+a)%n) % n
	}
	return x
}

func forward_exp(b, a, p, t, n int64) int64 {
	return (mypow(b, t, n)*p%n + a*((mypow(b, t, n)-1)*modularInverse(b-1, n))%n + n) % n
}

func mypow(a, b, n int64) int64 {
	p := int64(1)
	for idx := int64(0); idx < b; idx++ {
		p = p * a % n
	}
	return p
}

func PartTwo() {
	lines := util.ReadLines("ch22/input.txt")

	n := int64(119315717514047)
	t := int64(101741582076661)
	p := int64(2020)

	pp := forward(p, n, lines)
	b, a := find_coefficients(n, p, pp, lines)

	ni := big.NewInt(n)
	pi := big.NewInt(p)

	bi := big.NewInt(b)
	ai := big.NewInt(a)

	bni := pow(bi, t, n)
	bni.Add(bni, ni)
	bni.Mod(bni, ni)

	bi_inv := new(big.Int).ModInverse(bni, ni)

	// bn := mypow((n+b)%n, t, n)
	// b_inv := modularInverse(bn, n)
	// an := a * ((bn - 1) % n * modularInverse((n+b-1)%n, n)) % n
	// fmt.Printf("rev exp: %d\n", (b_inv*(pp-an)%n+n)%n)

	ani := new(big.Int).Sub(bni, big.NewInt(1))
	ani.Mul(ani, ai)
	ani.Mul(ani, new(big.Int).ModInverse(new(big.Int).Sub(bi, big.NewInt(1)), ni))
	ani.Mod(ani, ni)

	next := new(big.Int).Add(new(big.Int).Mul(bni, pi), ani)
	next.Mod(next, ni)
	next.Add(next, ni)
	next.Mod(next, ni)
	fmt.Printf("next >> %s\n", next.String())

	reverse := new(big.Int).Sub(pi, ani)
	reverse.Mul(reverse, bi_inv)
	reverse.Mod(reverse, ni)
	reverse.Add(reverse, ni)
	reverse.Mod(reverse, ni)

	fmt.Printf("reverse >> %s\n", reverse.String())

}
