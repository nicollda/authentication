/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main


import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
	bl ChaincodeBusinessLayer
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	
	
	//initialize our repositories
	t.bl.initObjects(stub)
	
	t.bl.writeOut("in init")
	
	_, err := t.bl.registerUser("David", "Password1", "[\"Button1\",\"Button2\"]")
	if err != nil {
		return nil, err
	}

	return nil, nil
}



// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("Running delete")
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Invoke callback representing the invocation of a chaincode
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "authenticate" {
		fmt.Printf("Function is authenticate")
		if len(args) != 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting 2")
		}
		
		return t.bl.authenticate(args[0], args[1])
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Handle different functions
	if function == "authenticate" {
		fmt.Printf("Function is authenticate")
		
		if len(args) != 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting 2")
		}
		//return []byte("[\"Button1\",\"Button2\"]"), nil
		return t.bl.authenticate(args[0], args[1])
		
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	}

	return nil, errors.New("Received unknown function invocation")
}	
	

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}