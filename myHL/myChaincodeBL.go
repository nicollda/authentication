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



//todo:  add constants for all string litterals
//todo:  need to make consitent status.  need better way to take them out of the process when closed
//todo: data abstraction layer, abstract persistance
//todo: add security to get user names



package main

import (
	"errors"
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const separator = 		"."
const userIndex =		"UserIndex" + separator
const roleIndex =		"RoleIndex" + separator

type ChaincodeBusinessLayer struct {
	userRep			UserRepository 
	roleRep			RoleRepository
	stub shim.ChaincodeStubInterface
}

func (t *ChaincodeBusinessLayer) initObjects(stub shim.ChaincodeStubInterface) error {
	t.stub = stub
	t.writeOut("in init objects")
	
	
	//initialize our repositories
	t.roleRep.init(stub)
	t.userRep.init(stub)
	
	return nil
}





//********************************************************************************************************
//****                        Debug function inplimentations                                          ****
//********************************************************************************************************


const debug = true


func (t *ChaincodeBusinessLayer) writeOut(out string) ([]byte, error) {
	if debug {
		curOutByteA,err := t.stub.GetState("currentOutput")
		outByteA := []byte(string(curOutByteA) + ":::" + out)
		err = t.stub.PutState("currentOutput", outByteA)
		return nil, err
	}
	
	return nil, nil
}



func (t *ChaincodeBusinessLayer) readOut() string {
	if debug {
		curOutByteA, err := t.stub.GetState("currentOutput")
		if err != nil {
			return "error"
		}
		
		return string(curOutByteA)
	}
	
	return ""
}





//********************************************************************************************************
//****                        Invoke function implimentations                                         ****
//********************************************************************************************************


// register user
func (t *ChaincodeBusinessLayer) registerUser(userID string, password string, roles string) ([]byte, error) {
	fmt.Printf("Running registerUser")
	//need to make sure the user is not already registered
	
	index, err := t.userRep.newUser(userID, password, roles, "Active")
	if err != nil {
		return nil, err
	}
	
	return []byte(index), nil
}


// register role
func (t *ChaincodeBusinessLayer) registerRole(roleID string, name string) ([]byte, error) {
	fmt.Printf("Running registerRole")
	//need to make sure the user is not already registered
	
	index, err := t.roleRep.newRole(roleID, name, "Active")
	if err != nil {
		return nil, err
	}
	
	return []byte(index), nil
}


func (t *ChaincodeBusinessLayer) encrypt(pwd string) ([]byte, error) {

	return []byte(pwd), nil
}

func (t *ChaincodeBusinessLayer) authenticate(userID string, password string) ([]byte, error) {
	fmt.Printf("Running invoke")
	
	var Password []byte
	var err error
	var user User
	
	Password, _ = t.encrypt(password)
	user, err = t.userRep.getUser(userID)

	
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	
	if user.Password == "" {
		return nil, nil
	}
	
	
	if bytes.Equal([]byte(user.Password), Password) {
		return user.getRoles(t.roleRep)
	} else {
		return nil, nil
	}

	return nil, nil   
}


func (t *ChaincodeBusinessLayer) getRoles(user User) ([]byte, error) {
	var roleArray []string
	var roleOut string
	
	rolejson := user.Roles
	roleOut = ""
	
	err := json.Unmarshal([]byte(rolejson), &roleArray) 
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	
	for _, roleID := range roleArray {
		role, err := t.roleRep.getRole(roleID)
		if err != nil {
			return nil, errors.New("Failed to get state")
		}
	
		if roleOut == "" {
			roleOut = "\"" + role.Name + "\""
		} else {
			roleOut = roleOut + ", \"" + role.Name + "\""
		}
	}
	
	return []byte("[" + roleOut + "]"), nil	
}



//   curently not used but should be used in place of taking the user id via the interface.  user id should come from the security model
func (t *ChaincodeBusinessLayer) getUserID(args []string) ([]byte, error) {
	//returns the user's ID 
	
	return nil, nil  //dont know how to get the current user
}