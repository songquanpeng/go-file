<p align="center">
  <a href="https://github.com/songquanpeng/go-file"><img src="https://user-images.githubusercontent.com/39998050/108494937-1a573e80-72e3-11eb-81c3-5545d7c2ed6e.jpg" width="200" height="200" alt="go-file"></a>
</p>

<div align="center">

# Go File

_✨ 文件分享工具，仅单个可执行文件，开箱即用，可用于局域网内分享文件和文件夹，直接跑满本地带宽 ✨_  

</div>

<p align="center">
  <a href="https://raw.githubusercontent.com/songquanpeng/go-file/master/LICENSE">
    <img src="https://img.shields.io/github/license/songquanpeng/go-file?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/songquanpeng/go-file/releases/latest">
    <img src="https://img.shields.io/github/v/release/songquanpeng/go-file?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://github.com/songquanpeng/go-file/releases/latest">
    <img src="https://img.shields.io/github/downloads/songquanpeng/go-file/total?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://goreportcard.com/report/github.com/songquanpeng/go-file">
  <img src="https://goreportcard.com/badge/github.com/songquanpeng/go-file" alt="GoReportCard">
  </a>
</p>

<p align="center">
  <a href="https://github.com/songquanpeng/go-file/projects/1">开发规划</a>
  ·
  <a href="https://github.com/songquanpeng/go-file/releases">程序下载</a>
  ·
  <a href="https://iamazing.cn/page/LAN-SHARE-使用教程">使用教程</a>
  ·
  <a href="#演示">截图展示</a>
</p>


<details>
<summary><strong><i>English</i></strong></summary>
<div>

Warning: The English version is outdated.

## Description
File sharing tool, can be used to share files in a LAN.

## Features
1. No need to configure environment and there is only a single executable file.
2. Automatically open browser to make you share file more quickly.
3. Generate QR codes for your mobile phone to scan.
4. Easily share all the content of a local dir.

## Usage
*For v0.3.4 and below.*

Just double-click to use with default port `3000` and default token (used to verify identity when user try to delete files) `token`.

If you want to change the port and token, run it like this:`./go-file.exe --port 80 --token private`.

Your can also public a local path by providing a `path` like this : `./go-file.exe --path ./this/is/a/path` 

```
Usage of go-file.exe:
  -host string
        the server's ip address or domain (default "localhost")
  -path string
        specify a local path to public
  -port int
        specify the server listening port. (default 3000)
  -token string
        specify the private token. (default "token")
  -video string
        specify a video folder to public
```

## Demo
Please visit https://go-file.herokuapp.com/ to have a try yourself.

![index page](https://user-images.githubusercontent.com/39998050/130427067-80bf3cc5-5fee-488a-bea5-e323b9458064.png)
![explorer page](https://user-images.githubusercontent.com/39998050/177032568-8af95d7e-87ab-4e60-804b-5e49addfb6ab.png)
![image page](https://user-images.githubusercontent.com/39998050/177032659-c8c68186-09f4-4142-9f57-70bcb4a4cda1.png)
![video page](https://user-images.githubusercontent.com/39998050/177032588-8946abde-a8da-45a2-a389-c16dba9cea34.png)


## Others
[Node.js version is here.](https://github.com/songquanpeng/lan-share)
</div>
</details>


## 特点
1. 无需配置环境，仅单个可执行文件，直接双击即可开始使用。
2. 自动打开浏览器，分享文件快人一步。
3. 提供二维码，可供移动端扫描下载文件，告别手动输入链接。
4. 支持分享本地文件夹。
5. 适配移动端。
6. 内置图床，支持直接粘贴上传图片，提供图片上传 API。
7. 内置视频播放页面，可用于在其他设备上在线博客自己电脑上的视频，轻松跨设备在线看视频。
8. 支持拖拽上传，拷贝上传。
9. 允许对不同类型的用户设置文件访问权限限制。
10. 访问频率限制。
11. 支持 Token API 验证，便于与其他系统整合。

## 使用方法
> v0.3.3 以及之前版本的使用方法请[点击此处](https://github.com/songquanpeng/go-file/tree/52e8303e33e99bbcaf583d2d5a5bb0ec197bc676#使用方法)。

直接双击即可使用，默认端口为 `3000`，程序在第一次启动时会自动创建管理员账户，用户名为 `admin`，密码为 `123456`，记得登录后到 `管理页面` 下的 `账户管理` 标签页中更改你的用户密码。

之后程序将自动为你打开浏览器，点击右上角的 `上传` 按钮即可上传，支持拖放上传，支持同时上传多个文件。

**进阶使用：**
1. 如果要修改端口，启动时请指定 `port` 参数：`./go-file.exe --port 80`。
2. 如果需要分享文件夹，启动时请指定 `path` 参数：`./go-file.exe --path ./this/is/a/path`，之后点击导航栏上的 `文件` 即可。
3. 如果需要分享本地的视频资源，加 `video` 参数：`./go-file.exe --video ./this/is/a/path`，之后点击导航栏上的 `视频` 即可。
4. 如果需要启用访问速率控制，需要在启动前设置 Redis 连接字符串环境变量 `REDIS_CONN_STRING`。 
5. 如果想使用 MySQL，需要先登录 MySQL 创建一个空的数据库 `gofile`，然后设置 `SQL_DSN` 环境变量即可，例如：`root:123456@tcp(localhost:3306)/gofile`。
6. 修改默认的 SQLite 数据库文件的位置，请设置 `SQLITE_PATH` 环境变量。
7. 设置会话密钥（默认随机生成），请设置 `SESSION_SECRET` 环境变量。
8. 设置文件上传路径（默认为工作目录下面的 `upload` 目录），请设置 `UPLOAD_PATH` 环境变量。
9. 禁止自动打开浏览器，启动时请指定 `no-browser` 参数：`./go-file.exe --no-browser true`。
10. 如果想要使用 Token 访问 API，请先前往个人账户管理页面生成 Token，之后在请求时加上 `Authorization` HTTP 头部，值为 `YOUR_TOKEN` 或者 `Bearer YOUR_TOKEN`。

**如果你不知道怎么加参数：**
1. 打开 go-file 所在的文件夹，
2. 按住 shift 并右键空白区域，
3. 选择`在此处打开 PowerShell`（如果是 Windows 11 的话则需要先点击`显示更多选项`），
4. 在打开的终端中输入：`./go-file --port 80 --video ./path/to/video`

**注意：**
1. 如果主机有多个 ip 地址，请使用 host 参数指定一个其他设备可访问的 ip 地址，如：`go-file.exe --host xxx.xxx.xxx.xxx`，否则二维码将生成错误。
2. 默认配置下访客可以上传和下载文件，可在 `管理` -> `系统设置` 中修改权限配置。
3. 如果是公网部署，务必记得第一时间更改默认密码！ 

## 演示
在线试用（用户名为 `admin`，密码为 `123456`）：https://go-file.herokuapp.com/

注意，以下展示图片可能没有得到及时跟新。
![index page](https://user-images.githubusercontent.com/39998050/178138784-2fc53a83-917d-4d2e-9aad-6c6c796bd9c8.png)
![file page](https://user-images.githubusercontent.com/39998050/178138792-1d9256f2-2ada-43c4-b646-28a93a919596.png)
![image page](https://user-images.githubusercontent.com/39998050/178138803-2a4da042-c29a-47c5-9e71-ebfac02cdf48.png)
![video page](https://user-images.githubusercontent.com/39998050/177032588-8946abde-a8da-45a2-a389-c16dba9cea34.png)
![setting page](https://user-images.githubusercontent.com/39998050/178138817-3f9caf95-ffc9-45fe-b2af-32c4a2e7b085.png)
![setting page 2](https://user-images.githubusercontent.com/39998050/178138833-d10e6f5a-aeea-4af3-8ae1-c0b3ab1d92f7.png)

## 其他
[Node.js 版本在此](https://github.com/songquanpeng/lan-share)
