package gocatch

import (
	"fmt"
	"log"
	"runtime"
)

func Catch(f func(interface{})){
	if err := recover(); err != nil{
		if f == nil{
			return
		}
		defer func(){
			if e := recover(); e != nil{
				log.Printf("Catch package: callback error: %v", e)
			}
		}()
		f(err)
	}
}

//print stack
func CatchWithStack(f func(interface{})){
	if err := recover(); err != nil{
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		fmt.Printf("[error stack] ==> %s\n", string(buf[:n]))

		if f == nil{
			return
		}
		defer func(){
			if e := recover(); e != nil{
				log.Printf("Catch package: callback error: %v", e)
			}
		}()
		f(err)
	}
}

//
func CatchStack(f func(interface{}, string)){
	if err := recover(); err != nil{
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)

		if f == nil{
			return
		}
		defer func(){
			if e := recover(); e != nil{
				log.Printf("Catch package: callback error: %v", e)
			}
		}()
		f(err, string(buf[:n]))
	}
}