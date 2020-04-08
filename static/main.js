function deleteFile(id, link) {
  let token = localStorage.getItem('token');
  if (token === undefined) {
    token = askUserInputToken();
  }
  $.ajax({
    url: `/delete`,
    type: 'POST',
    data: {
      id: id,
      link: link,
      token: token
    },
    success: function(result) {
      showMessage(result.message);
      if (!result.success) {
        localStorage.removeItem('token');
        askUserInputToken();
      } else {
        $(`#file-${id}`).hide();
      }
    }
  });
}

function askUserInputToken() {
  let token = prompt('Please input token for authentication');
  token = token.trim();
  localStorage.setItem('token', token);
  return token;
}

function showMessage(message) {
  $(document).ready(function() {
    const messageToast = $('#messageToast');
    messageToast.show();
    document.getElementById('messageToastText').innerText = message;
    messageToast.delay(1000).hide(300);
  });
}
