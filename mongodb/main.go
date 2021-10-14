package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/GGGxie/dataStructure/mongodb/model"
	"github.com/GGGxie/dataStructure/mongodb/utils.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	//1.建立连接
	var (
		client     = utils.GetMgoCli()
		err        error
		collection *mongo.Collection
		lr         *model.LogRecord
		iResult    *mongo.InsertOneResult
		id         primitive.ObjectID
	)
	//2.选择数据库 my_db
	db := client.Database("runoob")

	//3.选择表 my_collection
	collection = db.Collection("test")
	a := make(map[string]interface{})
	z, _ := json.Marshal(lr)
	json.Unmarshal(z, a)
	//4.插入数据
	if iResult, err = collection.InsertOne(context.TODO(), a); err != nil {
		fmt.Print(err)
		return
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())

}
