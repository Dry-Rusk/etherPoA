// Copyright 2016 The go-ethereum Authors
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

import (
	"math/big"

	"github.com/Exgibichi/go-etf/common"
)

// DAOForkBlockExtra is the block header extra-data field to set for the DAO fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("dao-hard-fork").
var (
	POAForkBlockExtra        = common.FromHex("0x796f62610000000000000000000000000000000000000000000000000000000001a4260e6348ddd60ee1a597d1c6a102a076e9040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	POAForkExtraRange        = big.NewInt(10)
	POABlock                 = MainnetChainConfig.POABlock
	POAForkGasLimit   uint64 = 105e+6
	POAForkDiff              = big.NewInt(2)
	POAEngine                = &CliqueConfig{
		Period: 5,
		Epoch:  30000,
	}
)

// DAODrainList is the list of accounts whose full balances will be moved into a
// refund contract at the beginning of the dao-fork block.
func POA() []common.Address {
	return []common.Address{
		common.HexToAddress("0x01a4260e6348ddd60ee1a597d1c6a102a076e904"),
		common.HexToAddress("0x13678c1285ff6b9b84e566d2e374e7d1be45fc8f"),
		common.HexToAddress("0x939f2f2fc4f1650a28733e3e52d844be12ce7fa7"),
		common.HexToAddress("0x983e7d95b9629ee41e44592fba8ebdf72bb5260c"),
		common.HexToAddress("0x69b565fdacb27c7ae652e3aa13322347e5398630"),
		common.HexToAddress("0xe939ec7525f8787f8c0ec4eef25be6e2b5dd34f0"),
		common.HexToAddress("0x8986578a8d9951a4B55038B01262d81b2a030D33"),
	}
}
