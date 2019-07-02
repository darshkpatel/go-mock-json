package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestJSON struct {
	data   []byte
	expect string
	got    string
}

type TestAPI struct {
	port     int
	endpoint string
	dataJSON string
}

func TestPrettyprint_empty(t *testing.T) {
	t.Log("Testing JSON Parsing and Pretty Print")
	var data map[string]interface{}

	var testUnit TestJSON
	testUnit.data = []byte(`{}`)
	testUnit.expect = `{}`

	if err := json.Unmarshal(testUnit.data, &data); err != nil {
		t.Error("Cannot Parse JSON , Probably a problem with json.Unmarshal")
	} else {
		if testUnit.got = prettyprint(data); testUnit.got != testUnit.expect {
			t.Errorf("Unexpected Linted JSON , Expected: %v but Got: %v", testUnit.expect, testUnit.got)
		}

	}
}

func TestPrettyprint_keys(t *testing.T) {
	t.Log("Testing JSON Parsing and Pretty Print")
	var data map[string]interface{}

	var testUnit TestJSON
	testUnit.data = []byte(`{"key": "value"}`)

	//Storing Expect value as byte string to prevent errors due to multiline string indentation
	bytestring := []rune{123, 10, 32, 32, 32, 32, 32, 34, 107, 101, 121, 34, 58, 32, 34, 118, 97, 108, 117, 101, 34, 10, 32, 125}
	testUnit.expect = string(bytestring)

	if err := json.Unmarshal(testUnit.data, &data); err != nil {
		t.Error("Cannot Parse JSON , Probably a problem with json.Unmarshal")
	} else {
		if testUnit.got = prettyprint(data); strings.Compare(testUnit.got, testUnit.expect) != 0 {
			t.Errorf("Unexpected Linted JSON ,  Expected: %v but Got: %v", testUnit.expect, testUnit.got)
			bs := []byte(testUnit.got)
			fmt.Print(bs)
		}

	}
}

func TestAPIdetails(t *testing.T) {
	t.Run("responseHandler", func(t *testing.T) {
		var api APIdetails
		api.port = 8080
		api.endpoint = "/test/endpoint"

		//Storing Expect value as byte string to prevent errors due to multiline string indentation
		bytestring := []rune{123, 10, 32, 32, 32, 32, 32, 34, 107, 101, 121, 34, 58, 32, 34, 118, 97, 108, 117, 101, 34, 10, 32, 125}
		api.dataJSON = string(bytestring)

		req, err := http.NewRequest("GET", api.endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(api.responseHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := api.dataJSON
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}

	})
}
