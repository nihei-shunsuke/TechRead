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
        mode:'cors',
        headers:{
            'Content-Type':'application/json'
        },
        body: JSON.stringify({
            user_name: username,
            email: email,
            password: password

        })
    }
    const r = fetch('localhost:8080/sign-up', parameter).then((response) => {
        return response.json();
    })
    .catch(error => {
        console.error('失敗', error)
    })
    console.log(r)
}

const btn = document.getElementById('btn');

btn.onclick = btn1;