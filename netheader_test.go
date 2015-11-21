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
	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestreadFrom(c *C) {
	nh := netHeader{}

	source_bytes := []byte("123456789")
	nh.readFrom(source_bytes)

	destination := make([]byte, 15, 15)
	nh.writeTo(destination)
	// First 8 bytes of the header need to match
	c.Assert(source_bytes[:8], DeepEquals, destination[:8])

	c.Assert(nh.RequestId, Equals, binary.BigEndian.Uint32([]byte("5678")))
	c.Assert(nh.MessageSize, Equals, binary.BigEndian.Uint32([]byte("1234")))
}
