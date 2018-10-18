/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BankChaincode struct {
}
var logger = shim.NewLogger("Bank Chaincode")
func (t *BankChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	level, err := shim.LogLevel("DEBUG")
	if err != nil {
		return shim.Error("Problems with loggin level")
	}
	logger.SetLevel(shim.LoggingLevel(level))

	logger.Debug("[Bank Chaincode][Init]Instanciating chaincode...")
	return shim.Success(nil)
}

func (t *BankChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	logger.Debug("[Bank Chaincode][Invoke]Invoking chaincode..." + function)

	if function == "initAccount" {
		// Make payment of X units from A to B
		return t.initAccount(stub, args)
	} else if function == "withdraw" {
		// Deletes an entity from its state
		return t.withdraw(stub, args)
	} else if function == "getHistory" {
		// the old "Query" is now implemtned in invoke
		return t.getHistory(stub, args)
	} else if function == "getBalance" {
		// the old "Query" is now implemtned in invoke
		return t.getBalance(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Init account
func (t *BankChaincode) initAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var account string    // Entities
	var balance int // Asset holdings
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	account = args[0]
	balance, err = strconv.Atoi(args[1])
	
	err = stub.PutState(account, []byte(strconv.Itoa(balance)))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// withdraw money
func (t *BankChaincode) withdraw(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var balance int

	account := args[0]
	ammount, err := strconv.Atoi(args[1])

	// Delete the key from the state in ledger
	balanceBytes, err := stub.GetState(account)
	if err != nil {
		return shim.Error("Failed get state")
	}

	balance, _ = strconv.Atoi(string(balanceBytes))

	if(balance < ammount){
		return shim.Error("Insuficient funds")
	}
	balance = balance - ammount

	// Update de state
	
	err = stub.PutState(account, []byte (strconv.Itoa(balance)))
	if err != nil {
		return shim.Error("Fail updating balance")
	}
	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *BankChaincode) getBalance(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var account string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	account = args[0]

	// Get the state from the ledger
	balanceBytes, err := stub.GetState(account)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + account + "\"}"
		return shim.Error(jsonResp)
	}

	if balanceBytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + account + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + account + "\",\"Amount\":\"" + string(balanceBytes) + "\"}"
	logger.Debug("Query Response:%s\n", jsonResp)
	return shim.Success(balanceBytes)
}

func (t *BankChaincode) getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	account := args[0]
	historyIer, err := stub.GetHistoryForKey(account)
	defer historyIer.Close()
	if err != nil {
		fmt.Println(err)
		return shim.Error("Error getting history")
	}

	for historyIer.HasNext() {
		modification, err := historyIer.Next()
		if err != nil {
			fmt.Println(err)
			return shim.Error("Error getting history")
		}
		fmt.Println("Returning information about", string(modification.Value))
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(BankChaincode))
	if err != nil {
		fmt.Printf("Error starting Bank chaincode: %s", err)
	}
}