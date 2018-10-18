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
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, dni string) {
	bytes := stub.State[dni]

	if bytes == nil {
		fmt.Println("State", dni, "failed to get nombre")
		t.FailNow()
	}
	
}
func checkQuery(t *testing.T, stub *shim.MockStub, dni string, nombre string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte(dni)})
	if res.Status != shim.OK {
		fmt.Println("Query", dni, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", dni, "failed to get nombre")
		t.FailNow()
	}
	var result Person
	json.Unmarshal(res.Payload, &result)
	if result.Name != nombre {
		fmt.Println("Query nombre ", dni, " was not ", nombre, " as expected")
		t.FailNow()
	}
}

func checkQueryAddress(t *testing.T, stub *shim.MockStub, dni string, address string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte(dni)})
	if res.Status != shim.OK {
		fmt.Println("Query", dni, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", dni, "failed to get nombre")
		t.FailNow()
	}
	var result Person
	json.Unmarshal(res.Payload, &result)
	if result.Address != address {
		fmt.Println("Query nombre ", dni, " was not ", address, " as expected")
		t.FailNow()
	}
	fmt.Println("The new address is", result.Address)
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}
func TestExample02_Init(t *testing.T) {
	scc := new(SimpleChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("512712712A"), []byte(`{"Name": "Pepe", "Age": "53", "Address": "Desengano 21"}`)})

	checkState(t, stub, "512712712A")
	checkQuery(t, stub, "512712712A", "Pepe")

	checkInvoke(t, stub, [][]byte{[]byte("modifyAddress"), []byte("512712712A"), []byte("Nueva calle")})
	checkQueryAddress(t, stub, "512712712A", "Nueva calle")
}


