package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/xtdlib/trace"
)

var counter atomic.Int64

func main() {
	http.HandleFunc("/x", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	// tr := trace.New(fmt.Sprintf("mypkg.Foo_%d", counter.Add(1)), "")
	tr := trace.New("foo", "")
	defer tr.Finish()

	if counter.Load() < 9 {
		time.Sleep(time.Second)
		counter.Add(1)
	}
	tr.LazyPrintf("some message %q", "to log")
	tr.LazyPrintf("somethingImportant failed: %v", fmt.Errorf("some error"))
	// tr.SetError()
	fmt.Fprintf(w, "Hello, World!")
}
