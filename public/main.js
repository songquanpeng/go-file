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
    let url = window.location.origin + link;
    url = encodeURI(url)
    console.log(url)
    let qr = new QRious({
        element: document.getElementById('qrcode'),
        value: url,
        size: 200,
    });
}