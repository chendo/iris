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

// Package config contains the hard-coded configuration values of the system.
package config

import (
	"crypto"
	"crypto/aes"
	"crypto/md5"
	"math/big"
	"time"
)

// Cyclic group for the STS cryptography (2448 bits).
var StsGroup = new(big.Int).SetBytes([]byte{
	0xed, 0xeb, 0xe6, 0x33, 0xb2, 0x8e, 0xa8, 0xb0,
	0x69, 0x57, 0xa1, 0x69, 0xd6, 0x9a, 0x5e, 0x09,
	0x80, 0xd4, 0x91, 0x78, 0xfe, 0xdf, 0x9c, 0x5e,
	0xc6, 0xfa, 0x7d, 0xd9, 0x37, 0xc3, 0x26, 0xe9,
	0xe2, 0xe0, 0xe5, 0x50, 0x11, 0xe4, 0x23, 0xa9,
	0x5e, 0x33, 0x4d, 0x43, 0x2d, 0x4f, 0x08, 0x10,
	0xb4, 0xc6, 0xc5, 0x59, 0x42, 0xb3, 0x26, 0x72,
	0x6e, 0xe8, 0xac, 0xce, 0xfb, 0xdc, 0xcd, 0x79,
	0x5f, 0x75, 0xa1, 0x14, 0x38, 0x25, 0x6a, 0x70,
	0xf9, 0x86, 0xf6, 0x79, 0x9f, 0x60, 0xc7, 0xee,
	0xc2, 0x56, 0x15, 0xd0, 0xf2, 0x95, 0x18, 0xdb,
	0xf1, 0x90, 0x56, 0xdc, 0x96, 0x79, 0xb1, 0xd6,
	0x42, 0x1e, 0xcb, 0xe1, 0x84, 0x59, 0x57, 0xbb,
	0xf3, 0x97, 0x5f, 0x64, 0xbf, 0x7f, 0x27, 0xc8,
	0xe1, 0x8c, 0xfa, 0x58, 0x19, 0x07, 0x28, 0x2b,
	0x2a, 0xb0, 0x03, 0x32, 0xac, 0x26, 0x13, 0xcd,
	0xf1, 0x47, 0x29, 0xab, 0x9a, 0x08, 0x56, 0x61,
	0x7e, 0xbe, 0x2a, 0x79, 0x9a, 0x22, 0x0e, 0x03,
	0xf1, 0x38, 0x1d, 0x9e, 0x0e, 0xc3, 0x35, 0xe1,
	0x3c, 0x39, 0xeb, 0xfd, 0x4d, 0x23, 0x4a, 0x9f,
	0xa3, 0xc0, 0xed, 0x66, 0xab, 0xfd, 0x6b, 0x4e,
	0x0f, 0x4d, 0x6a, 0xf4, 0x58, 0x66, 0xda, 0xe8,
	0x20, 0xa4, 0x95, 0xa1, 0xb2, 0xc1, 0x8c, 0xfa,
	0x49, 0x9c, 0xb8, 0xed, 0x21, 0x2d, 0x88, 0x7b,
	0x52, 0x0a, 0xac, 0xbd, 0xf6, 0xf4, 0x0b, 0x76,
	0x87, 0x0b, 0xcb, 0xa5, 0x0f, 0xa0, 0xd0, 0x0f,
	0xc1, 0x16, 0x7c, 0x86, 0x38, 0x3c, 0x16, 0x7d,
	0x40, 0xf4, 0x4a, 0xa0, 0x9c, 0x23, 0x2b, 0x6a,
	0x14, 0xe1, 0xff, 0x42, 0xdf, 0x42, 0x3f, 0x58,
	0xc8, 0xa6, 0xc8, 0xec, 0x03, 0xdf, 0xab, 0xd6,
	0x05, 0x74, 0x81, 0x63, 0x03, 0x9e, 0x6a, 0xb8,
	0x29, 0x78, 0xa7, 0xe4, 0x8c, 0x78, 0xb7, 0x42,
	0x6b, 0xe8, 0x99, 0x69, 0x41, 0x6c, 0xce, 0x50,
	0x31, 0x05, 0xca, 0xc1, 0xe4, 0x77, 0x24, 0xd5,
	0x8d, 0xe8, 0xd3, 0x6b, 0xd2, 0xfd, 0xd5, 0xd9,
	0xa2, 0x1d, 0xda, 0xf1, 0xc7, 0x11, 0x2c, 0xe6,
	0xc9, 0x9f, 0x1d, 0x7d, 0x62, 0x94, 0x62, 0x43,
	0x86, 0x9b, 0x7c, 0x2e, 0x77, 0x83, 0x81, 0xb7,
	0xeb, 0x33,
})

// Cyclic group generator for the STS cryptography (2448 bits).
var StsGenerator = new(big.Int).SetBytes([]byte{
	0x9b, 0xe1, 0x3f, 0x9c, 0xb0, 0x3a, 0x34, 0x18,
	0x50, 0x1d, 0xdf, 0xe2, 0xdf, 0x25, 0x4b, 0x13,
	0x6a, 0x1a, 0x0e, 0xf9, 0xbd, 0x2b, 0x1a, 0x06,
	0x54, 0xd3, 0x2d, 0x86, 0x41, 0xb7, 0x66, 0xca,
	0x8c, 0x76, 0x43, 0x65, 0xfe, 0xf8, 0xc1, 0xb2,
	0xce, 0x5a, 0x44, 0xfd, 0x99, 0x18, 0x82, 0x5a,
	0x35, 0x7e, 0xa8, 0xe0, 0xad, 0xaa, 0x28, 0x7e,
	0x19, 0x2d, 0x02, 0x86, 0x9f, 0xcf, 0x6e, 0xc5,
	0x84, 0x62, 0x60, 0x3b, 0x6a, 0xe3, 0x57, 0x74,
	0xcf, 0x6e, 0x5a, 0x89, 0xd6, 0xc1, 0x38, 0x03,
	0x66, 0x5a, 0xbb, 0x52, 0x11, 0xd4, 0xe9, 0x3d,
	0x73, 0x83, 0xb3, 0x17, 0xd3, 0x0f, 0x2e, 0x8b,
	0x16, 0xe1, 0xd9, 0x91, 0x97, 0x2f, 0xf8, 0x53,
	0x2a, 0x6a, 0xf5, 0x83, 0xe3, 0x1b, 0x39, 0x69,
	0x3e, 0xe6, 0x68, 0x01, 0xe9, 0x84, 0xdf, 0x76,
	0x8e, 0x96, 0xf1, 0x11, 0x19, 0x45, 0x03, 0x18,
	0x2d, 0x68, 0x29, 0xc5, 0xcc, 0x0b, 0x20, 0x2c,
	0x1f, 0xaf, 0x51, 0x97, 0xd3, 0x2c, 0xdb, 0xc9,
	0xfd, 0x98, 0x4d, 0x95, 0xb8, 0x13, 0x5f, 0xe4,
	0x51, 0x83, 0x1f, 0x15, 0x0c, 0xd9, 0x84, 0x42,
	0x57, 0x30, 0xc6, 0x46, 0x0f, 0x3f, 0x19, 0x44,
	0xd8, 0x77, 0xf1, 0x54, 0x22, 0x36, 0xa7, 0x4f,
	0x32, 0x5f, 0xe2, 0xe5, 0xf8, 0xd3, 0x55, 0x84,
	0x85, 0x5a, 0x9f, 0x15, 0x16, 0x14, 0x16, 0x7a,
	0x89, 0x18, 0x9e, 0x20, 0x34, 0x55, 0x50, 0xaf,
	0x21, 0x21, 0xff, 0xa6, 0xde, 0x0d, 0xc8, 0xad,
	0x72, 0xb7, 0x56, 0x6a, 0xae, 0x6c, 0xff, 0x79,
	0x10, 0xf3, 0x47, 0xd7, 0xee, 0xaf, 0x4f, 0x82,
	0x3b, 0x22, 0x93, 0xd8, 0x52, 0xba, 0x37, 0x9e,
	0xf3, 0x6a, 0x7d, 0xa2, 0x67, 0x61, 0x9a, 0xb8,
	0xcc, 0xd5, 0xf6, 0x23, 0xeb, 0x08, 0x87, 0x12,
	0xf3, 0x88, 0xaa, 0x38, 0xaf, 0xcc, 0x03, 0xf2,
	0xe5, 0xf8, 0x5c, 0x52, 0xc4, 0x08, 0x49, 0xd2,
	0xcf, 0x75, 0x58, 0xe3, 0xf4, 0x00, 0x99, 0x89,
	0x75, 0x87, 0x1a, 0x9f, 0x7b, 0x60, 0x9e, 0x39,
	0xb4, 0x89, 0xf5, 0xdf, 0x39, 0x49, 0xa8, 0xf3,
	0x41, 0x95, 0x8e, 0x67, 0x19, 0x78, 0x34, 0x50,
	0x9d, 0x1e, 0x52, 0xd5, 0x31, 0x6a, 0x3d, 0xf0,
	0xaa, 0xa0,
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

// Maximum allowed time to complete a session connection.
var SessionDialTimeout = time.Second

// Maximum allowed time to handle a session connection.
var SessionAcceptTimeout = time.Second

// Maximum allowed time to complete the session control channel setup.
var SessionShakeTimeout = 3 * time.Second

// Maximum allowed time to complete the session data channel setup.
var SessionLinkTimeout = time.Second

// Time allowance to gracefully terminate a session link.
var SessionGraceTimeout = 3 * time.Second

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
var BootScan = 100

// Virtual address space (bits).
var PastrySpace = 40

// Number of matching bits for the next hop.
var PastryBase = 4

// Number of closest nodes to track in the virtual network.
var PastryLeaves = 8

// Hash for mapping external ids into the overlay id space.
var PastryResolver = md5.New

// Time after booting to consider the overlay a single node in the network.
var PastryBootTimeout = 10 * time.Second

// Idle time after which to consider the overlay converged.
var PastryConvTimeout = 3 * time.Second

// Heartbeat period to ensure connections are alive and tear down unused ones.
var PastryBeatPeriod = 3 * time.Second

// Number of missed heartbeats after which to consider a node down.
var PastryKillCount = 3

// Maximum time to queue an authenticated session connection before dropping it.
var PastryAcceptTimeout = time.Second

// Time to wait after session setup for the init packet.
var PastryInitTimeout = 5 * time.Second

// Time limit for sending a message before the connection is dropped.
var PastrySendTimeout = 3 * time.Second

// Messages to buffer to and from the network.
var PastryNetBuffer = 64

// Maximum number of authentications allowed concurrently (per half duplex).
var PastryAuthThreads = 8

// Maximum number of state exchanges allowed concurrently.
var PastryExchThreads = 128

// Heartbeat period to distribute current CPU load and also check liveliness (ms).
var ScribeBeatPeriod = time.Second

// Number of missed heartbeats after which to consider a node down.
var ScribeKillCount = 3

// Application identifier space (bits).
var ScribeSpace = 32

// Number of messages to buffer for application delivery before dropping.
var ScribeAppBuffer = 128

// Number of sub-clusters an app cluster or topic is split into.
var IrisClusterSplits = 5

// Maximum number of handlers allowed concurrently per Iris application.
var IrisHandlerThreads = 16

// Maximum time to queue an established tunnel stream before dropping it.
var IrisTunnelAcceptTimeout = time.Second

// Maximum time to wait for a client init packet.
var IrisTunnelInitTimeout = time.Second

// Send and receive window for tunnel ordering and throttling.
var IrisTunnelBuffer = 256

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
