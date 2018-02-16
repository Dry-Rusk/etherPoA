// Copyright 2018 The go-etf Authors
// This file is part of the go-etf library.
//
// The go-etf library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-etf library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-etf library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/Exgibichi/go-etf/common"
)

// POA fork params
var (
	POAForkBlockExtra        = common.FromHex("0x796f62610000000000000000000000000000000000000000000000000000000001a4260e6348ddd60ee1a597d1c6a102a076e9040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	POAForkExtraRange        = big.NewInt(10)
	POAConfig                = &CliqueConfig{Period: 5, Epoch: 30000}
	POAForkGasLimit   uint64 = 105e+6
	POAChainID               = big.NewInt(455446)
	POAFoundersReward        = new(big.Int).Mul(big.NewInt(1e+18), big.NewInt(10e+6))
)

// POA Founders accounts
func POA() []common.Address {
	return []common.Address{
		common.HexToAddress("0xc3e980ccdbeed1975ceb89018470ea3cc087548f"),
		common.HexToAddress("0x674903fb269ce2e395b85746702455712af7913e"),
		common.HexToAddress("0x46c595f6f942bf06c7cb8d388cabb3bba2678a3e"),
		common.HexToAddress("0x9171f8f7e745ae7da611decbc2975eacc05be51c"),
		common.HexToAddress("0x44d33c88771c9911c7ed9b3f75bf69218ae1248e"),
	}
}
