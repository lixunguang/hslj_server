package service

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type Name interface {
	Run()
}

type AA struct {
	x int
	X int
}

type BB struct {
	y int
}

type CC struct {
	y int
}

func (a *AA) Run() {
	a.x = a.X
	fmt.Println("a run", a.x)
}

func (b BB) Run() {
	fmt.Println("b run")
}

func FunA(n Name) {

	res := reflect.TypeOf(n).String()
	fmt.Println(res)
	if strings.Contains(res, "*") {
		fmt.Printf("FunA:%p\n", n)
	} else {
		fmt.Printf("FunA:%p\n", &n)
	}

	n.Run()
}

func TestXXX(t *testing.T) {

	var a = AA{X: 11}
	fmt.Printf("&a:%p\n", &a)
	FunA(&a)

	var b = BB{}
	fmt.Printf("&b:%p\n", &b)
	FunA(b)

}
