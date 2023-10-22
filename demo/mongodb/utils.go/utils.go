package utils

import (
	"context"
	"log"
	"sync"

	"github.com/qiniu/qmgo"
)

var mgoCli *qmgo.Client
var once sync.Once

// func initEngine() {
// 	var err error
// 	clientOptions := options.Client().ApplyURI("mongodb://39.108.148.65:27017")

//		// 连接到MongoDB
//		mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
//		if err != nil {
//			log.Fatal(err)
//		}
//		// 检查连接
//		err = mgoCli.Ping(context.TODO(), nil)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
func initEngine() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://39.108.148.65:27017"})
	if err != nil {
		log.Fatal(err)
		return
	}

	mgoCli = client
}
func GetMgoCli() *qmgo.Client {
	once.Do(initEngine)
	// if mgoCli == nil {
	// 	initEngine()
	// }
	return mgoCli
}
