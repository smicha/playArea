package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
/*	fmt.Printf("hello, world\n")*/
	wg.Add(2)
    go foo()
	go bar()
	wg.Wait()

}

func foo() {

for i := 0; i < 100; i++ {
fmt.Println("Foo:", i)
}
	wg.Done()
}


func bar() {
	for i:=0;i<100;i++ {
	fmt.Println("Bar:",i)
}
	wg.Done()

}