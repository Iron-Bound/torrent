package main

import(
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
)

func main() {
    response, error := http.Get("https://kickass.to/usearch/avengers/")
    checkerror(error)
    
    
    
    var bodyout []byte 
    fmt.Println(len(bodyout))
    bodyout, error = ioutil.ReadAll(response.Body)
    checkerror(error)
    
    var output string = string(bodyout)
    //fmt.Println(bodyout)
    fmt.Println(output)
}

func checkerror(error error) {
    if error != nil{
        log.Panic(error)
    }
}