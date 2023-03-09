function btn1(){    
    // 入力フォームの内容を取得
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    // 入力内容を出力
    console.log(username);
    console.log(email);
    console.log(password);

    const parameter = {
        method:'POST',
        mode:'no-cors',
        headers:{
            'Content-Type':'application/json'
        },
        body: JSON.stringify({
            user_name: username,
            email: email,
            password: password
        })
    };
    fetch('http://localhost:8080/sign-up', parameter).then((response) => {
        console.log(response);
    })
    .catch(err => {
        console.error(err);
    })
    // console.log(r);
}

const btn = document.getElementById('btn');

btn.onclick = btn1;