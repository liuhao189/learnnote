# curl

curl是利用URL语法在命令行方式下工作的开源文件传输工具，广泛应用在Uinx，Linux，并且有DOS，Win32和Win64下的移植版本。

# 常用命令

## 下载单个文件

默认将输出打印到标准输出中。

```bash
curl https://www.bing.com
```

通过-o/-O选项保存下载的文件到指定的文件中

-o，文件保存为命令行中指定的文件名的文件中

-O，URL中默认的文件名保存文件到本地。

```bash
curl -o my.html https://www.baidu.com

curl -O https://www.baidu.com
```

可以使用转向字符>对输出进行转向输出。

```bash
curl https://www.baidu.com > baidu.html
```

同时获取多个文件

curl -O URL1 -O URL2

同时从同一个站点下载多个文件时，curl会尝试重用链接。

```bash
curl -o main.html https://www.baidu.com -o news.html https://news.baidu.com/
```

通过-L 选项进行重定向。

```bash
curl www.meituan.com
curl -L www.meituan.com
```

断点续传，通过使用-C 选项可对大文件使用断点续传功能。

```bash
curl -C - -o rom.zip  http://openrom-dl.flyme.cn/12156221/20180404/1522826900034/flyme_Redmi-Note4x_136fei_6.8.4.3R-beta.zip
```

# 传递请求数据

默认curl使用GET方式请求数据，这种方式下直接通过URL传递数据，可以通过--data/-d方式制定使用POST方式传递数据。

```bash
curl -X POST --data "message=testcurl&nick=userone" localhost:8080/form_post
```

# 保存与使用网站cookie信息

```bash
curl -D siteCookie -L https://www.baidu.com
```

# 使用—-X指定其他协议

```bash
curl -X DELETE https://api.github.com/one.txt
```

# 更改HTTP Header

对于User-Agent，Cookie，Host这类标准的HTTP头部字段，通常会有另外一种设置方法。Curl命令提供了特定的选项来对这些头部字段进行设置。

-A，--user-agent，-b，--cookie，-e，--referer

```bash
curl -H "User-Agent:test browser" https://lvyou.xxx.com/trip-package-auditapi/api/biz/spu/image/upload
curl -A "test browser" https://lvyou.xxx.com/trip-package-auditapi/api/biz/spu/image/upload
```

# 上传文件

使用了-F参数，curl会以multipart/form-data的方式发送POST请求，-F参数以name=value的方式来指定参数内容，如果值是一个文件，则需要以name=@file的方式来指定。

```bash
curl -H "cookie:It is me" -F "file=@paojie.jpeg" https://www.xxx.com/upload
curl -H "cookie:It is me" -d "{\"name\":\"liuhao\"}" -H "Content-Type:application/json;charset=UTF-8" https://www.xx.com/changemodel
```