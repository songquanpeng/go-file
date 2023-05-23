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
  <a href="https://hub.docker.com/repository/docker/justsong/go-file">
    <img src="https://img.shields.io/docker/pulls/justsong/go-file?color=brightgreen" alt="docker pull">
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
  <a href="https://github.com/songquanpeng/gofile-launcher">启动器下载</a>
  ·
  <a href="https://github.com/songquanpeng/gofile-cli">CLI 下载</a>
  ·
  <a href="https://iamazing.cn/page/LAN-SHARE-使用教程">使用教程</a>
  ·
  <a href="#演示">截图展示</a>
</p>

> **Note**：推荐使用官方的 [Go File 启动器](https://github.com/songquanpeng/gofile-launcher)，免去命令行操作。

## 特点
1. 无需配置环境，仅单个可执行文件，**直接双击即可开始使用**。
2. 自动打开浏览器，分享文件快人一步。
3. 提供**二维码**，可供移动端扫描下载文件，告别手动输入链接。
4. 支持**分享本地文件夹**。
5. 适配移动端。
6. 内置**图床**，支持直接粘贴上传图片，提供图片上传 API。
7. 内置**视频播放**页面，可用于在其他设备上在线博客自己电脑上的视频，轻松跨设备在线看视频。
8. 支持**拖拽上传，拷贝上传**。
9. 允许对不同类型的用户设置文件访问权限限制。
10. 访问频率限制。
11. 支持 Token API 验证，便于与其他系统整合。
12. 为不熟悉命令行的用户制作了**启动器**，[详见此处](https://github.com/songquanpeng/gofile-launcher)。
13. **支持 PicGo**，插件搜索 `gofile` 即可安装，[详见此处](https://github.com/songquanpeng/picgo-plugin-gofile)。
14. 配套 CLI 工具，支持命令行上传文件，支持 P2P 模式文件分享，[详见此处](https://github.com/songquanpeng/gofile-cli)。
15. Docker 一键部署：`docker run -d --restart always -p 3000:3000 -e TZ=Asia/Shanghai -v /home/ubuntu/data/go-file:/data justsong/go-file`

## 使用方法
> v0.3.3 以及之前版本的使用方法请[点击此处](https://github.com/songquanpeng/go-file/tree/52e8303e33e99bbcaf583d2d5a5bb0ec197bc676#使用方法)。

直接双击即可使用，默认端口为 `3000`，程序在第一次启动时会自动创建管理员账户，用户名为 `admin`，密码为 `123456`，记得登录后到 `管理页面` 下的 `账户管理` 标签页中更改你的用户密码。

之后程序将自动为你打开浏览器，点击右上角的 `上传` 按钮即可上传，支持拖放上传，支持同时上传多个文件。

**进阶使用：**
1. 如果要修改端口，启动时请指定 `port` 参数：`./go-file.exe --port 80`。
2. 如果需要分享文件夹，启动时请指定 `path` 参数：`./go-file.exe --path ./this/is/a/path`，之后点击导航栏上的 `文件` 即可。
3. 如果需要分享本地的视频资源，加 `video` 参数：`./go-file.exe --video ./this/is/a/path`，之后点击导航栏上的 `视频` 即可。
4. 如果需要启用访问速率控制，需要在启动前设置 Redis 连接字符串环境变量 `REDIS_CONN_STRING`，例如：`redis://default:redispw@localhost:49153`。 
5. 如果想使用 MySQL，需要先登录 MySQL 创建一个空的数据库 `gofile`，然后设置 `SQL_DSN` 环境变量即可，例如：`root:123456@tcp(localhost:3306)/gofile`。
6. 修改默认的 SQLite 数据库文件的位置，请设置 `SQLITE_PATH` 环境变量，默认在工作目录下，名称为 `go-file.db`。
7. 设置会话密钥（默认随机生成），请设置 `SESSION_SECRET` 环境变量。
8. 设置文件上传路径（默认为工作目录下面的 `upload` 目录），请设置 `UPLOAD_PATH` 环境变量。
9. 禁止自动打开浏览器，启动时请指定 `no-browser` 参数：`./go-file.exe --no-browser true`。
10. 如果想要使用 Token 访问 API，请先前往个人账户管理页面生成 Token，之后在请求时加上 `Authorization` HTTP 头部，值为 `YOUR_TOKEN` 或者 `Bearer YOUR_TOKEN`。
    + 例如作为 Typora 的 Image Uploader：[./script/typora.py](./script/typora.py)

**如果你不知道怎么加参数：**
1. 打开 go-file 所在的文件夹，
2. 按住 shift 并右键空白区域，
3. 选择`在此处打开 PowerShell`（如果是 Windows 11 的话则需要先点击`显示更多选项`），
4. 在打开的终端中输入：`./go-file --port 80 --video ./path/to/video`

建议直接使用[启动器](https://github.com/songquanpeng/gofile-launcher)。

**使用 Docker 进行部署：**
执行：`docker run -d --restart always -p 3000:3000 -e TZ=Asia/Shanghai -v /home/ubuntu/data/go-file:/data justsong/go-file`

数据将会保存在宿主机的 `/home/ubuntu/data/go-file` 目录。

**注意：**
1. 如果主机有多个 ip 地址，请使用 host 参数指定一个其他设备可访问的 ip 地址，如：`go-file.exe --host xxx.xxx.xxx.xxx`，否则二维码将生成错误。
2. 默认配置下访客可以上传和下载文件，可在 `管理` -> `系统设置` 中修改权限配置。
3. 如果是公网部署，务必记得第一时间更改默认密码！ 

## 演示
在线试用（用户名为 `admin`，密码为 `123456`）：https://go-file.onrender.com

注意，以下展示图片可能没有得到及时跟新。
![index page](https://user-images.githubusercontent.com/39998050/178138784-2fc53a83-917d-4d2e-9aad-6c6c796bd9c8.png)
![file page](https://user-images.githubusercontent.com/39998050/178138792-1d9256f2-2ada-43c4-b646-28a93a919596.png)
![image page](https://user-images.githubusercontent.com/39998050/178138803-2a4da042-c29a-47c5-9e71-ebfac02cdf48.png)
![video page](https://user-images.githubusercontent.com/39998050/177032588-8946abde-a8da-45a2-a389-c16dba9cea34.png)
![setting page](https://user-images.githubusercontent.com/39998050/178138817-3f9caf95-ffc9-45fe-b2af-32c4a2e7b085.png)
![setting page 2](https://user-images.githubusercontent.com/39998050/178138833-d10e6f5a-aeea-4af3-8ae1-c0b3ab1d92f7.png)

[启动器](https://github.com/songquanpeng/gofile-launcher)截图：

![launcher](https://raw.githubusercontent.com/songquanpeng/gofile-launcher/main/demo.png)

## 其他
[Node.js 版本在此](https://github.com/songquanpeng/lan-share)
