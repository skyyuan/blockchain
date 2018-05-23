package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"blockchain/blocks"
	"fmt"
	"encoding/json"
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
		out, err := json.Marshal(user)
		if err != nil {
			panic (err)
		}

		fmt.Println(string(out))

		bc.AddBlock(string(out))
	}
	return
}


func GetUser(db *mgo.Database, name string)(user User, err error){
	collection := db.C("users")
	err = collection.Find(bson.M{"name": name}).One(&user)
	return
	return
}


func GetUserByAccount(db *mgo.Database, account string)(user User, err error){
	collection := db.C("users")
	err = collection.Find(bson.M{"account": account}).One(&user)
	return
	return
}

func GetUsers(db *mgo.Database)(users []User, err error){

	return
}
