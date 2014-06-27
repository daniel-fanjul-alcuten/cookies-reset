package main

import (
	"flag"
	"log"
	"math/big"
)

func main() {

	m := flag.Uint64("m", 100, "multiplier")
	k := flag.Uint64("k", 0, "kittens")
	l := flag.Uint64("l", 0, "milk")
	hc1 := flag.Uint64("hc1", 0, "heavenly chips before")
	hc2 := flag.Uint64("hc2", 0, "heavenly chips after")
	flag.Parse()

	multiplier := new(big.Int).SetUint64(*m)
	milk := new(big.Int).SetUint64(*l)
	chips1 := new(big.Int).SetUint64(*hc1)
	chips2 := new(big.Int).SetUint64(*hc2)

	kittensMultiplier := big.NewRat(1, 1)
	if *k >= 1 {
		f := big.NewInt(5)
		f.Mul(f, milk)
		g := new(big.Rat).SetFrac(f, big.NewInt(10000))
		g.Add(g, big.NewRat(1, 1))
		kittensMultiplier.Mul(kittensMultiplier, g)
	}
	if *k >= 2 {
		f := big.NewInt(10)
		f.Mul(f, milk)
		g := new(big.Rat).SetFrac(f, big.NewInt(10000))
		g.Add(g, big.NewRat(1, 1))
		kittensMultiplier.Mul(kittensMultiplier, g)
	}
	if *k >= 3 {
		f := big.NewInt(20)
		f.Mul(f, milk)
		g := new(big.Rat).SetFrac(f, big.NewInt(10000))
		g.Add(g, big.NewRat(1, 1))
		kittensMultiplier.Mul(kittensMultiplier, g)
	}
	if *k >= 4 {
		f := big.NewInt(20)
		f.Mul(f, milk)
		g := new(big.Rat).SetFrac(f, big.NewInt(10000))
		g.Add(g, big.NewRat(1, 1))
		kittensMultiplier.Mul(kittensMultiplier, g)
	}
	if *k >= 5 {
		f := big.NewInt(20)
		f.Mul(f, milk)
		g := new(big.Rat).SetFrac(f, big.NewInt(10000))
		g.Add(g, big.NewRat(1, 1))
		kittensMultiplier.Mul(kittensMultiplier, g)
	}

	// multiplier before kittensMultiplier
	m1 := new(big.Rat).Quo(new(big.Rat).SetInt(multiplier), kittensMultiplier)
	// multiplier before heavenly chips 1
	m2 := new(big.Rat).Sub(m1, new(big.Rat).SetInt(new(big.Int).Mul(chips1, big.NewInt(2))))

	// multiplier after heavenly chips 2
	m3 := new(big.Rat).Add(m2, new(big.Rat).SetInt(new(big.Int).Mul(chips2, big.NewInt(2))))
	// multiplier after kittensMultiplier
	m4 := new(big.Rat).Mul(m3, kittensMultiplier)

	log.Printf("old multiplier: %v", multiplier)
	log.Printf("new multiplier: %v", new(big.Int).Quo(m4.Num(), m4.Denom()))
	log.Printf("ratio: %v", new(big.Rat).Quo(m4, new(big.Rat).SetInt(multiplier)).FloatString(4))
}
