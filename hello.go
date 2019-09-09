//package main

//import "fmt"

//func main() {
//    fmt.Println("Hello, World")
//}


package main

import "fmt"
import "net/http"

func main() {
    http.HandleFunc("/", handlerIndex)
    http.HandleFunc("/index", handlerIndex)
    http.HandleFunc("/hello", handlerHello)

        var address = ":8080"
        fmt.Printf("server started at %s\n", address)

        server := new(http.Server)
        server.Addr = address
        err := server.ListenAndServe()
        if err != nil {
                fmt.Println(err.Error())
           }
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
    var message = "Welcome"
    w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
    var message = "Hello World!"
    w.Write([]byte(message))
}
