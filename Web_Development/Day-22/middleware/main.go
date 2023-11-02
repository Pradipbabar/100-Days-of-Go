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

//NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
    return &Logger{handlerToWrap}
}

func middlewareGreetingsHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Happy New Year, 2022!"))
}

func middlewareTimeHandler(w http.ResponseWriter, r *http.Request) {
    curTime := time.Now().Format(time.Kitchen)
    w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func middleware(handler http.Handler) http.Handler {
     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         fmt.Println("Executing middleware before request phase!")
         // Pass control back to the handler
         handler.ServeHTTP(w, r)         
		 fmt.Println("Executing middleware after response phase!")
     })
 }

func mainLogic(w http.ResponseWriter, r *http.Request) {
     // Business logic goes here
     fmt.Println("Executing mainHandler...")
     w.Write([]byte("OK")) 
	 } 
	 
func main() {
	addr := "localhost:8080"
     // HandlerFunc returns a HTTP Handler
    //  mainLogicHandler := http.HandlerFunc(mainLogic)
    //  http.Handle("/", middleware(mainLogicHandler))
	 
	 mux := http.NewServeMux()
	 mux.HandleFunc("/v1/greetings", middlewareGreetingsHandler)
	 mux.HandleFunc("/v1/time", middlewareTimeHandler)
    //  http.ListenAndServe(":8000", nil)
	log.Printf("server is listening at %s", addr)
    log.Fatal(http.ListenAndServe(addr, mux))

}