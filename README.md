<p align="center">
  <a href="https://github.com/songquanpeng/go-file"><img src="https://user-images.githubusercontent.com/39998050/108494937-1a573e80-72e3-11eb-81c3-5545d7c2ed6e.jpg" width="200" height="200" alt="go-file"></a>
</p>

<div align="center">

# Go File

_✨ 文件分享工具，仅单个可执行文件，开箱即用，可用于局域网内分享文件和文件夹，直接跑满本地带宽 ✨_  

</div>

<p align="center">
  <a href="https://raw.githubusercontent.com/songquanpeng/go-file/master/LICENSE">
    <img src="https://img.shields.io/github/license/songquanpeng/go-file" alt="license">
  </a>
  <a href="https://github.com/songquanpeng/go-file/releases/latest">
    <img src="https://img.shields.io/github/v/release/songquanpeng/go-file?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://goreportcard.com/report/github.com/songquanpeng/go-file">
  <img src="https://goreportcard.com/badge/github.com/songquanpeng/go-file" alt="GoReportCard">
  </a>
</p>

<p align="center">
  <a href="https://github.com/songquanpeng/go-file/projects/1">开发规划</a>
  ·
  <a href="https://github.com/songquanpeng/go-file/releases">下载</a>
  ·
  <a href="https://iamazing.cn/page/LAN-SHARE-使用教程">使用教程</a>
</p>


<details>
<summary><strong><i>Click here to expend the English readme</i></strong></summary>
<div>

## Description
File sharing tool, can be used to share files in a LAN.

## Features
1. No need to configure environment and there is only a single executable file.
2. Automatically open browser to make you share file more quickly.
3. Generate QR codes for your mobile phone to scan.
4. Easily share all the content of a local dir.

## Usage
Just double-click to use with default port `3000` and default token (used to verify identity when user try to delete files) `token`.

If you want to change the port and token, run it like this:`./go-file.exe -port 80 -token private`.

Your can also public a local path by providing a `path` like this : `./go-file.exe -path ./this/is/a/path` 

## Demo
![desktop](https://user-images.githubusercontent.com/39998050/130427067-80bf3cc5-5fee-488a-bea5-e323b9458064.png)
![explorer view](https://user-images.githubusercontent.com/39998050/144734218-d8969c22-f626-464d-b0c5-c32ec61b4e7d.png)
![mobile](https://user-images.githubusercontent.com/39998050/130427229-10da003f-8d9a-4591-b32c-efedbac419fb.png)
## Others
[Node.js version is here.](https://github.com/songquanpeng/lan-share)
</div>
</details>


## 特点
1. 无需配置环境，仅单个可执行文件，直接双击即可开始使用。
2. 自动打开浏览器，分享文件快人一步。
3. 提供二维码，可供移动端扫描下载文件，告别手动输入链接。
4. 支持分享本地文件夹。

## 使用方法
直接双击即可使用，默认端口为 `3000`，默认的 token（用于删除文件时验证身份）为 `token`。

注意，如果主机有多个 ip 地址，请使用 host 参数指定一个其他设备可访问的 ip 地址，如：`go-file.exe -host xxx.xxx.xxx.xxx`，否则其他设备将无法访问！

如果需要修改端口，加参数即可：`./go-file.exe -port 80 -token private`。

如果需要分享文件夹，加 `path` 参数：`./go-file.exe -path ./this/is/a/path`，之后点击导航栏上的 `Explorer` 即可。

## 演示
![desktop](https://user-images.githubusercontent.com/39998050/130427067-80bf3cc5-5fee-488a-bea5-e323b9458064.png)
![explorer view](https://user-images.githubusercontent.com/39998050/144734218-d8969c22-f626-464d-b0c5-c32ec61b4e7d.png)
![mobile](https://user-images.githubusercontent.com/39998050/130427229-10da003f-8d9a-4591-b32c-efedbac419fb.png)

## 其他
[Node.js 版本在此](https://github.com/songquanpeng/lan-share)
