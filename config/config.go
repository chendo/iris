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
package config

import (
	"crypto"
	"crypto/aes"
	"crypto/md5"
	"math/big"
)

// Cyclic group for the STS cryptography (2248 bits).
var StsGroup = new(big.Int).SetBytes([]byte{
	0xdc, 0x28, 0x29, 0xab, 0xca, 0xc5, 0x7d, 0x0d,
	0xf7, 0x44, 0xa4, 0x9a, 0x42, 0x7e, 0x5b, 0xe9,
	0xa7, 0xf8, 0xd3, 0x3f, 0x87, 0x01, 0xfa, 0x37,
	0x3d, 0xfe, 0x1b, 0x31, 0xec, 0x03, 0x48, 0x9f,
	0x77, 0xe3, 0x2f, 0xc1, 0x8b, 0xc2, 0x3a, 0xa5,
	0x95, 0x2f, 0x19, 0x04, 0x76, 0xba, 0xe7, 0xef,
	0xeb, 0x80, 0xd7, 0xf8, 0x72, 0xca, 0x34, 0xfe,
	0x88, 0xb5, 0x28, 0x0e, 0x41, 0x33, 0x16, 0x8d,
	0xee, 0x27, 0x4b, 0x0a, 0xf1, 0x9e, 0xfa, 0xe4,
	0xf0, 0xed, 0x86, 0x22, 0x8d, 0xd8, 0xa3, 0x9f,
	0x61, 0xd8, 0xaf, 0x77, 0xb1, 0x9d, 0xf8, 0x2d,
	0x3b, 0x5d, 0x3f, 0x49, 0xb4, 0xe3, 0x9c, 0xb8,
	0xeb, 0xa5, 0x32, 0xf4, 0xa8, 0xf9, 0x48, 0x5b,
	0x6d, 0xac, 0xee, 0x4e, 0xd5, 0xe6, 0x81, 0x1e,
	0xfd, 0x60, 0x43, 0x28, 0xd3, 0x4b, 0xd8, 0xca,
	0x52, 0xf7, 0x3f, 0x5e, 0xfc, 0x80, 0x11, 0x9d,
	0x74, 0x58, 0x8c, 0x83, 0x1f, 0x0f, 0x1e, 0x0e,
	0xd6, 0x0e, 0xe8, 0xc5, 0x72, 0x1d, 0x8f, 0x0e,
	0x4e, 0x14, 0x45, 0xfa, 0x46, 0x1e, 0xa9, 0xf8,
	0x67, 0xd8, 0x02, 0xfa, 0x88, 0x35, 0xe5, 0x39,
	0xf9, 0xa6, 0x09, 0xba, 0xda, 0x7f, 0x78, 0x72,
	0x0b, 0x14, 0xd1, 0xef, 0xff, 0x70, 0xfd, 0x05,
	0x62, 0x7c, 0x93, 0xde, 0x22, 0x17, 0x8f, 0xe1,
	0xab, 0x37, 0x9c, 0xc5, 0xa4, 0xab, 0x10, 0x4c,
	0x1d, 0xf0, 0xc3, 0xa7, 0xd3, 0xad, 0x9f, 0x97,
	0xd9, 0xea, 0xd9, 0xe4, 0x1a, 0xbd, 0xfe, 0x84,
	0x9b, 0x72, 0xec, 0x27, 0xf3, 0xd5, 0x83, 0x39,
	0x70, 0x19, 0x23, 0xcc, 0xd9, 0x51, 0x1e, 0xb2,
	0x9d, 0x3f, 0x38, 0x64, 0x04, 0x36, 0x13, 0xcc,
	0xbc, 0xb8, 0x62, 0xcb, 0x1e, 0xbf, 0x30, 0x08,
	0x2f, 0xe5, 0xca, 0xdc, 0x8a, 0xb5, 0xd7, 0x91,
	0x0f, 0x60, 0x99, 0x1d, 0x0b, 0x3a, 0x70, 0x16,
	0x59, 0x42, 0x4a, 0x5d, 0xde, 0x5d, 0x10, 0x5b,
	0xbc, 0x30, 0x60, 0xb9, 0x59, 0x37, 0xf2, 0xe8,
	0x50, 0xa3, 0x68, 0x02, 0x15, 0x27, 0xc4, 0xee,
	0x53,
})

// Cyclic group generator for the STS cryptography (2248 bits).
var StsGenerator = new(big.Int).SetBytes([]byte{
	0x09, 0x50, 0x1e, 0x53, 0xeb, 0xce, 0xd4, 0xc8,
	0x05, 0x0d, 0x76, 0x90, 0xee, 0xf5, 0x48, 0x06,
	0x18, 0xca, 0xd2, 0x9e, 0x75, 0x37, 0x9d, 0x0b,
	0x7f, 0x6f, 0x47, 0xe0, 0xe9, 0xe8, 0xd1, 0xd0,
	0x16, 0xbd, 0xf1, 0xa8, 0xc2, 0x73, 0x19, 0x93,
	0xa4, 0xf3, 0x42, 0x58, 0x8c, 0x4e, 0x7b, 0x8b,
	0x62, 0xa5, 0x23, 0xc1, 0xe6, 0xec, 0x89, 0xa5,
	0xdc, 0x49, 0xa4, 0xcd, 0xb7, 0x54, 0xfc, 0xba,
	0x32, 0xef, 0x14, 0x16, 0xc3, 0x3b, 0xb0, 0xcc,
	0xfc, 0xe4, 0x81, 0xd2, 0x3d, 0x16, 0x79, 0x3a,
	0x46, 0xaf, 0x1e, 0xd3, 0x2a, 0x97, 0x7a, 0xb4,
	0xfa, 0x91, 0x0f, 0x64, 0x8b, 0x56, 0x65, 0xce,
	0xe0, 0x97, 0x09, 0xf6, 0xf0, 0x91, 0x26, 0x63,
	0xa2, 0x27, 0xd0, 0x15, 0xf2, 0xd0, 0x56, 0x4b,
	0x08, 0xcc, 0xeb, 0x4e, 0x84, 0xba, 0xdb, 0x33,
	0x17, 0x2b, 0xe9, 0xbb, 0xfa, 0xa4, 0x50, 0xb7,
	0x80, 0x9d, 0xd6, 0x96, 0xb2, 0xfc, 0xcb, 0x5c,
	0x35, 0xee, 0xa7, 0x3a, 0x2a, 0xd5, 0x0d, 0xeb,
	0x3d, 0xbb, 0xde, 0x21, 0x2a, 0x39, 0xfa, 0x2a,
	0x55, 0x4b, 0xf4, 0x8e, 0x8e, 0x99, 0xca, 0xae,
	0x44, 0x72, 0x55, 0x90, 0xb9, 0xe4, 0xc6, 0x8b,
	0x14, 0x2d, 0xf7, 0x3e, 0x77, 0xf3, 0x7b, 0x2f,
	0xcc, 0x69, 0xb1, 0x2c, 0xb6, 0x2c, 0xba, 0x46,
	0x47, 0xa7, 0xc3, 0x2f, 0xbf, 0x37, 0xe7, 0x80,
	0x4d, 0xe9, 0x0e, 0x92, 0xc9, 0x57, 0x08, 0x8a,
	0x0a, 0x37, 0x6f, 0xde, 0xf8, 0xa7, 0xf9, 0xa3,
	0x3a, 0xdf, 0x45, 0x0d, 0x3c, 0xde, 0xbe, 0x3a,
	0x14, 0x8e, 0xd2, 0x3b, 0xfc, 0x20, 0xfd, 0xf9,
	0xe6, 0x3d, 0x43, 0x5a, 0xb8, 0x4d, 0xef, 0xf4,
	0x23, 0x02, 0x77, 0x9d, 0x3a, 0xfa, 0xba, 0xee,
	0x97, 0xbe, 0x15, 0x94, 0xcc, 0xa3, 0x69, 0x0b,
	0x6c, 0x95, 0xcc, 0x5c, 0xb2, 0x40, 0x40, 0x1d,
	0x7e, 0xa7, 0x9a, 0xe5, 0x4e, 0x76, 0x92, 0xd1,
	0x3d, 0x91, 0x9e, 0x24, 0xde, 0xbb, 0x03, 0x8d,
	0x71, 0x7f, 0x1d, 0xbb, 0xe5, 0xd9, 0x78, 0xbb,
	0x96,
})

// Symmetric cipher to use for the STS encryption.
var StsCipher = aes.NewCipher

// Key size for the symmetric cipher (bits).
var StsCipherBits = 128

// Hash type for the RSA signature/verification.
var StsSigHash = crypto.MD5

// Hash type for the HMAC within HKDF.
var HkdfHash = crypto.MD5

// Salt value for the HKDF key extraction.
var HkdfSalt = []byte("iris.proto.session.hkdf.salt")

// Info value for the HKDF key expansion.
var HkdfInfo = []byte("iris.proto.session.hkdf.info")

// Symmetric cipher to use for session encryption.
var SessionCipher = aes.NewCipher

// Key size for the session symmetric cipher (bits).
var SessionCipherBits = 128

// Hash creator for the session HMAC.
var SessionHash = md5.New

// Symmetric cipher for the temporary message encryption.
var PacketCipher = aes.NewCipher

// Key size for the temporary cipher (bits).
var PacketCipherBits = 128

// Bootstrapping ports to use.
var BootPorts = []int{14142, 27182, 31415, 45654, 22222, 33333}

// Number of heartbeats to queue before blocking.
var BootBeatsBuffer = 32

// Probing interval during bootstrapping in startup mode (ms).
var BootFastProbe = 250

// Probing interval during bootstrapping in maintenance mode (ms).
var BootSlowProbe = 1000

// Scanning interval during bootstrapping (ms).
var BootScan = 250

// Virtual address space (bits).
var OverlaySpace = 40

// Number of matching bits for the next hop.
var OverlayBase = 4

// Number of closest nodes to track in the virtual network.
var OverlayLeaves = 8

// Number of closes nodes to track in the real network.
var OverlayNeighbors = 8

// Hash for mapping external ids into the overlay id space.
var OverlayResolver = md5.New

// Heartbeat period to ensure connections are alive and tear down unused ones (ms).
var OverlayBeatPeriod = 10000

// Time to wait after session setup for the init packet (ms).
var OverlayInitTimeout = 5000

// Time limit for sending a message before the connection is dropped (ms).
var OverlaySendTimeout = 10000

// Messages to buffer inside the overlay going out to one peer.
var OverlayNetPreBuffer = 64

// Messages to buffer to and from the network.
var OverlayNetBuffer = 64

// Maximum number of authentications allowed concurrently.
var OverlayAuthThreads = 8

// Maximum number of state exchanges allowed concurrently.
var OverlayExchThreads = 128

// Heartbeat period to distribute current cpu load and also check liveliness (ms).
var CarrierBeatPeriod = 500

// Number of missed heartbeats after which to consider a node down.
var CarrierKillCount = 3

// Application identifier space (bits).
var CarrierSpace = 32

// Number of messages to buffer for application delivery before dropping.
var CarrierAppBuffer = 128

// Number of sub-clusters an app cluster or topic is split into.
var IrisClusterSplits = 5

// Maximum number of handlers allowed concurrently per Iris application.
var IrisHandlerThreads = 16

// Send and receive window for tunnel ordering and throttling.
var IrisTunnelWindow = 256

// Timeout value for receiving a new acknowledgement.
var IrisTunnelTimeout = 3000

// Use in case of federated applications.
var AppParentId = []byte(nil)

// Protocol version to ensure compatible connections.
var ProtocolVersion = "v0.1-pre"

// Maximum number of handlers allowed concurrently per relay connection.
var RelayHandlerThreads = 8

// Number of messages to buffer per outbound tunnel.
var RelayTunnelBuffer = 128

// Time alloted to a client to acknowledge a tunnel.
var RelayTunnelTimeout = 3000

// Block time when trying a tunnel read (ms).
var RelayTunnelPoll = 1000
