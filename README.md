# gocatch
A Golang error capture package that calls back a custom method after capturing an error.  Golang | Catch | Recover | Try | 

## install

### normal
``` go get github.com/wwaayyaa/gocatch```

### dep
``` 
dep ensure -add github.com/wwaayyaa/gocatch
dep ensure -v
```

### Methods
#### Catch
if your goroutine crash, will pass ```error``` to you callback with ```interface{}``` type.

```go
defer gocatch.Catch(func(e interface{}){
	///
}
```
#### CatchWithStack
if your goroutine crash, will print stack and pass ```error``` to you callback with ```interface
}{}``` type.

```go
defer gocatch.Catch(func(e interface{}){
	///
}
```
#### CatchStack
if your goroutine crash, will pass ```error``` and stack-string to you callback.

```go
defer gocatch.Catch(func(e interface{}, stack string){
	///
}
```
### example

```go
func foo() {
	defer gocatch.Catch(func(e
	interface{}){
		fmt.Printf("catch %+v", e)
		// maybe custom func throw panic
		panic("panic again.")
		})
	fmt.Printf("d	oing...\r\n")
	panic("i am panic!")
}

func main() {
	go foo()
	time.Sleep(1 * time.Second)
	fmt.Println("end\r\n")
}


// doing...
// catch panic: i am panic!
// 2019/03/07 15:00:25 Catch package: callback error: panic again.
// end
```

