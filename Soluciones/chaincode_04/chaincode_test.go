/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
	fmt.Println("State value", name, "is", value, "as expected")
}

func checkGetBalance(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getBalance"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("GetBalance", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("GetBalance", name, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("GetBalance value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkInitAccount(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init Account", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkWithdraw(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("checkWithdraw", args, "failed", string(res.Message))
		t.FailNow()
	}
}


func checkWithdrawFail(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status == shim.OK {
		fmt.Println("checkWithdrawFail", args, "failed", string(res.Message))
		t.FailNow()
	}
	fmt.Println(res)
}

func checkGetHistory(t *testing.T, stub *shim.MockStub, name string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("getHistory"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("getHistory", name, "failed", string(res.Message))
		t.FailNow()
	}
	fmt.Println(res)
}


func checkSendMoney(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("SendMoney", args, "failed", string(res.Message))
		t.FailNow()
	}
}


func TestBank1_Init(t *testing.T) {
	scc := new(BankChaincode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{})

}


func TestBank1_InitAccount(t *testing.T) {
	scc := new(BankChaincode)
	stub := shim.NewMockStub("ex02", scc)

	checkInit(t, stub, [][]byte{})

	checkInitAccount(t, stub, [][]byte{[]byte("initAccount"), []byte("Maria"), []byte("100")})
	checkState(t, stub, "Maria", "100")
	checkInitAccount(t, stub, [][]byte{[]byte("initAccount"), []byte("Pedro"), []byte("100")})
	checkState(t, stub, "Pedro", "100")

}

func TestBank1_Withdraw(t *testing.T) {
	scc := new(BankChaincode)
	stub := shim.NewMockStub("ex02", scc)

	checkInit(t, stub, [][]byte{})

	// // Invoke A->B for 123
	checkInitAccount(t, stub, [][]byte{[]byte("initAccount"), []byte("Maria"), []byte("100")})
	checkState(t, stub, "Maria", "100")
	checkWithdraw(t, stub, [][]byte{[]byte("withdraw"), []byte("Maria"), []byte("10")})
	checkState(t, stub, "Maria", "90")
	checkWithdrawFail(t, stub, [][]byte{[]byte("withdraw"), []byte("Maria"), []byte("120")})
	checkWithdraw(t, stub, [][]byte{[]byte("withdraw"), []byte("Maria"), []byte("10")})
	checkWithdraw(t, stub, [][]byte{[]byte("withdraw"), []byte("Maria"), []byte("5")})
	checkWithdraw(t, stub, [][]byte{[]byte("withdraw"), []byte("Maria"), []byte("6")})
	checkGetHistory(t, stub, "Maria")
}

