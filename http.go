package GO_TIMELINE_HTTP

import (
    "net/http"
    "log"
)

type TestHandler struct {
    str string
}

func SayHello(w http.ResponseWriter, r *http.Request){
    log.Printf("HandleFunc")
    w.Write([]byte(string("HandleFunc")))
}
//ServeHTTP方法，绑定TestHandler
func (th *TestHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
    log.Printf("Handle")
    w.Write([]byte(string("Handle")))
}

func main(){
    http.Handle("/", &TestHandler{"Hi"})
    http.HandleFunc("/test", SayHello)
    http.ListenAndServe("127.0.0.1:8000",nil)
}