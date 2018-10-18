/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package example02

import (
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
	
}

type Person struct {
	Name string  `json:"Name"`
	Age string `json:"Age"`
	Address string `json:"Address"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var dni string    // value
	
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	rawIn := json.RawMessage(args[1])
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		fmt.Printf("[Management Chaincode][StoreCode]Error Marshaling Object")
		return shim.Error(err.Error())
	}
	fmt.Println("Argumnts received to store " + args[1])
	var persona Person
	err = json.Unmarshal(bytes, &persona)
	if err != nil {
		fmt.Printf("[Management Chaincode][StoreCode]Error Marshaling Object")
		return shim.Error(err.Error())
	}
	dni = args[0]
	fmt.Println("Person to store DNI  " + dni + " with Name " + persona.Name)
	
	data, err := json.Marshal(persona)
	if err != nil {
		fmt.Printf("[Management Chaincode][StoreCode]Error Marshaling Object")
		return shim.Error(err.Error())
	}
	// Initialize the chaincode
	// Write the state to the ledger
	err = stub.PutState(dni, []byte(data))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "modifyAddress" {
		// Make payment of X units from A to B
		return t.modifyAddress(stub, args)
	} else if function == "query" {
		// Deletes an entity from its state
		return t.query(stub, args)
	} 

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) modifyAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var dni, address string    // Entities
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	dni = args[0]
	address = args[1]

	var personObject Person;
	personStored, err := stub.GetState(dni)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	err = json.Unmarshal(personStored, &personObject)
	
	fmt.Println("Modifying person DNI  " + dni + " with Name " + personObject.Name)
	personObject.Address = address

	data, err := json.Marshal(personObject)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dni, data)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}


// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var dni string // Entities
	var err error
	fmt.Println("Metodo consultar")
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	dni = args[0]

	// Get the state from the ledger
	storedValue, err := stub.GetState(dni)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + dni + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(storedValue)
}