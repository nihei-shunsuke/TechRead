function btn1(){    
    // 入力フォームの内容を取得
    const title = document.getElementById('title').value;
    const word = document.getElementById('word').value;
    // 入力内容を出力
    console.log(title);
    console.log(word);
}

const btn = document.getElementById('btn');

btn.onclick = btn1;