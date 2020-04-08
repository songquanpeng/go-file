function deleteFile(id) {
  let token = localStorage.getItem('token');
  if (token === undefined) {
    token = askUserInputToken();
  }
  $.ajax({
    url: `/`,
    type: 'DELETE',
    data: {
      id: id,
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
    console.log(message);
    document.getElementById('messageToastText').innerText = message;
    messageToast.delay(1000).hide(300);
  });
}
