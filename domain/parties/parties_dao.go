package parties

import (
	"context"
	"fmt"
	"github.com/dogancancelikk/election/dao/store"
	"github.com/dogancancelikk/election/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (party *Party) Save() *errors.RestError {
	collection := store.Database.Collection("party")
	insertResult, err := collection.InsertOne(context.TODO(), party)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(insertResult.InsertedID)
	return nil
}

func (party *Party) GetParties() ([]Party, *errors.RestError) {
	parties := make([]Party, 0)
	collection := store.Database.Collection("party")
	collection.FindOne(context.TODO(), bson.M{"name": "TİP"}).Decode(&parties)
	fmt.Println(parties)
	return parties, nil
}
