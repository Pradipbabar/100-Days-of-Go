package main 

import (
    "fmt"
    "net/http"
    "time"
	"log"
)

//Create a request logging middleware handler called Logger
type Logger struct {
    handler http.Handler
}

//ServeHTTP handles the request by passing it to the real handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    l.handler.ServeHTTP(w, r)
    log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}


func middlewareTimeHandler(w http.ResponseWriter, r *http.Request) {
    curTime := time.Now().Format(time.Kitchen)
    w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}
	 
func main() {
	addr := "localhost:8080"

	 mux := http.NewServeMux()
	 mux.HandleFunc("/", middlewareTimeHandler)
	log.Printf("server is listening at %s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))

}