# -*- coding: utf-8 -*-
#!/usr/bin/python
import sys
import base64
import requests


SSO_DOMAIN = "http://localhost:8080"
SSO_TOKEN_PATH = "dingtalk/token"
SSO_USERINFO_PATH = "dingtalk/userinfo"
APP_KEY = "dings251pjxcs810vukr"
SECRET_KEY = "MsIYJTJ0aDMHLYwA7bKfp5tG7DjovXeCgTTMGQwM5U-e2uMpccya5HFh_IeBuZXL"
def _http_request(method, url, headers=None, data=None, timeout=None, verify=False, cert=None, cookies=None):
    try:
        if method == "GET":
            resp = requests.get(
                url=url, headers=headers, params=data, timeout=timeout, verify=verify, cert=cert, cookies=cookies
            )
        elif method == "HEAD":
            resp = requests.head(url=url, headers=headers, verify=verify, cert=cert, cookies=cookies)
        elif method == "POST":
            resp = requests.post(
                url=url, headers=headers, json=data, timeout=timeout, verify=verify, cert=cert, cookies=cookies,
            )
        elif method == "DELETE":
            resp = requests.delete(
                url=url, headers=headers, json=data, timeout=timeout, verify=verify, cert=cert, cookies=cookies
            )
        elif method == "PUT":
            resp = requests.put(
                url=url, headers=headers, json=data, timeout=timeout, verify=verify, cert=cert, cookies=cookies
            )
        else:
            return False, None
    except requests.exceptions.RequestException:
        print("http request error! method: %s, url: %s, data: %s", method, url, data)
        # logger.exception("http request error! method: %s, url: %s, data: %s", method, url, data)
        return False, None
    else:
        if resp.status_code != 200:
            content = resp.content[:100] if resp.content else ""
            error_msg = "http request fail! method: %s, url: %s, " "response_status_code: %s, response_content: %s"
            # if isinstance(content, str):
            #     try:
            #         content = content.decode('utf-8')
            #     except Exception:
            #         content = content
            # logger.error(error_msg, method, url, resp.status_code, content)
            return False, None

        return True, resp.json()

def http_post(url, data, headers=None, verify=False, cert=None, timeout=None, cookies=None):
    if not headers:
        headers = _gen_header()
    return _http_request(
        method="POST", url=url, headers=headers, data=data, timeout=timeout, verify=verify, cert=cert, cookies=cookies
    )
def get_basic_auth_str(username, password):
    temp_str = username + ':' + password
    # 转成bytes string
    bytesString = temp_str.encode(encoding="utf-8")
    # base64 编码
    encodestr = base64.b64encode(bytesString)
    return 'Basic ' + encodestr.decode()

state=""
code = "745cb6b63d9933809f46218c8f67aec8"
token_url = str.format("{}/{}", SSO_DOMAIN, SSO_TOKEN_PATH)
data = str.format("code={}&state={}", code, state)
print(data)
headers={
                              "Authorization": get_basic_auth_str(APP_KEY, SECRET_KEY),
                              "Content-Type": "text/plain",
                          }
# ok, _data = http_post(token_url,
#                     data=data,
#                     headers={
#                         "Authorization": get_basic_auth_str(APP_KEY, SECRET_KEY),
#                         "Content-Type": "text/plain",
#                     },
#                     )
_data = requests.post(
        url=token_url, headers=headers, data=data, timeout=None, verify=None, cert=None, cookies=None).json()
# print(ok)
print(_data)

# print(get_basic_auth_str("dings251pjxcs810vukr","MsIYJTJ0aDMHLYwA7bKfp5tG7DjovXeCgTTMGQwM5U-e2uMpccya5HFh_IeBuZXL"))