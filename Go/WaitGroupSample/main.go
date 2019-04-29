package main

import (
	"log"
	"os"
	"sync"
)

func main(){
	logger := log.New(os.Stdout, "", 0)

	var wg sync.WaitGroup

	wg.Add(20)

	for i:=0; i<10; i++{
		go func(){
			defer wg.Done()
			logger.Println(i)
		}()
	}

	for i:=0; i<10; i++{
		go func(i int){
			defer wg.Done()
			logger.Println(i)
		}(i)
	}

	wg.Wait()
}
