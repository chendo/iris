// Iris - Decentralized cloud messaging
// Copyright (c) 2013 Project Iris. All rights reserved.
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

package session

import (
	"crypto/rand"
	"crypto/rsa"
	"net"
	"testing"
	"time"
)

// Tests whether the session handshake works.
func TestHandshake(t *testing.T) {
	t.Parallel()

	addr, _ := net.ResolveTCPAddr("tcp", "localhost:0")
	key, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Start the server
	sock, err := Listen(addr, key)
	if err != nil {
		t.Fatalf("failed to start the session listener: %v.", err)
	}
	sock.Accept(10 * time.Millisecond)

	// Connect with a few clients, verifying the crypto primitives
	for i := 0; i < 3; i++ {
		client, err := Dial("localhost", addr.Port, key)
		if err != nil {
			t.Fatalf("failed to connect to the server: %v.", err)
		}
		// Make sure the server also gets back a live session
		select {
		case server := <-sock.Sink:
			// Close the two sessions
			if err := client.Close(); err != nil {
				t.Fatalf("failed to close client session: %v.", err)
			}
			if err := server.Close(); err != nil {
				t.Fatalf("failed to close server session: %v.", err)
			}

		case <-time.After(10 * time.Millisecond):
			t.Fatalf("server-side handshake timed out.")
		}
	}
	// Ensure the listener can be torn down correctly
	if err := sock.Close(); err != nil {
		t.Fatalf("failed to terminate session listener: %v.", err)
	}
}

// Benchmarks the session setup performance.
func BenchmarkHandshake(b *testing.B) {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:0")
	key, _ := rsa.GenerateKey(rand.Reader, 2048)

	sock, err := Listen(addr, key)
	if err != nil {
		b.Fatalf("failed to start the session listener: %v.", err)
	}
	sock.Accept(10 * time.Millisecond)

	// Collectors for the established sessions
	sink := make(chan *Session)
	dump := make([]*Session, 0)

	// Execute the handshake benchmarks
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Start a dialer on a new thread
		go func() {
			sess, err := Dial("localhost", addr.Port, key)
			if err != nil {
				b.Fatalf("failed to connect to the server: %v.", err)
				close(sink)
			} else {
				sink <- sess
			}
		}()
		// Wait for the negotiated session from both client and server side
		client, ok := <-sink
		if !ok {
			b.Fatalf("client negotiation failed.")
		}
		dump = append(dump, client)

		select {
		case server := <-sock.Sink:
			dump = append(dump, server)
		case <-time.After(10 * time.Millisecond):
			b.Fatalf("server-side handshake timed out.")
		}
	}
	b.StopTimer()

	// Clean up the established sessions
	for _, sess := range dump {
		if err := sess.Close(); err != nil {
			b.Fatalf("failed to close session: %v.", err)
		}
	}
	// Tear down the listener
	if err := sock.Close(); err != nil {
		b.Fatalf("failed to terminate session listener: %v.", err)
	}
}
