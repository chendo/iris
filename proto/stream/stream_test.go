// Iris - Distributed Messaging Framework
// Copyright 2013 Peter Szilagyi. All rights reserved.
//
// Iris is dual licensed: you can redistribute it and/or modify it under the
// terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// The framework is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// Alternatively, the Iris framework may be used in accordance with the terms
// and conditions contained in a signed written agreement between you and the
// author(s).
//
// Author: peterke@gmail.com (Peter Szilagyi)
package stream

import (
	"net"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	sink, quit, err := Listen(31415)
	if err != nil {
		t.Errorf("failed to listen for incomming streams: %v", err)
	}
	for c := 0; c < 3; c++ {
		sock, err := net.Dial("tcp", "localhost:31415")
		if err != nil {
			t.Errorf("test %d: failed to connect to stream listener: %v", c, err)
		}
		timeout := time.Tick(time.Second)
		select {
		case <-sink:
			continue
		case <-timeout:
			t.Errorf("test %d: listener didn't return incoming stream", c)
		}
		sock.Close()
	}
	close(quit)
}

func TestDial(t *testing.T) {
	sock, err := net.Listen("tcp", "localhost:31416")
	if err != nil {
		t.Errorf("failed to listen for incoming TCP connections: %v", err)
	}
	for c := 0; c < 3; c++ {
		strm, err := Dial("localhost", 31416)
		if err != nil {
			t.Errorf("test %d: failed to connect to TCP listener: %v", c, err)
		}
		strm.Close()
	}
	sock.Close()
}

func TestSendRecv(t *testing.T) {
	sink, quit, err := Listen(31417)
	if err != nil {
		t.Errorf("failed to listen for incomming streams: %v", err)
	}
	c2s, err := Dial("localhost", 31417)
	if err != nil {
		t.Errorf("failed to connect to stream listener: %v", err)
	}
	s2c := <-sink

	var send1 = struct{ A, B int }{3, 14}
	var recv1 = struct{ A, B int }{}

	err = c2s.Send(send1)
	if err != nil {
		t.Errorf("failed to send client -> server: %v", err)
	}
	err = s2c.Recv(&recv1)
	if err != nil {
		t.Errorf("failed to recieve client -> server: %v", err)
	}
	if send1.A != recv1.A || send1.B != recv1.B {
		t.Errorf("sent/received mismatch: have %v, want %v", recv1, send1)
	}

	var send2 = struct{ A, B, C int }{3, 1, 4}
	var recv2 = struct{ A, C int }{}

	err = s2c.Send(send2)
	if err != nil {
		t.Errorf("failed to send client -> server: %v", err)
	}
	err = c2s.Recv(&recv2)
	if err != nil {
		t.Errorf("failed to recieve client -> server: %v", err)
	}
	if send2.A != recv2.A || send2.C != recv2.C {
		t.Errorf("sent/received mismatch: have %v, want %v", recv1, send1)
	}

	c2s.Close()
	s2c.Close()
	close(quit)
}
