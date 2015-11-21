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
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestResponseData(c *C) {
	rd_empty := responseData{}

	rd := responseData{
		Status:   "status",
		Msg:      "msg",
		Action:   "action",
		Id:       "id",
		PubSubId: "psid",
		Rows:     1,
		Fromrow:  2,
		Torow:    3,
		Columns:  []string{"A", "B"},
		Data:     [][]string{[]string{"a", "b"}, []string{"c", "d"}},
	}

	rd.reset()
	c.Assert(rd, DeepEquals, rd_empty)
}
