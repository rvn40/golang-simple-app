package main

import (
  "fmt"
  "log"
	"os"
  "net/http"
	"github.com/olekukonko/tablewriter"
)

func main() {
  http.HandleFunc("/", handlerTable)
  http.HandleFunc("/raw", handler)
  http.HandleFunc("/temp", handlerTemp)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type requestHeaders struct {
  Name  []string
  Value []string
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
  //Iterate over all header fields
  for k, v := range r.Header {
    fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
  }
  fmt.Fprintf(w, "Host = %q\n", r.Host)
  fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
  //Get value for a specified token
  fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])
}

func handlerTable(w http.ResponseWriter, r *http.Request) {
	data := [][]string{
    []string{"A", "The Good", "500"},
    []string{"B", "The Very very Bad Man", "288"},
    []string{"C", "The Ugly", "120"},
    []string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	for _, v := range data {
	    table.Append(v)
	}
	table.Render() // Send output

}

func handlerTemp(w http.ResponseWriter, r *http.Request) {
  var request [][]string
  table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Values"})

  for k, v := range r.Header {
    for _, value := range v {
        request = append(request,[]string{k, value})
    }
  }

  for _, v := range request {
	    table.Append(v)
	}
	table.Render()

}