/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package example02


import (
	"fmt"
	"testing"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {

}

func checkState(t *testing.T, stub *shim.MockStub, dni string) {

	
}
func checkQuery(t *testing.T, stub *shim.MockStub, dni string, nombre string) {
	
}

func checkQueryAddress(t *testing.T, stub *shim.MockStub, dni string, address string) {
	
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	
}
func TestExample02_Init(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)


}


