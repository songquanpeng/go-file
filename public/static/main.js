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
    let prompt = "";
    let files = document.getElementById('fileInput').files;
    if (files.length === 1) {
        prompt = 'Selected file: ' + files[0].name;
    } else {
        prompt = files.length + " files selected";
    }
    document.getElementById('uploadFileDialogTitle').innerText = prompt;
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
    document.addEventListener('DOMContentLoaded', () => {
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
    });
}

init();