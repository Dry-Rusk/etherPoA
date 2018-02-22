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
	"enode://e896d98f1c03d3b55f6b67063126e7533a612cc979215600263a95348094e3b223e66f8c129380033e3437a1faa81813cb3b9cb9a85ded7061459c05718af405@[159.65.42.209]:1515",
	"enode://cd155d4cbd681d65f04d4018ab5c609502bb0c914ee49b1d19ac6649362cd237681bb3958e5a6bbc6105f56242eaca5fe7b8372604d53cda1418d7e941bc7f92@[159.65.34.218]:2525",
	"enode://f78b9fc73868354d48ca9e11812f4261d836c4a5f2b989bec5f02043fd9ba2c93f8b5e73489e59b55079b524354db078688b29708ecc62615926c6592b0abf0f@[159.65.34.206]:3535",
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
