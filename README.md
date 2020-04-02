<img src="img/banner.png">

<p align="center" >

<a href="https://travis-ci.com/darshkpatel/go-mock-json">
<img src="https://travis-ci.com/darshkpatel/go-mock-json.svg?branch=master">
</a>

<a href="https://github.com/darshkpatel/go-mock-json/releases">
<img src="https://img.shields.io/github/release/darshkpatel/go-mock-json.svg">
</a>

<a href="https://github.com/darshkpatel/go-mock-json/blob/master/LICENSE.md">
<img src="https://img.shields.io/github/license/darshkpatel/go-mock-json.svg?color=yellow">
</a>

</p>

## Installation

Easiest way to install this , is by executing the following command if you have golang setup on your machine 

`go get -u github.com/darshkpatel/go-mock-json/cmd/go-mock-json`

You can also download one of the compiled binaries from the [Releases Page](https://github.com/darshkpatel/go-mock-json/releases) and add it to the system PATH or copy it to the bin folder.

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
