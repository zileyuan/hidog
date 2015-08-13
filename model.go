package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Dog struct {
	Id       bson.ObjectId `bson:"_id"`
	Title    string        `bson:"Title"`
	Subtitle string        `bson:"Subtitle"`
	Sex      string        `bson:"Sex"`
	Preview  string        `bson:"Preview"`
	Picture1 string        `bson:"Picture1"`
	Picture2 string        `bson:"Picture2"`
	Picture3 string        `bson:"Picture3"`
	Picture4 string        `bson:"Picture4"`
	Picture5 string        `bson:"Picture5"`
	Memo     string        `bson:"Memo"`
	Like     int           `bson:"Like"`
}

type Pup struct {
	Id       bson.ObjectId `bson:"_id"`
	Title    string        `bson:"Title"`
	Subtitle string        `bson:"Subtitle"`
	Sex      string        `bson:"Sex"`
	Preview  string        `bson:"Preview"`
	Picture1 string        `bson:"Picture1"`
	Picture2 string        `bson:"Picture2"`
	Picture3 string        `bson:"Picture3"`
	Picture4 string        `bson:"Picture4"`
	Picture5 string        `bson:"Picture5"`
	Memo     string        `bson:"Memo"`
	Like     int           `bson:"Like"`
}

type Comment struct {
	Id       bson.ObjectId `bson:"_id"`
	Title    string        `bson:"Title"`
	Content  string        `bson:"Content"`
	DateTime time.Time     `bson:"DateTime"`
}

type Global struct {
	Id    bson.ObjectId `bson:"_id"`
	Key   string        `bson:"Key"`
	Value string        `bson:"Value"`
}

type Account struct {
	Id       bson.ObjectId `bson:"_id"`
	UserName string        `bson:"UserName"`
	Password string        `bson:"Password"`
	Role     int           `bson:"Role"`
}
