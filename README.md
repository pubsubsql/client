[![Build Status](https://travis-ci.org/pubsubsql/client.svg?branch=master)](https://travis-ci.org/pubsubsql/client)

# Introduction

This is Go client for [PubSubSQL](https://github.com/pubsubsql/pubsubsql), an in-memory database with SQL-like syntax and usage but offering PUB-SUB functionality and MySQL as secondary datastore.
Supported versions of Go are 1.3 and up.


# Example
```go
package main

import (
	"fmt"

	"github.com/pubsubsql/client"
)

func checkError(client *pubsubsql.Client, str string) {
	if client.Failed() {
		fmt.Println("Error:", client.Error(), str)
		os.Exit(1)
	}
}

func main() {
	client := new(pubsubsql.Client)

	address := "localhost:7777"
	client.Connect(address)
	checkError(client, "client connect failed")

	filepath := "/bin/ls"

	mimetype, err := mm.TypeByFile(filepath)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	fmt.Printf("%s -> %s\n", filepath, mimetype)
}
```
