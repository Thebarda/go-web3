/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file web3-gasprice_test.go
 * @authors:
 *   Tom Darneix <tomdar87@outlook.com>
 * @date 2018
 */

package test

import (
	"math/big"
	"testing"

	web3 "github.com/regcostajr/go-web3"
	"github.com/regcostajr/go-web3/complex/types"
	"github.com/regcostajr/go-web3/dto"
	"github.com/regcostajr/go-web3/providers"
)

func TestEthGetFilterChanges(t *testing.T) {

	var connection = web3.NewWeb3(providers.NewHTTPProvider("vps409515.ovh.net:42669", 10, false))

	coinbase, err := connection.Eth.GetCoinbase()
	transaction := new(dto.TransactionParameters)
	transaction.From = coinbase
	transaction.To = coinbase
	transaction.Value = big.NewInt(10)
	transaction.Gas = big.NewInt(40000)
	transaction.Data = types.ComplexString("p2p transaction")
	_, errSend := connection.Eth.SendTransaction(transaction)

	if errSend != nil {
		t.Error(errSend)
		t.FailNow()
	}

	logs, err := connection.Eth.GetFilterChanges(coinbase)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log("coucou")
	t.Log(logs)
}
