package main

var HTMLTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Go File</title>
    <meta name="keywords" content="Go File">
    <meta name="description" content="LAN file sharing tool website. 局域网文件共享网站">
    <meta name="theme-color" content="#3F51B5"/>
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons"/>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap"/>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/mdui@0.4.3/dist/css/mdui.min.css">
    <script src="https://cdn.jsdelivr.net/npm/mdui@0.4.3/dist/js/mdui.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/qrious@4.0.2/dist/qrious.min.js"></script>
</head>
<body class="mdui-loaded mdui-drawer-body-left">
<script>
    function deleteFile(id, link) {
        let token = localStorage.getItem('token');
        if (!token) {
            token = askUserInputToken();
            if (token) {
                deleteFile(id, link);
            }
        } else {
           fetch("/delete", {
            method: 'post',
			headers: {
            	'Content-Type': 'application/json'
            },
			body: JSON.stringify({
				id: id,
                link: link,
                token: token
			})
        }).then(function(res) {
            res.json().then(function (data){
                showMessage(data.message);
                if (!data.success) {
                    localStorage.removeItem('token');
                    askUserInputToken();
                } else {
                    document.getElementById("file-"+id).style.display = 'none';
                }
            })
        });
        }
    }

    function askUserInputToken() {
        let token = prompt('Please input token for authentication');
        token = token.trim();
        localStorage.setItem('token', token);
        return token;
    }

    function showMessage(message) {
        const messageToast = document.getElementById('messageToast');
        messageToast.style.display = 'block';
        document.getElementById('messageToastText').innerText = message;
        setTimeout(function (){
            messageToast.style.display = 'none';
        }, 2000);
    }
    
    function showQRCode(link) {
        let url = window.location.href.slice(0, -1) + link;
        url = encodeURI(url)
        console.log(url)
		let qr = new QRious({
		  element: document.getElementById('qrcode'),
		  value: url,
		  size: 200,
		});
    }
</script>
<div class="mdui-appbar-with-toolbar mdui-theme-primary-indigo mdui-theme-accent-indigo">
    <div class="mdui-appbar mdui-appbar-fixed">
        <div class="mdui-toolbar mdui-color-white mdui-color-theme">
            <span mdui-drawer="{target: '.mc-drawer', swipe: true}"
                  class="mdui-btn mdui-btn-icon mdui-ripple mdui-ripple-white">
                <i class="mdui-icon material-icons">menu</i>
            </span>
            <a class="mdui-typo-headline" href="/">Go File</a>
            <div class="mdui-toolbar-spacer"></div>
            <a href="https://github.com/songquanpeng/go-file" target="_blank"
               class="mdui-btn mdui-btn-icon mdui-ripple mdui-ripple-white">
                <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg"
                     xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 36 36"
                     enable-background="new 0 0 36 36" xml:space="preserve" class="mdui-icon"
                     style="width: 24px;height:24px;">
        <path fill-rule="evenodd" clip-rule="evenodd" fill="#ffffff" d="M18,1.4C9,1.4,1.7,8.7,1.7,17.7c0,7.2,4.7,13.3,11.1,15.5
	c0.8,0.1,1.1-0.4,1.1-0.8c0-0.4,0-1.4,0-2.8c-4.5,1-5.5-2.2-5.5-2.2c-0.7-1.9-1.8-2.4-1.8-2.4c-1.5-1,0.1-1,0.1-1
	c1.6,0.1,2.5,1.7,2.5,1.7c1.5,2.5,3.8,1.8,4.7,1.4c0.1-1.1,0.6-1.8,1-2.2c-3.6-0.4-7.4-1.8-7.4-8.1c0-1.8,0.6-3.2,1.7-4.4
	c-0.2-0.4-0.7-2.1,0.2-4.3c0,0,1.4-0.4,4.5,1.7c1.3-0.4,2.7-0.5,4.1-0.5c1.4,0,2.8,0.2,4.1,0.5c3.1-2.1,4.5-1.7,4.5-1.7
	c0.9,2.2,0.3,3.9,0.2,4.3c1,1.1,1.7,2.6,1.7,4.4c0,6.3-3.8,7.6-7.4,8c0.6,0.5,1.1,1.5,1.1,3c0,2.2,0,3.9,0,4.5
	c0,0.4,0.3,0.9,1.1,0.8c6.5-2.2,11.1-8.3,11.1-15.5C34.3,8.7,27,1.4,18,1.4z">

        </path>
      </svg>
            </a>
        </div>
    </div>
    <div class="mc-drawer mdui-drawer">
        <div class="mdui-list">
            <a class="mdui-list-item mdui-ripple" href="/">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-blue">home</i>
                <div class="mdui-list-item-content">Home</div>
            </a>
            <a class="mdui-list-item mdui-ripple" href="https://iamazing.cn/page/LAN-SHARE-使用教程">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-deep-orange">help</i>
                <div class="mdui-list-item-content">Help</div>
            </a>
            <a class="mdui-list-item mdui-ripple" href="https://github.com/songquanpeng/go-file">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-purple">error_outline</i>
                <div class="mdui-list-item-content">About</div>
            </a>
            <a class="mdui-list-item mdui-ripple" href="https://github.com/songquanpeng/go-file/issues/new">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-green">message</i>
                <div class="mdui-list-item-content">Feedback</div>
            </a>
            <a class="mdui-list-item mdui-ripple" href="https://github.com/songquanpeng/go-file">
                <i class="mdui-list-item-icon mdui-icon material-icons mdui-text-color-yellow">star</i>
                <div class="mdui-list-item-content">Star</div>
            </a>
        </div>
    </div>
</div>
<div class="mdui-snackbar mdui-snackbar-bottom" id="messageToast" style="transform: translate(-50%, 0px);display: none" >
    <div class="mdui-snackbar-text" id="messageToastText"></div>
</div>
{{if ne .message ""}}
<script>
    showMessage({{.message}})
</script>
{{end}}

<div class="mdui-container">
    <div class="mdui-card" style="margin-top: 16px;padding-left: 8px;padding-right: 8px">
        <div class="mdui-textfield mdui-textfield-floating-label" style="padding-top: 8px;margin-bottom: 8px;margin-right: 8px">
            <i class="mdui-icon material-icons">search</i>
            <label class="mdui-textfield-label">Search files...</label>

            <input class="mdui-textfield-input" id="searchInput" type="text" autocomplete="off" autofocus style="cursor: text;">
            <div class="mdui-textfield-helper">Press "Enter" to search.</div>
        </div>
        <script>
            const input = document.getElementById("searchInput");
            input.addEventListener("keyup", function (event) {
                if (event.key === 'Enter') {
                    let value = input.value.trim();
                    if (value === "") {
                        location.href = "/";
                    }else {
                        location.href = "/?query="+value;
                    }
                }
            });
        </script>
    </div>

    {{range $index, $file := .files}}
        <div class="mdui-card" id="file-{{$file.Id}}" style="margin-top: 16px">
            <div class="mdui-card-primary" style="padding-top: 8px">
                <div class="mdui-card-primary-title">{{$file.Filename}}</div>
                <div class="mdui-card-primary-subtitle">
                    <span>{{$file.Uploader}}</span>
                    <span>{{$file.Time}}</span>
<!--                    <span>{{$file.DownloadCounter}}</span>-->
                </div>
            </div>
            <div class="mdui-card-content" style="padding-top: 8px;padding-bottom: 8px">
                {{$file.Description}}
            </div>
            <div class="mdui-card-actions">
                <a class="mdui-btn mdui-btn-icon mdui-float-right" download="{{$file.Filename}}" href="{{$file.Link}}"><i
                            class="mdui-icon material-icons mdui-text-color-green">cloud_download</i></a>
                <a class="mdui-btn mdui-btn-icon mdui-float-right" target="_blank" href="{{$file.Link}}"><i class="mdui-icon material-icons mdui-text-color-indigo">play_circle_filled</i></a>

<!--                <button class="mdui-btn mdui-btn-icon mdui-float-right"><i class="mdui-icon material-icons mdui-text-color-purple">thumb_down</i>-->
<!--                </button>-->
<!--                <button class="mdui-btn mdui-btn-icon mdui-float-right"><i-->
<!--                            class="mdui-icon material-icons mdui-text-color-blue">thumb_up</i>-->
<!--                </button>-->
                <button class="mdui-btn mdui-btn-icon mdui-float-right" onclick="showQRCode('{{$file.Link}}')" mdui-dialog="{target: '#qrcodeDialog'}"><i class="mdui-icon material-icons mdui-text-color-blue">camera_alt</i>
                </button>
                <button class="mdui-btn mdui-btn-icon mdui-float-right" onclick="deleteFile({{$file.Id}}, '{{$file.Link}}')"><i
                            class="mdui-icon material-icons mdui-text-color-red">delete</i>
                </button>
            </div>
        </div>
    {{end}}
</div>
<!--QR code-->
<div class="mdui-dialog" id="qrcodeDialog" style="width: 220px; height: 220px; padding: 10px">
	<canvas id="qrcode"></canvas>
</div>


<div class="mdui-fab-wrapper mdui-fab">
    <button class="mdui-fab mdui-ripple mdui-color-theme-accent" mdui-dialog="{target: '#uploadFileDialog'}">
        <i class="mdui-icon material-icons">add</i>
    </button>
</div>
<form method="post" action="/upload" enctype='multipart/form-data'>
    <div class="mdui-dialog" id="uploadFileDialog">
        <div class="mdui-dialog-title" id="uploadFileDialogTitle">Upload files</div>
        <div class="mdui-dialog-content">
            <input class="mdui-btn mdui-ripple" id="fileInput" type="file" name="file" multiple required style="display: none"
                   onchange="onFileInputChange()">
			<script>
			function onFileInputChange() {
			    let prompt = "";
			    let files = document.getElementById('fileInput').files;
			    if(files.length === 1) {
			        prompt = 'Selected file: '+ files[0].name;
			    } else {
			        prompt = files.length + " files selected";
			    }
			  	document.getElementById('uploadFileDialogTitle').innerText = prompt;
			}
			</script>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <i class="mdui-icon material-icons">account_circle</i>
                <label class="mdui-textfield-label">Your name (optional)</label>
                <textarea class="mdui-textfield-input" name="uploader" id="fileUploader"></textarea>
            </div>
            <div class="mdui-textfield mdui-textfield-floating-label">
                <i class="mdui-icon material-icons">textsms</i>
                <label class="mdui-textfield-label">Description (optional)</label>
                <textarea class="mdui-textfield-input" name="description" id="fileDescription"></textarea>
            </div>
        </div>
        <div class="mdui-dialog-actions">
            <button class="mdui-btn mdui-ripple" onclick="document.getElementById('fileInput').click()">choose</button>
            <button class="mdui-btn mdui-ripple" mdui-dialog-close>cancel</button>
            <input class="mdui-btn mdui-ripple" type="submit" name="SUBMIT">
        </div>
    </div>
</form>
</body>
</html>
`
