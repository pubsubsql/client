/* Copyright (C) 2014 CompleteDB LLC.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the Apache License Version 2.0 http://www.apache.org/licenses.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package pubsubsql

import (
	"fmt"
	"testing"
)

func TestNetworkHeader(t *testing.T) {
	header1 := netHeader{
		MessageSize: 32567,
		RequestId:   9875235,
	}
	var header2 netHeader
	bytes := make([]byte, 100, 100)
	//
	header1.writeTo(bytes)
	header2.readFrom(bytes)
	//
	if header1 != header2 {
		t.Error("NetworkHeader data does not match")
	}

	fmt.Println(header1.String())
}
