package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	sql := `SELECT
    rule_name,
    sum(c) c
from
    (
        SELECT
            a.severity severity,
            rule_name,
            COUNT(distinct d.event_id) c
        FROM
            metis.events AS a
            LEFT JOIN metis.alert_event_feedback AS b ON (
                a.id = b.event_id
                AND b.created_at > addWeeks(today(), -1)
            )
            INNER JOIN (
                SELECT
                    event_id,
                    ifnull(
                        arrayJoin(
                            splitByString(
                                ';',
                                reverse(
                                    splitByString('@', reverse(ifNull(receiver, ''))) [ 1 ]
                                )
                            )
                        ),
                        '群告警'
                    ) AS receiver,
                    MIN(status) as status
                FROM
                    metis.notifications
                where
                    created_at >= addWeeks(today(), -1)
                GROUP BY
                    event_id,
                    receiver
            ) AS d ON d.event_id = a.id
        WHERE
            a.status IN ('alerting')
            and a.created_at >= toDateTime(1677231482)
            AND a.created_at < toDateTime(1677836282)
            and receiver in ('邓盾')
        GROUP BY
            severity,
            business,
            rule_name,
            receiver
    )
GROUP BY
    rule_name,
    severity
ORDER BY c DESC 
    FORMAT JSON`
	//发起 http get 请求
	resp, err := Get("https://monitor.onething.net/api/datasources/proxy/21",
		nil,
		map[string]string{
			"Authorization": "Basic Z3JhZmFuYV9ib3RAb25ldGhpbmcubmV0OmdyYWZhbmFfYm90IUAjb3QxIw==",
		},
		map[string]string{
			"query": sql,
		})
	fmt.Println(err, resp.Code, string(resp.RespBody))

	var gResp GrafanaQueryResp
	json.Unmarshal(resp.RespBody, &gResp)
	fmt.Printf("%+v", gResp)
}

type GrafanaQueryResp struct {
	Meta       []Meta                   `json:"meta"`
	Data       []map[string]interface{} `json:"data"`
	Rows       int                      `json:"rows"`
	Statistics Statistics               `json:"statistics"`
}
type Meta struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Statistics struct {
	Elapsed   float64 `json:"elapsed"`
	RowsRead  int     `json:"rows_read"`
	BytesRead int     `json:"bytes_read"`
}

func Get(url string, data interface{}, header map[string]string, query map[string]string) (*Resp, error) {

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	//设置header参数
	for k, v := range header {
		req.Header.Set(k, v)
	}

	//设置query参数
	tempQuery := req.URL.Query()
	for k, v := range query {
		tempQuery.Set(k, v)
	}

	req.URL.RawQuery = tempQuery.Encode()
	//请求 https 网站跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//获取response
	result, err := ParseResp(resp)
	return &Resp{
		Code:     resp.StatusCode,
		RespBody: result,
	}, err
}

type Resp struct {
	RespBody []byte
	Code     int
}

func ParseResp(resp *http.Response) ([]byte, error) {
	if result, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
