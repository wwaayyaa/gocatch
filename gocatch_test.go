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
func TestCatchReturnNull(t *testing.T){
	fmt.Println("test 1")
	defer Catch(nil)
	z := 0
	a := 100 / z
	t.Errorf("%v", a)
}
func TestCatchCallbackError(t *testing.T){
	fmt.Println("test 1")
	defer Catch(func(i interface{}) {
		panic("internal")
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
func TestCatchWithStackReturnNull(t *testing.T){
	fmt.Println("test 2")
	defer CatchWithStack(nil)
	z := 0
	a := 100 / z
	t.Errorf("%v", a)
}
func TestCatchWithStackCallbackError(t *testing.T){
	fmt.Println("test 2")
	defer CatchWithStack(func(i interface{}) {
		panic("internal")
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
func TestCatchStackReturnNull(t *testing.T){
	fmt.Println("test 3")
	defer CatchStack(nil)
	z := 0
	a := 100 / z
	t.Errorf("%v", a)
}
func TestCatchStackCallbackError(t *testing.T){
	fmt.Println("test 3")
	defer CatchStack(func(i interface{}, stack string) {
		panic("internal")
	})
	z := 0
	a := 100 / z
	t.Errorf("%v", a)
}