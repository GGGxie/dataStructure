package main

type AA struct {
	a *A
}

type A struct {
	B int
}

func (a *A) change() {
	a.B = 3
}

func main() {
	a := A{1}
	aa := AA{
		&a,
	}
	aa.a.change()
}
