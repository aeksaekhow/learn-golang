package main

import "fmt"

func main() {

	var h human
	aek := aek{}
	mook := mook{}
	h = aek
	h.breath()
	h = mook
	h.eat()
}

type aek struct{}

func (a aek) breath() {
	fmt.Println("Aek breath")
}

func (a aek) eat() {
	fmt.Println("Aek eat")
}

type mook struct{}

func (a mook) breath() {
	fmt.Println("Mook breath")
}

func (a mook) eat() {
	fmt.Println("Mook eat")
}

type breathable interface {
	breath()
}

type eatable interface {
	eat()
}

type human interface {
	breathable
	eatable
}
