function showUploadModal() {
    document.getElementById('uploaderNameInput').value = localStorage.getItem('uploaderName');
    showModal('uploadModal');
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
        }).then(function (res) {
            res.json().then(function (data) {
                // showMessage(data.message);
                if (!data.success) {
                    console.error(data.message);
                    showMessage(data.message, true);
                    localStorage.removeItem('token');
                    askUserInputToken();
                } else {
                    document.getElementById("file-" + id).style.display = 'none';
                }
            })
        });
    }
}


function onFileInputChange() {
    let prompt;
    let files = document.getElementById('fileInput').files;
    if (files.length === 1) {
        prompt = 'Selected file: ' + files[0].name;
    } else {
        prompt = files.length + " files selected";
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
    closeUploadModal();
    let files = document.getElementById('fileInput').files;
    let formData = new FormData();
    for (let i = 0; i < files.length; i++) {
        formData.append("file", files[i]);
    }
    if (files.length === 1) {
        fileUploadTitle.innerText = `Uploading 1 file`;
    } else {
        fileUploadTitle.innerText = `Uploading ${files.length} files`;
    }

    let fileUploader = new XMLHttpRequest();
    fileUploader.upload.addEventListener("progress", ev => {
        let percent = (ev.loaded / ev.total) * 100;
        fileUploadProgress.value = Math.round(percent);
        fileUploadDetail.innerText = `Processing ${byte2mb(ev.loaded)} MB / ${byte2mb(ev.total)} MB.`
    }, false);
    fileUploader.addEventListener("load", ev => {
        fileUploadTitle.innerText = files.length === 1 ? `File uploaded.` : `Files uploaded.`;
        // setTimeout(()=>{
        //     fileUploadCard.style.display = 'none';
        // }, 5000);
    }, false);
    fileUploader.addEventListener("error", ev => {
        fileUploadTitle.innerText = `File uploading failed.`;
        console.error(ev);
    }, false);
    fileUploader.addEventListener("abort", ev => {
        fileUploadTitle.innerText = `File uploading aborted.`;
    }, false);
    fileUploader.open("POST", "/upload");
    fileUploader.send(formData);
}

function dropHandler(ev) {
    ev.preventDefault();
    document.getElementById('fileInput').files = ev.dataTransfer.files;
    onFileInputChange();
}

function dragOverHandler(ev) {
    document.getElementById('uploadFileDialogTitle').innerText = "Release to this dialog";
    ev.preventDefault();
}


function updateToken() {
    let token = document.getElementById('tokenInput').value;
    token = token.trim();
    localStorage.setItem('token', token);
    closeModal('tokenModal');
}

function askUserInputToken() {
    showModal('tokenModal');
}

function onUploaderNameChange() {
    localStorage.setItem('uploaderName', document.getElementById('uploaderNameInput').value);
}

function showMessage(message, isError = false) {
    const messageToast = document.getElementById('messageToast');
    messageToast.style.display = 'block';
    messageToast.className = isError ? "message is-danger" : "message";
    let timeout = isError ? 5000 : 2000;
    document.getElementById('messageToastText').innerText = message;
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

function toLocalTime(str) {
    let date = Date.parse(str);
    return date.toLocaleString()
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
}

document.addEventListener('DOMContentLoaded', init)