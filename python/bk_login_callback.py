SSO_DOMAIN = "https://sso.onething.net"
SSO_TOKEN_PATH = "dingtalk/token"
SSO_USERINFO_PATH = "dingtalk/userinfo"
APP_KEY = "dings251pjxcs810vukr"
SECRET_KEY = "MsIYJTJ0aDMHLYwA7bKfp5tG7DjovXeCgTTMGQwM5U-e2uMpccya5HFh_IeBuZXL"

def DingTalk_callback(request):
    # 获取 authcode
    code = request.GET.get("code", "")
    state = request.GET.get("state", "")

    # 调用 https://sso.onething.net/dingtalk/token 获取token
    token_url = str.format("{}/{}", SSO_DOMAIN, SSO_TOKEN_PATH)
    data = str.format("code={}&state={}", code, state)
    headers = {
                                  "Authorization": get_basic_auth_str(APP_KEY, SECRET_KEY),
                                  "Content-Type": "text/plain",
                              }
    _data = requests.post(url=token_url, headers=headers, data=data, timeout=None, verify=None, cert=None, cookies=None).json()
    access_token = _data['access_token']

    # 调用 https://sso.onething.net/dingtalk/userinfo 获取用户信息
    userinfo_url = str.format("{}/{}", SSO_DOMAIN, SSO_USERINFO_PATH)
    ok, _data = http_get(userinfo_url,
                         data=None,
                         headers={
                             "Authorization": access_token,
                         },
                         )
    if not ok:
        logger.error("DingTalk_callback:get userinfo error")
        return login_failed_response(request, redirect_to, app_id=None)

    # 校验用户信息
    # ...
    return repsonse