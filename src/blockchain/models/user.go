package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"blockchain/blocks"
	"fmt"
	"bytes"
)

type User struct {
	Id_     bson.ObjectId
	Name    string
	Account string
	Pwd     string
}

func Newuser(db *mgo.Database, name, account, pwd string)(user User, err error){
	user.Id_ = bson.NewObjectId()
	user.Name = name
	user.Account = account
	user.Pwd = pwd
	userCollection := db.C("users")
	err = userCollection.Insert(&user)
	if err == nil {
		bc := blocks.GetBlockchain()
		defer bc.DBClose()
		json :=  make(map[string]interface{})
		json["id"] = user.Id_.Hex()
		json["name"] = user.Name
		json["account"] = user.Account
		json["pwd"] = user.Pwd
		b := new(bytes.Buffer)
		for key, value := range json {
			fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
		}
		bc.AddBlock(b.String(), user.Id_.Hex())
	}
	return
}


func GetUser(db *mgo.Database, name string)(user User, err error){
	collection := db.C("users")
	err = collection.Find(bson.M{"name": name}).One(&user)
	return
	return
}


func GetUsers(db *mgo.Database)(users []User, err error){

	return
}
