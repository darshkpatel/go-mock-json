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

//APIdetails - Struct to store api details
type APIdetails struct {
	port     int
	endpoint string
	dataJSON string
	allPaths bool
}

func main() {

	//Define CLI Flags
	var (
		srcFile   = flag.String("json", "", "Path to json file to serve")
		port      = flag.Int("port", 8080, "Port to bind mock json server")
		printJSON = flag.Bool("show", false, "Pretty Print JSON")
		allPaths  = flag.Bool("allpaths", false, "Serve JSON on all URL paths")
		endpoint  = flag.String("endpoint", "/", "endpoint to serve json at, eg: /api/json ")
	)

	flag.Parse()

	// Sanity Check For Arguments

	//Check if no args provided
	if len(os.Args) == 1 {
		fmt.Println("No Arguments Provided, use flag -h for help ")
		os.Exit(1)
	}

	//Check JSON file path
	if _, err := os.Stat(*srcFile); err != nil {
		log.Fatal(*srcFile + " does not exist")
	}

	// Read JSON file
	raw, err := ioutil.ReadFile(*srcFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Loading JSON file " + *srcFile)

	//Parse JSON
	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Loaded JSON File")
	}

	var api APIdetails
	api.endpoint = *endpoint
	api.port = *port
	api.allPaths = *allPaths
	api.dataJSON = prettyprint(data)

	//Prints JSON to console if flag is set
	if *printJSON {
		fmt.Println(api.dataJSON)
	}

	//Start HTTP SERVER
	log.Printf("Starting server to serve File: %v on Port: %v", *srcFile, api.port)

	http.HandleFunc(api.endpoint, api.responseHandler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(api.port), nil))

}

func (api *APIdetails) responseHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	log.Printf("%v : %v - %v \n", readUserIP(r), r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if (api.endpoint != "" && api.endpoint == r.URL.Path) || (api.allPaths) {
		w.WriteHeader(200)
		JSONstring := api.dataJSON
		_, err := fmt.Fprint(w, JSONstring)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		w.WriteHeader(404)
		JSONstring := "404 Not Found"
		_, err := fmt.Fprint(w, JSONstring)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func prettyprint(data interface{}) string {
	prettyJSON, err := json.MarshalIndent(data, " ", "    ")
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(prettyJSON)
}
