package main

import (
	"errors"
//	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)



//********************************
//         User Repository
//********************************

type UserRepository struct {
	LinkedList ChainLinkedList
}



func (self *UserRepository) init(stub shim.ChaincodeStubInterface) error {
	
	return self.LinkedList.init(stub, userIndex)
}



func (self *UserRepository) newUser(userID string, password string, roles string, status string) (string, error) {
	
	var user User
	user.init(userID, password, roles, status)
	
	key, err := self.LinkedList.put(userID, user)
	if err != nil {
		return "", err
	}
	
	/*
	//debug code
	user, err = self.getUser(userID)
	if err != nil {
		return "", err
	}
	
	curOutByteA,err := self.LinkedList.stub.GetState("currentOutput")
	outByteA := []byte(string(curOutByteA) + ":::debug for userID " + user.UserID)
	err = self.LinkedList.stub.PutState("currentOutput", outByteA)
	*/
	
	return key, nil
}



func (self *UserRepository) getFirstUser() (User, error) {
	var user User
	
	err := self.LinkedList.getFirst(&user)
	return user, err
}



func (self *UserRepository) getNextUser() (User, error) {
	var user User
	
	err := self.LinkedList.getNext(&user)
	return user, err
}



func (self *UserRepository) getUser(userId string) (User, error) {
	var user User
	var err error
	
	err = self.LinkedList.get(userId, &user)

	return user, err
}



func (self *UserRepository) updateUser(user User) (string, error) {

	return self.LinkedList.put(user.UserID, user)
}



func (self *UserRepository) deleteUser(userID string) error {
	
	return self.LinkedList.del(userID)
}



//********************************
//         Role Repository
//********************************

type RoleRepository struct {
	LinkedList ChainLinkedList
}



func (self *RoleRepository) init(stub shim.ChaincodeStubInterface) error {
	
	return self.LinkedList.init(stub, roleIndex)
}



func (self *RoleRepository) newRole(roleID string, name string, status string) (string, error) {
	
	
	//todo:  swtich to role elements
	var role Role
	role.init(roleID, name, status)
	
	key, err := self.LinkedList.put(roleID, role)
	if err != nil {
		return "", err
	}
	
	/*
	//debug code
	user, err = self.getUser(userID)
	if err != nil {
		return "", err
	}
	
	curOutByteA,err := self.LinkedList.stub.GetState("currentOutput")
	outByteA := []byte(string(curOutByteA) + ":::debug for userID " + user.UserID)
	err = self.LinkedList.stub.PutState("currentOutput", outByteA)
	*/
	
	return key, nil
}



func (self *RoleRepository) getFirstRole() (Role, error) {
	var role Role
	
	err := self.LinkedList.getFirst(&role)
	return role, err
}



func (self *RoleRepository) getNextRole() (Role, error) {
	var role Role
	
	err := self.LinkedList.getNext(&role)
	return role, err
}



func (self *RoleRepository) getRole(roleId string) (Role, error) {
	var role Role
	var err error
	
	err = self.LinkedList.get(roleId, &role)

	return role, err
}



func (self *RoleRepository) updateRole(role Role) (string, error) {

	return self.LinkedList.put(role.RoleID, role)
}



func (self *RoleRepository) deleteRole(roleID string) error {
	
	return self.LinkedList.del(roleID)
}


//********************************
//         User
//********************************

type User struct {
	UserID		string	`json:"userID"`
	Status		string	`json:"status"`
	Password	string	`json:"password"`
	Roles		string	`json:"roles"`
}

func (self *User) init(userID string, password string, roles string, status string) error {
	self.UserID = userID
	self.Status = status
	self.Password = password
	self.Roles = roles

	return nil
}

func (self *User) getRoles(roleRep RoleRepository) ([]byte, error) {
	var roleArray []string
	var roleOut string

	rolejson := self.Roles
	roleOut = ""

	err := json.Unmarshal([]byte(rolejson), &roleArray)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}

	for _, roleID := range roleArray {
		role, err := roleRep.getRole(roleID)
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

//********************************
//         Role
//********************************


type Role struct {
	RoleID string `json:"roleID"`
	Status string `json:"status"`
	Name string `json:"name"`
}


func (self *Role) init(roleID string, name string, status string) error {
	self.RoleID = roleID
	self.Status = status
	self.Name = name

	return nil
} 