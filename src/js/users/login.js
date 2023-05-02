<<<<<<< Updated upstream
const postLogin = async () => {
  // 入力フォームの内容を取得
  const email = await document.getElementById('email').value;
  const password = await document.getElementById('password').value;

  // 入力内容を出力
  console.log(email);
  console.log(password);

  const parameter = await {
    method: 'POST',
    cache: "no-cache",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      email: email,
      password: password
    })
  };

  const response = await fetch("http://localhost:8080/login", parameter);
  const res = await response.json();
  console.log(res);

  if (res.res_state == 'success') {
    window.location.href = '../html/event-list.html';
  };
}
=======
const postLogin = async () => {
  // 入力フォームの内容を取得
  const email = await document.getElementById('email').value;
  const password = await document.getElementById('password').value;

  // 入力内容を出力
  console.log(email);
  console.log(password);

  const parameter = await {
    method: 'POST',
    cache: "no-cache",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      email: email,
      password: password
    })
  };

  const response = await fetch("http://localhost:8080/login", parameter);
  const res = await response.json();
  console.log(res);

  if (res.res_state == 'success') {
    Cookies.set('user_id',res.user_id);
    window.location.href = '../html/event-list.html';
  } else {
    alert('パスワードもしくはメールアドレスが間違っています');
  };
}
>>>>>>> Stashed changes
