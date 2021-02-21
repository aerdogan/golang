package main

import (
	"fmt"
	"math"
)

type sekil interface {
	alanhesapla() float64
}

type dikdortgen struct {
	en, boy float64
}

type cember struct {
	yaricap float64
}

func (d dikdortgen) alanhesapla() float64 {
	return d.en * d.boy
}

func (c cember) alanhesapla() float64 {
	return math.Pi * c.yaricap * c.yaricap
}

func hesapla(s sekil) {
	fmt.Println("Seklin Alani : ", s.alanhesapla())
}

func main() {
	d := dikdortgen{en: 10, boy: 6}
	hesapla(d)

	c := cember{yaricap: 5}
	hesapla(c)
}
