package main

import (
//	"errors"
//	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
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
	
	//debug code
	user, err = self.getUser(userID)
	if err != nil {
		return "", err
	}
	
	curOutByteA,err := self.LinkedList.stub.GetState("currentOutput")
	outByteA := []byte(string(curOutByteA) + ":::debug for userID " + user.UserID)
	err = self.LinkedList.stub.PutState("currentOutput", outByteA)
	
	
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
	user.Password = "Password1"
	user.Roles = "[\"Button1\",\"Button2\"]"
	return user, err
}



func (self *UserRepository) updateUser(user User) (string, error) {

	return self.LinkedList.put(user.UserID, user)
}



func (self *UserRepository) deleteUser(userID string) error {
	
	return self.LinkedList.del(userID)
}



//********************************
//         User
//********************************

type User struct {
	UserID		string	`json:"userID"`
	Status		string	`json:"status"`
	Password	string	`json:"ballance"`
	Roles		string	`json:"roles"`
}



func (self *User) init(userID string, password string, roles string, status string) error {
	self.UserID = userID
	self.Status = status
	self.Password = password
	self.Roles = roles
	
	return nil
}