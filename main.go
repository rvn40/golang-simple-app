package main

import (
  "fmt"
  "log"
  "sort"
  "net/http"
	"github.com/olekukonko/tablewriter"
	"github.com/prometheus/client_golang/prometheus"
//  "html/template"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	httpReqCounters = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "webapp_counter",
		Help: "Root endpoint total http requests counter.",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(httpReqCounters)
}


func main() {
  cpuTemp.Set(69.3)

  http.HandleFunc("/", handlerTemp)
  http.HandleFunc("/raw", handlerRaw)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerTemp(w http.ResponseWriter, r *http.Request) {
  var request [][]string
  keys := make([]string, 0, len(r.Header))
  
  table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Name", "Values"})

  for k := range r.Header {
	  keys = append(keys, k)
	}

  sort.Strings(keys)

  // convert value from map to string
  // r.Header disamble from multidimensional string array to string array 
  for _, key := range keys {
  // Resamble string array to multidimensional string array
    for _, value := range r.Header[key] {
      request = append(request, []string{key, value})
	    //fmt.Println(key, r.Header[key])
    }
  }

  for _, v := range request {
	    table.Append(v)
	}
	table.Render()
  // fmt.Fprintf(w, "%s", table.Render()
  httpReqCounters.Inc()
}

func handlerRaw(w http.ResponseWriter, r *http.Request) {
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
