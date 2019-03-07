package gocatch

import "log"

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