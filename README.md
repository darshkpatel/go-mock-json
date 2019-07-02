<img src="img/banner.png">

[![Build Status](https://travis-ci.com/darshkpatel/go-mock-json.svg?branch=master)](https://travis-ci.com/darshkpatel/go-mock-json)

## Compiling and Installation (UNIX)
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

[] Improve Logging
[] Improve URL Handling and error messages
