package main

import "testing"

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

func (c Cat) Quack() {}

func BenchmarkDirectCall(b *testing.B) {
	c := Cat{Name: "draven"}
	for n := 0; n < b.N; n++ {
		// MOVQ	AX, (SP)
		// MOVQ	$6, 8(SP)
		// CALL	"".Cat.Quack(SB)
		c.Quack()
	}
}

func BenchmarkDynamicDispatch(b *testing.B) {
	c := Duck(Cat{Name: "draven"})
	for n := 0; n < b.N; n++ {
		// MOVQ	16(SP), AX
		// MOVQ	24(SP), CX
		// MOVQ	AX, "".d+32(SP)
		// MOVQ	CX, "".d+40(SP)
		// MOVQ	"".d+32(SP), AX
		// MOVQ	24(AX), AX
		// MOVQ	"".d+40(SP), CX
		// MOVQ	CX, (SP)
		// CALL	AX
		c.Quack() //值类型的变量实现接口对性能影响较大
	}
}

func BenchmarkDynamicDispatch2(b *testing.B) {
	c := Duck(&Cat{Name: "draven"})
	for n := 0; n < b.N; n++ {
		// MOVQ	"".d+56(SP), AX
		// MOVQ	24(AX), AX
		// MOVQ	"".d+64(SP), CX
		// MOVQ	CX, (SP)
		// CALL	AX
		c.Quack() //指针类型的变量实现接口对性能影响小一点,但还是有影响
	}
}

func BenchmarkDynamicDispatch3(b *testing.B) {
	c := Duck(&Cat{Name: "draven"})
	for n := 0; n < b.N; n++ {
		c.(*Cat).Quack() //断言后再调用，几乎没有影响
	}
}
