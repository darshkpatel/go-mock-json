<img src="img/banner.png">

[![Build Status](https://travis-ci.com/darshkpatel/go-mock-json.svg?branch=master)](https://travis-ci.com/darshkpatel/go-mock-json)

## Installation

You can download one of the compiled binaries from the [Releases Page](https://github.com/darshkpatel/go-mock-json/releases) and add it to the system PATH or copy it to the bin folder.

You can also compile from source using the instructions mentioned below

## Compiling from Source (UNIX)
To Build and run binary  
```  git clone https://github.com/darshkpatel/go-mock-json && cd go-mock-json  ```

``` go build cmd/go-mock-json/go-mock-json.go ```

```./cmd/go-mock-json/go-mock-json```

To install:

You can either copy the binary to ```bin``` directory of your system OR add go-mock-json to PATH


## Usage
```
Usage of ./go-mock-json:
  -endpoint string 
        Endpoint to serve json at, eg: /api/json  (default "/")

  -json string
    	Path to json file to serve
        
  -port int
    	Port to bind mock json server (default 8080)

  -show
    	Pretty Print JSON
```

### To-Do
* Improve Logging
* Improve URL Handling and error messages
