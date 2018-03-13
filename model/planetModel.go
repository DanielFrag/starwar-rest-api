package model

import (
	"gopkg.in/mgo.v2/bson"
)

//Planet domain object
type Planet struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	Climate       string        `bson:"climate" json:"climate"`
	Terrain       string        `bson:"terrain" json:"terrain"`
	ForeignID     int32         `bson:"foreignId" json:"-"`
	NumberOfFilms int32         `json:"numberOfFilms"`
}
