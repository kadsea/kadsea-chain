// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://87990d4910dd6ad415c58927fb9434f29247dc3d582d73dfc54c170839186b8a96930054a58252e61a7a9e00fab33f13144e9e2175d9f64dc4579bc92fae8876@18.140.65.3:32668",
	"enode://9a31326fa2a286c82374a2cb8a27be93fe79dee7b417b6ee07be5d665bd1091606c2be2acc5984ff442b2d8471575addbde82d302899292d6d9184c6b65d7827@192.168.100.134:32668",
	"enode://6bfae61649991a2516374ff5cf05d75067210fe056e59be1785add0f705fd9dabe2709f14ba5d8bcde7a82f3ae691e1ccbbf3b53618fd866d5ba93438ab2bdce@192.168.100.135:32668",
	"enode://501ebfe6270fe103b1834f6b3b5d5d7af1ec7320d5a1f40613bc9a5e3ef65f0d410dfe0bb20272e3298300e79f7a84657cc3f8199286266b0acd4ebdd086b9f7@192.168.100.133:32668",
	"enode://74d8f5f89ed0e58fd2a26d8fc75299f969f1aceb6e5f4e2a87fac8acc26173983b8736d4bdeda07c1085a6c289dd3385dca6fcdde01ae0e9e36065062fa4f166@112.213.84.68:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://924543a43d18bc5759a8bdcd17fa9c7c35df63968e9333640b80b58dab94b17a012371c9d46bed10ce7508a607cac76828ca04685893958eee44ade83b856dc2@47.242.237.63:32668",
	"enode://ebad898d980b520ef6adb54ffb6a68117686e7332f1ea01f7551b7a296a34dd945445a078d7cad019d864c5ef0e0b7f2b5777d94f93adf7dc59f798af72609ac@47.242.235.121:32668",
}

var V5Bootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
