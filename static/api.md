- [获取钉钉扫码登录页面](#%E8%8E%B7%E5%8F%96%E9%92%89%E9%92%89%E6%89%AB%E7%A0%81%E7%99%BB%E5%BD%95%E9%A1%B5%E9%9D%A2)
  - [简要描述](#%E7%AE%80%E8%A6%81%E6%8F%8F%E8%BF%B0)
  - [请求URL](#%E8%AF%B7%E6%B1%82url)
  - [请求方式](#%E8%AF%B7%E6%B1%82%E6%96%B9%E5%BC%8F)
  - [参数](#%E5%8F%82%E6%95%B0)
  - [返回说明](#%E8%BF%94%E5%9B%9E%E8%AF%B4%E6%98%8E)
- [sso 回调接口](#sso-%E5%9B%9E%E8%B0%83%E6%8E%A5%E5%8F%A3)
  - [简要描述](#%E7%AE%80%E8%A6%81%E6%8F%8F%E8%BF%B0-1)
  - [请求URL](#%E8%AF%B7%E6%B1%82url-1)
  - [请求方式](#%E8%AF%B7%E6%B1%82%E6%96%B9%E5%BC%8F-1)
  - [参数](#%E5%8F%82%E6%95%B0-1)
  - [返回示说明](#%E8%BF%94%E5%9B%9E%E7%A4%BA%E8%AF%B4%E6%98%8E)
- [获取access_token](#%E8%8E%B7%E5%8F%96access_token)
  - [简要描述](#%E7%AE%80%E8%A6%81%E6%8F%8F%E8%BF%B0-2)
  - [请求URL](#%E8%AF%B7%E6%B1%82url-2)
  - [请求方式](#%E8%AF%B7%E6%B1%82%E6%96%B9%E5%BC%8F-2)
  - [头部参数](#%E5%A4%B4%E9%83%A8%E5%8F%82%E6%95%B0)
  - [Body参数](#body%E5%8F%82%E6%95%B0)
  - [返回示例](#%E8%BF%94%E5%9B%9E%E7%A4%BA%E4%BE%8B)
  - [返回参数说明](#%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E)
- [获取用户信息接口](#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF%E6%8E%A5%E5%8F%A3)
  - [简要描述](#%E7%AE%80%E8%A6%81%E6%8F%8F%E8%BF%B0-3)
  - [请求URL](#%E8%AF%B7%E6%B1%82url-3)
  - [请求方式](#%E8%AF%B7%E6%B1%82%E6%96%B9%E5%BC%8F-3)
  - [参数](#%E5%8F%82%E6%95%B0-2)
  - [返回示例](#%E8%BF%94%E5%9B%9E%E7%A4%BA%E4%BE%8B-1)
  - [返回参数说明](#%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E-1)

<br/> <br/> <br/> <br/> 

# 获取钉钉扫码登录页面   
## 简要描述

- 获取钉钉扫码登录页面

## 请求URL
- ` https://sso.onething.net/dingtalk/authorize `
  
## 请求方式
- GET 

## 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|redirect_url |是  |string |第三方平台callback接口的url（需要进行url编码）   |
|client_id |是  |string | h5微应用的AppKey（默认填写：dings251pjxcs810vukr）    |

## 返回说明
返回指向钉钉扫码登录页面的重定向Response

<br/> <br/> <br/> <br/> 


# sso 回调接口

## 简要描述

- sso 回调接口

## 请求URL
- ` https://sso.onething.net/dingtalk/callback `
  
## 请求方式
- GET 

## 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|forward_url |是  |string |第三方平台回调接口的url   |
|authCode |是  |string | 钉钉用户授权码   |
|state     |否  |string | 安全校验随机数    |

## 返回示说明
返回调用第三方平台回调接口的重定向Response

<br/> <br/> <br/> <br/> 

# 获取access_token
## 简要描述

- 获取access_token

## 请求URL
- ` https://sso.onething.net/dingtalk/token`
  
## 请求方式
- POST

## 头部参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|Authorization |是  |string |Basic Auth（默认为：Basic ZGluZ3MyNTFwanhjczgxMHZ1a3I6TXNJWUpUSjBhRE1ITFl3QTdiS2ZwNXRHN0Rqb3ZYZUNnVFRNR1F3TTVVLWUydU1wY2N5YTVIRmhfSWVCdVpYTA==）   |
|Content-Type |是  |string |Body数据类型类型（默认为：text/plain）   |

## Body参数
code={{authCode}}&state={{state}}


## 返回示例 

``` 
  {
    "access_token": 5b148efae4f9371ab52386f64109495e,
    "expiry_in": 7200,
	"refresh_token": "ab76ccb32a293e0e823554583e2680aa",
	"token_type": "Bearer"
  }
```

## 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|access_token |string   |生成的accessToken。  |
|refresh_token |string   |生成的refresh_token。可以使用此刷新token，定期的获取用户的accessToken  |
|expiry_in |long   |超时时间，单位秒。  |
|token_type |string   |默认：Bearer  |

## 特别说明
body中数据为一串字符串，客户端回调第三方平台回调接口时会携带code和state的值

<br/> <br/> <br/> <br/> 


# 获取用户信息接口
## 简要描述

- 获取用户信息接口

## 请求URL
- ` https://sso.onething.net/dingtalk/userinfo `

## 请求方式
- GET 

## 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|Authorization |是  |string |获取的access_token   |


## 返回示例 

``` 
  {
    "username":"xxx"
	"email":"xxx@onething.net"
  }
```

## 返回参数说明 
|参数名|类型|说明|
|:-----  |:-----|-----                           |
|username |string   |用户名  |
|email |string   |邮箱  |
