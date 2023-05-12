let hiddenTextArea = undefined;

function showUploadModal() {
    if (location.href.split('/')[3].startsWith("explorer")) {
        let path = getPathParam();
        document.getElementById('uploadFileDialogTitle').innerText = `上传文件到 "${path}"`;
    }
    showModal('uploadModal');
}

function getPathParam() {
    let url = new URL(location.href);
    let searchParams = new URLSearchParams(url.search);
    let path = "/";
    if (searchParams.has('path')) {
        path = searchParams.get('path');
    }
    if (path === "") path = "/";
    return path;
}

function closeUploadModal() {
    document.getElementById('uploadModal').className = "modal";
}


function showModal(id) {
    document.getElementById(id).className = "modal is-active";
}

function closeModal(id) {
    document.getElementById(id).className = "modal";
}

function onChooseBtnClicked(e) {
    document.getElementById('fileInput').click();
    e.preventDefault();
}

function deleteFile(id, link) {
    fetch("/api/file", {
        method: 'delete',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            id: id,
            link: link
        })
    }).then(function (res) {
        res.json().then(function (data) {
            // showMessage(data.message);
            if (!data.success) {
                console.error(data.message);
                showMessage(data.message, true);
            } else {
                document.getElementById("file-" + id).style.display = 'none';
                showToast(`文件删除成功：${link}`)
            }
        })
    });
}

function deleteImage() {
    let e = document.getElementById("inputDeleteImage");
    if (e.value === "") return;
    let tmpList = e.value.split("/");
    let filename = tmpList[tmpList.length - 1];

    fetch("/api/image", {
        method: 'delete',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            filename: filename,
        })
    }).then(function (res) {
        res.json().then(function (data) {
            if (data.success) {
                e.value = "";
                showToast("图片已成功删除");
            } else {
                showToast(data.message, "danger");
            }
        })
    });
}


function updateDownloadCounter(id) {
    let e = document.getElementById(id);
    let n = parseInt(e.innerText.replace("次下载", ""));
    e.innerText = `${n + 1} 次下载`;
}

function onFileInputChange() {
    let prompt;
    let files = document.getElementById('fileInput').files;
    if (files.length === 1) {
        prompt = '已选择文件: ' + files[0].name;
    } else {
        prompt = `已选择 ${files.length} 个文件`;
    }
    document.getElementById('uploadFileDialogTitle').innerText = prompt;
}

function byte2mb(n) {
    let sizeMb = 1024 * 1024;
    n /= sizeMb;
    return n.toFixed(2);
}

function uploadFile() {
    let fileUploadCard = document.getElementById('fileUploadCard');
    let fileUploadTitle = document.getElementById('fileUploadTitle');
    let fileUploadProgress = document.getElementById('fileUploadProgress');
    let fileUploadDetail = document.getElementById('fileUploadDetail');
    fileUploadCard.style.display = 'block';
    let files = document.getElementById('fileInput').files;
    let description = document.getElementById("fileUploadDescription").value;
    if (files.length === 0 && description === "") {
        return;
    }
    closeUploadModal();
    let formData = new FormData();
    for (let i = 0; i < files.length; i++) {
        formData.append("file", files[i]);
    }
    formData.append("description", description);

    let path = "";
    if (location.href.split('/')[3].startsWith("explorer")) {
        path = getPathParam();
    }
    formData.append("path", path);

    fileUploadTitle.innerText = `正在上传 ${files.length} 个文件`;

    let fileUploader = new XMLHttpRequest();
    fileUploader.upload.addEventListener("progress", ev => {
        let percent = (ev.loaded / ev.total) * 100;
        fileUploadProgress.value = Math.round(percent);
        fileUploadDetail.innerText = `处理中 ${byte2mb(ev.loaded)} MB / ${byte2mb(ev.total)} MB...`
    }, false);
    fileUploader.addEventListener("load", ev => {
        fileUploadTitle.innerText = `已上传 ${files.length} 个文件`;
        if (fileUploader.status === 403) {
            location.href = "/login";
        } else {
            location.reload();
        }
        // setTimeout(()=>{
        //     fileUploadCard.style.display = 'none';
        // }, 5000);
    }, false);
    fileUploader.addEventListener("error", ev => {
        if (fileUploader.status === 403) {
            location.href = "/login";
        } else {
            fileUploadTitle.innerText = `文件上传失败`;
        }
        console.error(ev);
    }, false);
    fileUploader.addEventListener("abort", ev => {
        fileUploadTitle.innerText = `文件上传已终止`;
    }, false);
    fileUploader.open("POST", "/api/file");
    fileUploader.send(formData);
}

function dropHandler(ev) {
    ev.preventDefault();
    document.getElementById('fileInput').files = ev.dataTransfer.files;
    onFileInputChange();
}

function dragOverHandler(ev) {
    document.getElementById('uploadFileDialogTitle').innerText = "释放文件至此对话框";
    ev.preventDefault();
}

function imageDropHandler(ev) {
    ev.preventDefault();
    document.getElementById('fileInput').files = ev.dataTransfer.files;
    uploadImage();
}

function uploadImage() {
    document.getElementById("promptBox").style.display = "block";
    let imageUploadProgress = document.getElementById('imageUploadProgress');
    let imageUploadStatus = document.getElementById('imageUploadStatus');
    imageUploadStatus.innerText = "上传中..."

    let files = document.getElementById('fileInput').files;
    let formData = new FormData();
    for (let i = 0; i < files.length; i++) {
        if (files[i]['type'].split('/')[0] === 'image') {
            formData.append("image", files[i]);
        }
    }

    let fileUploader = new XMLHttpRequest();
    fileUploader.upload.addEventListener("progress", ev => {
        let percent = (ev.loaded / ev.total) * 100;
        imageUploadProgress.value = Math.round(percent);
    }, false);
    fileUploader.addEventListener("load", ev => {
        // Uploading is done.
        if (fileUploader.status === 200) {
            imageUploadStatus.innerText = "文件上传成功";
        } else if (fileUploader.status === 403) {
            location.href = "/login";
        }
    }, false);
    fileUploader.addEventListener("error", ev => {
        imageUploadStatus.innerText = "文件上传失败";
        console.error(ev);
    }, false);
    fileUploader.addEventListener("abort", ev => {
        imageUploadStatus.innerText = "文件上传终止";
    }, false);
    fileUploader.addEventListener("readystatechange", ev => {
        if (fileUploader.readyState === 4) {
            let res = JSON.parse(fileUploader.response);
            console.log(res);
            if (fileUploader.status === 200) {
                let filenames = res.data;
                let imageUploadPanel = document.getElementById('imageUploadPanel');
                filenames.forEach(filename => {
                    let url = location.href + '/' + filename;
                    imageUploadPanel.insertAdjacentHTML('afterbegin', `
                <div class="field has-addons">
                    <div class="control is-light is-expanded">
                        <input class="input url-input" type="text" value="${url}" readonly>
                    </div>
                    <div class="control">
                        <a class="button is-light" onclick="copyText('${url}')">
                            复制链接
                        </a>
                    </div>
                    <div class="control">
                        <a class="button is-light" onclick="copyText('![${filename}](${url})')">
                            复制 Markdown 代码
                        </a>
                    </div>
                </div>
                `);
                });
            } else if (fileUploader.status === 403) {
                location.href = "/login";
            }
        }
    });
    fileUploader.open("POST", "/api/image");
    fileUploader.send(formData);
}

function imageDragOverHandler(ev) {
    ev.preventDefault();
}

function showMessage(message, isError = false) {
    const messageToast = document.getElementById('messageToast');
    messageToast.style.display = 'block';
    messageToast.className = isError ? "message is-danger" : "message";
    let timeout = isError ? 5000 : 2000;
    document.getElementById('messageToastText').innerText = message;
    if (isError) {
        document.getElementById("nav").scrollIntoView();
    }
    setTimeout(function () {
        messageToast.style.display = 'none';
    }, timeout);
}

function showQRCode(link) {
    let url = window.location.origin + link;
    url = encodeURI(url)
    console.log(url)
    let qr = new QRious({
        element: document.getElementById('qrcode'),
        value: url,
        size: 200,
    });
    showModal('qrcodeModal');
}

function copyLink(link) {
    let url = window.location.origin + link;
    url = decodeURI(url);
    copyText(url);
    showToast(`已复制：${url}`, 'success');
}

function toLocalTime(str) {
    let date = Date.parse(str);
    return date.toLocaleString()
}

function copyText(text) {
    const textArea = document.getElementById('hiddenTextArea');
    textArea.textContent = text;
    document.body.append(textArea);
    textArea.select();
    document.execCommand("copy");
}

function showToast(message, type = "success", timeout = 2900) {
    let toast = document.getElementById("toast");
    toast.innerText = message;
    toast.className = `show notification is-${type}`;
    setTimeout(() => {
        toast.className = "";
    }, timeout);
}

function showGeneralModal(title, content) {
    document.getElementById("generalModalTitle").innerText = title;
    document.getElementById("generalModalContent").innerHTML = content;
    showModal("generalModal");
}

async function loadOptions() {
    let tab = document.getElementById('settingTab');
    let html = ""
    let response = await fetch("/api/option");
    let result = await response.json();
    if (result.success) {
        for (let i = 0; i < result.data.length; i++) {
            let key = result.data[i].key;
            let value = result.data[i].value;
            html += `
            <div>
                <label class="label">${key}</label>
                <div class="field has-addons">
                    <p class="control is-expanded">
                        <input class="input" id="inputOption${key}" type="text" placeholder="请输入新的配置" value="${value}">
                    </p>
                    <p class="control">
                        <a class="button" onclick="updateOption('${key}', 'inputOption${key}')">提交</a>
                    </p>
                </div>
            </div>`;
        }
    } else {
        html = `<p>选项加载失败：${result.message}</p>`
    }
    tab.innerHTML = html;
}

async function updateOption(key, inputElementId, originValue = "") {
    let inputElement = document.getElementById(inputElementId);
    let value = inputElement.value;
    let response = await fetch("/api/option", {
        method: "PUT",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            key: key,
            value: value
        })
    });
    let result = await response.json();
    if (result.success) {
        showToast(`更新成功`, "success");
    } else {
        showToast(`更新失败：${result.message}`, "danger");
        if (originValue !== "") {
            inputElement.value = originValue;
        }
    }
}

async function updateUser(key, inputElementId) {
    let inputElement = document.getElementById(inputElementId);
    let value = inputElement.value;
    if (value === "") return
    let data = {};
    data[key] = value;
    let response = await fetch("/api/user", {
        method: "PUT",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    let result = await response.json();
    if (result.success) {
        showToast(`更新信息成功`, "success");
    } else {
        showToast(`更新信息失败：${result.message}`, "danger");
    }
}

async function createUser() {
    let username = document.getElementById("newUserName").value;
    let password = document.getElementById("newUserPassword").value;
    if (!username || !password) return;
    let type = document.getElementById("newUserType").value;
    let data = {
        username: username,
        password: password,
        type: type
    }
    let response = await fetch("/api/user", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    let result = await response.json();
    if (result.success) {
        showToast(`添加用户成功`, "success");
    } else {
        showToast(`添加用户失败：${result.message}`, "danger");
    }
}

async function manageUser() {
    let username = document.getElementById("manageUserName").value;
    let action = document.getElementById("manageAction").value;
    if (!username) return;

    let data = {
        username: username,
        action: action,
    }
    let response = await fetch("/api/manage_user", {
        method: "PUT",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
    let result = await response.json();
    if (result.success) {
        showToast(`操作成功`, "success");
    } else {
        showToast(`操作失败：${result.message}`, "danger");
    }
}

async function generateNewToken() {
    let response = await fetch("/api/token", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        }
    });
    let result = await response.json();
    if (result.success) {
        showToast(`Token 已重置为：${result.data}`, "success");
    } else {
        showToast(`操作失败：${result.message}`, "danger");
    }
}

function isMobile() {
    return window.innerWidth <= 768;
}

function getFileExt(link) {
    let parts = link.split('.');
    if (parts.length === 1) return "";
    return parts[parts.length - 1].toLowerCase();
}

function getFilename(link) {
    let parts = link.split('/');
    return parts[parts.length - 1];
}

function displayFile(link) {
    // TODO: text file preview support
    let ext = getFileExt(link);
    let filename = getFilename(link);
    console.log(link, ext, filename)
    document.getElementById("displayModalTitle").innerText = filename;
    if (ext === "mp3" || ext === "wav" || ext === "ogg") {
        document.getElementById("displayModalContent").innerHTML = `
        <audio controls>
            <source src="${link}" type="audio/${ext}">
        </audio>`;
    } else if (ext === "mp4" || ext === "webm" || ext === "ogv") {
        document.getElementById("displayModalContent").innerHTML = `
        <video controls style="width: 100%">
            <source src="${link}" type="video/${ext}">
        </video>`;
    } else if (ext === "png" || ext === "jpg" || ext === "jpeg" || ext === "gif") {
        document.getElementById("displayModalContent").innerHTML = `
        <img src="${link}" alt="${filename}" width="100%">`;
    } else if (ext === "pdf") {
        if (isMobile()) {
            window.open(link);
            return;
        }
        document.getElementById("displayModalContent").innerHTML = `
        <div style="width:100%; height: 600px!important;">
            <iframe src="${link}" width="100%" height="100%"></iframe>
        </div>`;
    } else {
        window.open(link);
        return;
    }
    showModal("displayModal");
}

function init() {
    const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);
    if ($navbarBurgers.length > 0) {
        $navbarBurgers.forEach(el => {
            el.addEventListener('click', () => {
                const target = el.dataset.target;
                const $target = document.getElementById(target);
                el.classList.toggle('is-active');
                $target.classList.toggle('is-active');
            });
        });
    }

    hiddenTextArea = document.createElement('textarea');
    hiddenTextArea.setAttribute("id", "hiddenTextArea");
    hiddenTextArea.style.cssText = "height: 0px; width: 0px";
    document.body.appendChild(hiddenTextArea);
}

document.addEventListener('DOMContentLoaded', init)