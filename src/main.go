package main

import (
  "net/http"
  "html/template"
  "log"
)


var tmpl *template.Template
type Todo struct {
    Item string
    Done bool
}
type PageData struct{
    Title string
    Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
   data := PageData {
     Title: "TODO List",
     Todos: [] Todo{
       {Item: "install", Done: true},
       {Item: "Learn", Done: false},
       {Item: "Learn more", Done: false},
       {Item: "Learn more and more", Done: false},
     },
   }

   tmpl.Execute(w, data)
}

func main(){

   mux := http.NewServeMux()
   tmpl = template.Must(template.ParseFiles("../template/index.gohtml"))
   fs :=  http.FileServer(http.Dir("../static"))
   mux.Handle("/static/", http.StripPrefix("/static/", fs))
   mux.HandleFunc("/todo", todo)

   log.Fatal(http.ListenAndServe(":9091", mux))


}
