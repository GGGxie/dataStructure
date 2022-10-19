window.onload = function () {
    var domain = "https://devops-stage.onething.net:443";
    dd.ready(function () {
        // var corpId;
        //corpId = 'dingff6a1f72f162ef7235c2f4657eb6378f';
        corpId = 'dinge70f54a68fcfc101f2c783f7214b6d69'
        dd.runtime.permission.requestAuthCode({
            corpId: corpId, //三方企业ID
            onSuccess: function (result) {
                alert("5555555");
                var httpRequest = new XMLHttpRequest();//第一步：建立所需的对象
                var url = domain + "/login/?c_url=https%3A//devops-stage.onething.net/o/bk_itsm/&authCode=" + result.code;
                httpRequest.open('GET', url, true);//第二步：打开连接  将请求参数写在url中  ps:"./Ptest.php?name=test&nameone=testone"
                httpRequest.setRequestHeader('Accept-Language', 'zh-CN,zh;q=0.9');
                httpRequest.send();//第三步：发送请求  将请求参数写在URL中
                /**
                 * 获取数据后的处理程序
                 */
                httpRequest.onreadystatechange = function () {
                    alert(this.readyState);
                    alert(this.status);
                    if (this.readyState == 4 && this.status == 200) {
                        dd.biz.util.openLink({
                            url: this.responseURL,//要打开链接的地址
                            onSuccess: function (result) {
                                /**/
                            },
                            onFail: function (err) {
                            }
                        });
                    }
                };
            },
            onFail: function (err) {
                alert("77777777");
                alert(JSON.stringify(err))
            }
        });
    }, []);
}