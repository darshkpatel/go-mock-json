package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	HeaderContentType              = "Content-Type"
	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + "charset=UTF-8"
	MIMETextPlain                  = "text/plain"
	MIMETextPlainCharsetUTF8       = MIMETextPlain + "; " + "charset=UTF-8"
)

type API struct {
	port     int
	endpoint string
	dataJSON interface{}
}

var api API

func main() {
	//Parse CLI Flags
	var (
		srcFile   = flag.String("json", "", "Path to json file to serve")
		port      = flag.Int("port", 8080, "Port to bind mock json server")
		printJSON = flag.Bool("show", false, "Pretty Print JSON")
		endpoint  = flag.String("endpoint", "/", "endpoint to serve json at, eg: /api/json ")
	)

	flag.Parse()

	// Sanity Check For Arguments

	//Check if no args provided
	if len(os.Args) == 1 {
		fmt.Println("No Arguments Provided, use flag -h for help ")
		os.Exit(1)
	}

	//Check json file path

	if _, err := os.Stat(*srcFile); err != nil {
		log.Fatal(*srcFile + " does not exist")
	}

	raw, err := ioutil.ReadFile(*srcFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Loading JSON file " + *srcFile)
	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		log.Fatal(err)
	}

	api.endpoint = *endpoint
	api.port = *port
	api.dataJSON = data

	log.Println("Loaded JSON File")

	if *printJSON {
		fmt.Println(prettyPrint(api.dataJSON))
	}

	//Start HTTP SERVER
	log.Printf("Serving File %v on port %v", *srcFile, api.port)
	http.HandleFunc(api.endpoint, Response)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(api.port), nil))

}

func Response(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	r.ParseForm()
	fmt.Println("method:", r.Method)
	fmt.Println("path:", r.URL.Path)
	w.Header().Set(HeaderContentType, MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(200)
	s := prettyPrint(api.dataJSON)
	b := []byte(s)
	w.Write(b)

}

func prettyPrint(data interface{}) string {
	prettyJson, err := json.MarshalIndent(data, " ", "    ")
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(prettyJson)
}
