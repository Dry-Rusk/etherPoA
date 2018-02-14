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

package misc

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/Exgibichi/go-etf/core/state"
	"github.com/Exgibichi/go-etf/core/types"
	"github.com/Exgibichi/go-etf/params"
)

var (
	// ErrBadProDAOExtra is returned if a header doens't support the DAO fork on a
	// pro-fork client.
	ErrBadProPOAExtra = errors.New("bad POA pro-fork extra-data")

	// ErrBadNoDAOExtra is returned if a header does support the DAO fork on a no-
	// fork client.
	ErrBadNoPOAExtra = errors.New("bad POA no-fork extra-data")
)

// VerifyDAOHeaderExtraData validates the extra-data field of a block header to
// ensure it conforms to DAO hard-fork rules.
//
// DAO hard-fork extension to the header validity:
//   a) if the node is no-fork, do not accept blocks in the [fork, fork+10) range
//      with the fork specific extra-data set
//   b) if the node is pro-fork, require blocks in the specific range to have the
//      unique extra-data set.
func VerifyPOAHeaderExtraData(config *params.ChainConfig, header *types.Header) error {
	// Short circuit validation if the node doesn't care about the DAO fork
	// if config.PoaBlock == nil {
	// 	return nil
	// }
	// // Make sure the block is within the fork's modified extra-data range
	// limit := new(big.Int).Add(config.PoaBlock, params.PoaExtraRange)
	// if header.Number.Cmp(config.PoaBlock) < 0 || header.Number.Cmp(limit) >= 0 {
	// 	return nil
	// }
	// // Depending on whether we support or oppose the fork, validate the extra-data contents
	// if config.PoaSupport {
	// 	if !bytes.Equal(header.Extra, params.PoaBlockExtra) {
	// 		return ErrBadProPOAExtra
	// 	}
	// } else {
	// 	if bytes.Equal(header.Extra, params.PoaBlockExtra) {
	// 		return ErrBadNoPOAExtra
	// 	}
	// }
	// // All ok, header has the same extra-data we expect
	return nil
}

// ApplyDAOHardFork modifies the state database according to the DAO hard-fork
// rules, transferring all balances of a set of DAO accounts to a single refund
// contract.
func ApplyPOAHardFork(statedb *state.StateDB) { // , header *types.Header
	var balance *big.Int = new(big.Int).Mul(big.NewInt(1e+18), big.NewInt(5e+6))
	fmt.Printf("apply POA\n")
	// debug.PrintStack()
	if statedb != nil {
		for _, addr := range params.POA() {
			if !statedb.Exist(addr) {
				statedb.CreateAccount(addr)
				//statedb.AddBalance(addr, balance)
				statedb.SetBalance(addr, balance)
			}
		}
	}
	// if header != nil {
	// 	if !bytes.Equal(header.Extra, params.POAForkBlockExtra) {
	// 		header.Extra = common.CopyBytes(params.POAForkBlockExtra)
	// 	}
	// 	if header.GasLimit != params.POAForkGasLimit && !bytes.Equal(header.Extra, params.POAForkBlockExtra) {
	// 		header.GasLimit = params.POAForkGasLimit
	// 	}
	// }
}
