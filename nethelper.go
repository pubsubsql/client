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
	"errors"
	"net"
	"time"
)

// message reader
type NetHelper struct {
	conn  net.Conn
	bytes []byte
}

func NewNetHelper(conn net.Conn, bufferSize int) *NetHelper {
	var ret NetHelper
	ret.Set(conn, bufferSize)
	return &ret
}

func (this *NetHelper) Set(conn net.Conn, bufferSize int) {
	this.conn = conn
	this.bytes = make([]byte, bufferSize, bufferSize)
}

func (this *NetHelper) Close() {
	if this.conn != nil {
		this.conn.Close()
		this.conn = nil
	}
}

func (this *NetHelper) Valid() bool {
	return this.conn != nil
}

func (this *NetHelper) WriteMessage(bytes []byte) error {
	leftToWrite := len(bytes)
	for {
		written, err := this.conn.Write(bytes)
		if err != nil {
			return err
		}
		leftToWrite -= written
		if leftToWrite == 0 {
			break
		}
		bytes = bytes[written:]
	}
	return nil
}

func (this *NetHelper) WriteHeaderAndMessage(requestId uint32, bytes []byte) error {
	err := this.WriteMessage(NewNetHeader(uint32(len(bytes)), requestId).GetBytes())
	if err != nil {
		return err
	}
	return this.WriteMessage(bytes)
}

func (this *NetHelper) ReadMessageTimeout(milliseconds int64) (*NetHeader, []byte, error, bool) {
	this.conn.SetReadDeadline(time.Now().Add(time.Duration(milliseconds) * time.Millisecond))
	header, bytes, err := this.ReadMessage()
	timedout := false
	if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
		timedout = true
		err = nil
	}
	return header, bytes, err, timedout
}

func (this *NetHelper) ReadMessage() (*NetHeader, []byte, error) {
	// header
	read, err := this.conn.Read(this.bytes[0:HEADER_SIZE])
	if err != nil {
		return nil, nil, err
	}
	if read < HEADER_SIZE {
		err = errors.New("Failed to read header.")
		return nil, nil, err
	}
	var header NetHeader
	header.ReadFrom(this.bytes)
	// prepare buffer
	if len(this.bytes) < int(header.MessageSize) {
		this.bytes = make([]byte, header.MessageSize, header.MessageSize)
	}
	// message
	bytes := this.bytes[:header.MessageSize]
	left := len(bytes)
	message := bytes
	read = 0
	for left > 0 {
		bytes = bytes[read:]
		read, err = this.conn.Read(bytes)
		if err != nil {
			return nil, nil, err
		}
		left -= read
	}
	return &header, message, nil
}
