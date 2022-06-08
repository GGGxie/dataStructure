// package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	dhttp "github.com/GGGxie/dataStructure/pkg/http"
// 	"github.com/gin-gonic/gin"
// )

// type CommonResp struct {
// 	Message string      `json:"message"`
// 	Code    int         `json:"code"`
// 	Data    interface{} `json:"data"`
// 	Result  bool        `json:"result"`
// }
// type respStruct struct {
// 	Errcode int    `json:"errcode"`
// 	Result  Result `json:"result"`
// 	Errmsg  string `json:"errmsg"`
// }
// type Result struct {
// 	AssociatedUnionid string `json:"associated_unionid"`
// 	Unionid           string `json:"unionid"`
// 	DeviceID          string `json:"device_id"`
// 	SysLevel          int    `json:"sys_level"`
// 	Name              string `json:"name"`
// 	Sys               bool   `json:"sys"`
// 	Userid            string `json:"userid"`
// }

// var db = make(map[string]string)

// func setupRouter() *gin.Engine {

// 	r := gin.Default()
// 	// r.Static("/code", "js/code.js")
// 	r.LoadHTMLFiles("js/code.html")
// 	// a, _ := ioutil.ReadFile("js/code.html")
// 	r.GET("/code", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "code.html", gin.H{
// 			"title": "main site",
// 		})
// 	})
// 	// Ping test
// 	r.GET("/getcode", func(c *gin.Context) {
// 		code := c.Query("code")
// 		url := "https://oapi.dingtalk.com/topapi/v2/user/getuserinfo" //必须要用https,否则报 missing code错误
// 		data := struct {
// 			Code string `json:"code"`
// 		}{
// 			Code: code,
// 		}
// 		query := map[string]string{
// 			"access_token": "0e5ff5012b173d31a67cf50e79e7469c",
// 		}

// 		resp, _ := dhttp.Post(url, data, nil, query)
// 		var Resp respStruct
// 		json.Unmarshal(resp.RespBody, &Resp)

// 		commonResp := &CommonResp{
// 			Result:  true,
// 			Message: "",
// 			Data:    Resp,
// 			Code:    resp.Code,
// 		}
// 		c.JSON(http.StatusOK, commonResp)
// 	})
// 	return r
// }

// func main() {
// 	r := setupRouter()
// 	// Listen and Server in 0.0.0.0:8080
// 	r.Run(":8080")
// }

// // func main() {
// // 	a := "str"
// // 	p1 := (*StringHeader)(unsafe.Pointer(&a))
// // 	fmt.Printf("%x,%d\n", &p1.Data, p1.Len)
// // 	b := "st2123123123123123123213123123123112312312312312312312"
// // 	a = a + b
// // 	p2 := (*StringHeader)(unsafe.Pointer(&a))
// // 	fmt.Printf("%x,%d\n", &p2.Data, p1.Len)
// // }

package main

import (
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dingtalkoauth2_1_0 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *dingtalkoauth2_1_0.Client, _err error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &dingtalkoauth2_1_0.Client{}
	_result, _err = dingtalkoauth2_1_0.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}
	getSsoUserInfoHeaders := &dingtalkoauth2_1_0.GetSsoUserInfoHeaders{}
	getSsoUserInfoHeaders.XAcsDingtalkAccessToken = tea.String("<your access token>")
	getSsoUserInfoRequest := &dingtalkoauth2_1_0.GetSsoUserInfoRequest{
		Code: tea.String("tokenxxxx"),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err = client.GetSsoUserInfoWithOptions(getSsoUserInfoRequest, getSsoUserInfoHeaders, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(util.Empty(err.Code)) && !tea.BoolValue(util.Empty(err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
		}

	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
