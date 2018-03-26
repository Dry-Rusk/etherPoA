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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// ETF Go Bootnodes
	"enode://e646d31323e6f74162bf7a6f6c43663a0b6047d7dc0d9744aa10d5913682f6c3f5341840238a62fd1e55511132fd8c5932333da4e96a06ab6ed322b323ec9589@[62.210.13.4]:30301",
	"enode://9e157f417e8f743868713ed046598ce82ef809927b188bfc13e350f95a8fb2dbbaa4a092a2a9d099d0a9531fa7b7fcdc982548a462c25dfe51fa8b734c67f967@[62.210.27.177]:30301",
	"enode://2a132a6cb83d607401941e334fe334faf1a3cae77b7488bf791c6a8bdc40dbc2bec4b5995e8a378214fbabba6d0b637cfcc4a6dfa05aaa4d46148688a187fb88@[62.210.29.177]:30301",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
