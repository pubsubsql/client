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
	"encoding/binary"
	"encoding/json"
)

/*
--------------------+--------------------
|   message size    |    request id     |
--------------------+--------------------
|      uint32       |      uint32       |
--------------------+--------------------
*/

type NetHeader struct {
	MessageSize uint32
	RequestId   uint32
}

var HEADER_SIZE = 8
var EMPTY_HEADER = make([]byte, HEADER_SIZE, HEADER_SIZE)

func NewNetHeader(messageSize uint32, requestId uint32) *NetHeader {
	return &NetHeader{
		MessageSize: messageSize,
		RequestId:   requestId,
	}
}

func (this *NetHeader) ReadFrom(bytes []byte) {
	this.MessageSize = binary.BigEndian.Uint32(bytes)
	this.RequestId = binary.BigEndian.Uint32(bytes[4:])
}

func (this *NetHeader) WriteTo(bytes []byte) {
	binary.BigEndian.PutUint32(bytes, this.MessageSize)
	binary.BigEndian.PutUint32(bytes[4:], this.RequestId)
}

func (this *NetHeader) GetBytes() []byte {
	bytes := make([]byte, HEADER_SIZE, HEADER_SIZE)
	this.WriteTo(bytes)
	return bytes
}

func (this *NetHeader) String() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}
