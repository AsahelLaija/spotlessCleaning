package main
import (
    "fmt"
    "log"
    "net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello"{
        http.Error(w, "404 no se encontro.", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method is not supported", http.StatusNotFound)
        return
    }
    fmt.Fprint(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParserForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request succesful")
    name := r.FormValue("name")
    address := r.FormValue("address")
    users := r.FormValue("users")
    fmt.Println(name)
    fmt.Println(address)
    fmt.Println(users)

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
    fmt.Fprintf(w, "users = %s\n", users)
}
func main () {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello",  helloHandler)

    fmt.Printf("Startinc server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
