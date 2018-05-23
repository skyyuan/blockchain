package utils

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"errors"
)

var session *mgo.Session

func init() {
	InitMgo()
}


func ConnMgo() *mgo.Session {
	return session.Copy()
}

func InitMgo() {
	url := beego.AppConfig.String("mongodb::url")

	sess, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	session = sess.Copy()

	session.SetMode(mgo.Monotonic, true)
}

func GetMgoDbSession() (*mgo.Database, *mgo.Session){
	// get db
	dbn := beego.AppConfig.String("mongodb::DataBaseName")
	if len(dbn) == 0 {
		panic(errors.New("No DataBaseName"))
	}
	mgoSession := ConnMgo()
	db := mgoSession.DB(dbn)
	return db, mgoSession
}