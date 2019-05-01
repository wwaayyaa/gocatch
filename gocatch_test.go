package gocatch

import (
	"fmt"
	"testing"
)

func TestCatch(t *testing.T){
	fmt.Println("test 1")
	defer Catch(func(e interface{}){
		fmt.Printf("%v \n", e)
	})
	z := 0
	a := 100 / z
	t.Errorf("%v", a)
}
func TestCatchWithStack(t *testing.T){
	fmt.Println("test 2")
	defer CatchWithStack(func(e interface{}){
		fmt.Printf("%v \n", e)
	})
	z := 0
	a := 100 / z
	t.Errorf("%v", a)
}
func TestCatchStack(t *testing.T){
	fmt.Println("test 3")
	defer CatchStack(func(e interface{}, stack string){
		fmt.Printf("%v \n", e)
		fmt.Printf("%v \n", stack)
	})
	z := 0
	a := 100 / z
	t.Errorf("%v", a)
}