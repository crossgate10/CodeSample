package main

import (
	"fmt"
	"net/http"
	"time"
)

//------------------------- First sample : counter
type Counter struct{
	n int
}

func (ctn *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctn.n++
	fmt.Fprintf(w, "counter = %d\n", ctn.n)
}

//------------------------- Second sample : semaphore
type ReqChan struct {
	ch chan *http.Request
}

var MaxOutstanding = 3  // equal Mutex when MaxOutstanding = 1
var sem = make(chan int, MaxOutstanding)

func process(r *http.Request){
	time.Sleep(time.Second*3)
	fmt.Printf("got some request: %+v\n", r)
}

func (reqc *ReqChan) log() {
	for req := range reqc.ch {
		fmt.Printf("start process %s request\n", req.Method)
	}
}

func (reqc *ReqChan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqc.ch <- req
	sem <- 1
	go func(req *http.Request){
		process(req)
		<-sem
	}(req)
	fmt.Fprint(w, "notification sent")
}

//------------------------- Third sample :

func main(){
	ctr := new(Counter)
	http.Handle("/counter", ctr)

	reqc := &ReqChan{make(chan *http.Request)}
	go reqc.log()
	http.Handle("/chan", reqc)

	http.ListenAndServe(":80", nil)
}
