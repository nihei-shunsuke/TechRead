// イベント作成のPOST関数
function create_event() {
    // 入力フォームの内容を取得
    const event_name = document.getElementById('event_name').value;
    const book_name = document.getElementById('book_name').value;
    const members = document.getElementsByClassName('member');
    let member_list = [];
    for(let i = 0; i < members.length; i++) {
        member_list.push(members.item(i).value);
    }

    // 入力内容を出力
    console.log(event_name);
    console.log(book_name);
    console.log(member_list);

    // 入力値をリクエスト
    const parameter = {
        method:'POST',
        mode:'no-cors',
        headers:{
            'Content-Type':'application/json'
        },
        body: JSON.stringify({
            event_name: event_name,
            book_name: book_name,
            member_list: member_list,
        })
    };
    const url = 'http://localhost:8080/create-event';

    fetch(url, parameter).then((response) => {
        console.log(response);
    })
    .catch(err => {
        console.error(err);
    })
}

const create_btn = document.getElementById('create-btn');
create_btn.addEventListener('click', create_event, false);

// メンバーを追加するフォームを追加する関数
function add_member_form() {
    const newForm = document.createElement('input');
    newForm.type = 'text';
    newForm.className = 'member';

    const br = document.createElement('br');
    document.getElementById('add-member').appendChild(br);
    document.getElementById('add-member').appendChild(newForm);
};

const add_member_button = document.getElementById('add-member-button');
add_member_button.onclick = add_member_form;
