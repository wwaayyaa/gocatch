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