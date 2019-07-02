package main

import (
	"encoding/json"
	"fmt"
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
	dataJSON interface{}
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
			t.Error("Pretty Print Not Linting JSON")
			t.Errorf("Expected: %v but Got: %v", testUnit.expect, testUnit.got)
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
		fmt.Print(data)
		if testUnit.got = prettyprint(data); strings.Compare(testUnit.got, testUnit.expect) != 0 {
			t.Error("Pretty Print Not Linting JSON")
			t.Errorf("Expected: %v but Got: %v", testUnit.expect, testUnit.got)
			bs := []byte(testUnit.got)
			fmt.Print(bs)
		}

	}
}
