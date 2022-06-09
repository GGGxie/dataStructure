package main

import (
	"encoding/json"
	"net/http"

	dhttp "github.com/GGGxie/dataStructure/pkg/http"
	"github.com/gin-gonic/gin"
)

type Resp2 struct {
	Success   bool        `json:"success"`
	ErrorCode interface{} `json:"errorCode"`
	ErrorMsg  interface{} `json:"errorMsg"`
	Data      Data        `json:"data"`
}
type Data struct {
	CorpID string `json:"corpId"`
}
type CommonResp struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Result  bool        `json:"result"`
}
type respStruct struct {
	Errcode int    `json:"errcode"`
	Result  Result `json:"result"`
	Errmsg  string `json:"errmsg"`
}
type Result struct {
	AssociatedUnionid string `json:"associated_unionid"`
	Unionid           string `json:"unionid"`
	DeviceID          string `json:"device_id"`
	SysLevel          int    `json:"sys_level"`
	Name              string `json:"name"`
	Sys               bool   `json:"sys"`
	Userid            string `json:"userid"`
}

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/config", func(c *gin.Context) {
		commonResp := &Resp2{
			Success: true,
			Data: Data{
				CorpID: "dinge70f54a68fcfc101f2c783f7214b6d69",
			},
			ErrorMsg:  nil,
			ErrorCode: nil,
		}
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "x-access-token, Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Cache-Control", "no-cache")
		c.JSON(http.StatusOK, commonResp)
	})

	r.GET("/login", func(c *gin.Context) {
		code := c.Query("authCode")
		url := "https://oapi.dingtalk.com/topapi/v2/user/getuserinfo" //必须要用https,否则报 missing code错误
		data := struct {
			Code string `json:"code"`
		}{
			Code: code,
		}
		query := map[string]string{
			"access_token": "a4dd15620138395e98c2dbb82f3d7118",
		}

		resp, _ := dhttp.Post(url, data, nil, query)
		var Resp respStruct
		json.Unmarshal(resp.RespBody, &Resp)

		commonResp := &CommonResp{
			Result:  true,
			Message: "",
			Data:    Resp,
			Code:    resp.Code,
		}
		c.JSON(http.StatusOK, commonResp)
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

// func main() {
// 	a := "str"
// 	p1 := (*StringHeader)(unsafe.Pointer(&a))
// 	fmt.Printf("%x,%d\n", &p1.Data, p1.Len)
// 	b := "st2123123123123123123213123123123112312312312312312312"
// 	a = a + b
// 	p2 := (*StringHeader)(unsafe.Pointer(&a))
// 	fmt.Printf("%x,%d\n", &p2.Data, p1.Len)
// }
