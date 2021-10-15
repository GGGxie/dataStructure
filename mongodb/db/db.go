package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/GGGxie/dataStructure/mongodb/model"
	"github.com/GGGxie/dataStructure/mongodb/utils.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// D：一个BSON文档。这种类型应该在顺序重要的情况下使用，比如MongoDB命令。
// M：一张无序的map。它和D是一样的，只是它不保持顺序。
// A：一个BSON数组。
// E：D里面的一个元素。

func InsertOne() {
	//1.建立连接
	var (
		client     = utils.GetMgoCli()
		err        error
		collection *mongo.Collection
		lr         *model.LogRecord
		iResult    *mongo.InsertOneResult
		id         primitive.ObjectID
	)
	//2.选择数据库和表 my_collection
	collection = client.Database("runoob").Collection("test")
	//把结构体转化成json
	a := make(map[string]interface{})
	z, _ := json.Marshal(lr)
	json.Unmarshal(z, a)
	// 4.插入数据
	if iResult, err = collection.InsertOne(context.TODO(), a); err != nil {
		fmt.Print(err)
		return
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())
}

func InsertMany() {
	//1.建立连接
	var (
		client     = utils.GetMgoCli()
		err        error
		collection *mongo.Collection
		iResult    *mongo.InsertManyResult
		id         primitive.ObjectID
	)
	//2.选择数据库和集合
	collection = client.Database("runoob").Collection("test")
	//3.插入数据
	//批量插入
	iResult, err = collection.InsertMany(context.TODO(), []interface{}{
		model.LogRecord{
			JobName: "job10",
			Command: "echo 1",
			Err:     "",
			Content: "1",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job10",
			Command: "echo 2",
			Err:     "",
			Content: "2",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job10",
			Command: "echo 3",
			Err:     "",
			Content: "3",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job10",
			Command: "echo 4",
			Err:     "",
			Content: "4",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if iResult == nil {
		log.Fatal("result nil")
	}
	for _, v := range iResult.InsertedIDs {
		id = v.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
	}
}

//查询
func Find() {
	var (
		client     = utils.GetMgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)
	//2.选择数据库和集合
	collection = client.Database("runoob").Collection("test")

	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	// cond := map[string]string{"jobName": "job10"}
	// cond := model.FindByJobName{JobName: "job10"}

	cond := bson.M{"jobName": "job10"}
	//按照jobName字段进行过滤jobName="job10",翻页参数0-2
	if cursor, err = collection.Find(context.TODO(), cond, options.Find().SetSkip(0), options.Find().SetLimit(2)); err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr model.LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Println(lr)
	}

	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.LogRecord
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func Update() {
	var (
		client     = utils.GetMgoCli()
		collection *mongo.Collection
		err        error
		uResult    *mongo.UpdateResult
		//upsertedID model.LogRecord
	)
	//2.选择数据库和集合
	collection = client.Database("runoob").Collection("test")

	filter := bson.M{"jobName": "job10"}
	//update := bson.M{"$set": bson.M{"command": "ByBsonM",}}
	update := bson.M{"$set": model.UpdateByJobName{Command: "byModel", Content: "model"}}
	//update := bson.M{"$set": model.LogRecord{JobName:"job10",Command:"byModel"}}
	if uResult, err = collection.UpdateMany(context.TODO(), filter, update); err != nil {
		log.Fatal(err)
	}
	//uResult.MatchedCount表示符合过滤条件的记录数，即更新了多少条数据。
	log.Println(uResult.MatchedCount)
}

func Delete() {
	var (
		client     = utils.GetMgoCli()
		collection *mongo.Collection
		err        error
		uResult    *mongo.DeleteResult
	//upsertedID model.LogRecord
	)
	//2.选择数据库和集合
	collection = client.Database("runoob").Collection("test")
	filter := bson.M{"jobName": "job10"}
	//3.删除开始时间早于当前时间的数据
	//
	if uResult, err = collection.DeleteMany(context.TODO(), filter); err != nil {
		log.Fatal(err)
	}
	log.Println(uResult.DeletedCount)
}
