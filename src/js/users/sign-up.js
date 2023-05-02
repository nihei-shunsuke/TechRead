<<<<<<< Updated upstream
const postSignUp = async () => {
  // 入力フォームの内容を取得
  const username = await document.getElementById('username').value;
  const email = await document.getElementById('email').value;
  const password = await document.getElementById('password').value;

  // 入力内容を出力
  console.log(username);
  console.log(email);
  console.log(password);

  const parameter = await {
    method:'POST',
    cache: "no-cache",
    headers:{
        'Content-Type':'application/json'
    },
    body: JSON.stringify({
        user_name: username,
        email: email,
        password: password
    })
  };

  const response = await fetch("http://localhost:8080/sign-up", parameter);
  const res = await response.json();
  console.log(res);

  if (res.res_state == 'success') {
    window.location.href = '../html/event-list.html';
  };
}
=======
const postSignUp = async () => {
  // 入力フォームの内容を取得
  const username = await document.getElementById('username').value;
  const email = await document.getElementById('email').value;
  const password = await document.getElementById('password').value;

  // 入力内容を出力
  console.log(username);
  console.log(email);
  console.log(password);

  const parameter = await {
    method:'POST',
    cache: "no-cache",
    headers:{
        'Content-Type':'application/json'
    },
    body: JSON.stringify({
        user_name: username,
        email: email,
        password: password
    })
  };

  const response = await fetch("http://localhost:8080/sign-up", parameter);
  const res = await response.json();
  console.log(res);

  if (res.res_state == 'success') {
    Cookies.set('user_id',res.user_id);
    window.location.href = '../html/event-list.html';
  } else {
    alert('メールアドレスが重複しています');
  };
}
>>>>>>> Stashed changes
