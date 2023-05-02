<<<<<<< Updated upstream


function create_chapter() {
  // 入力値を取り出す
  const date = document.getElementById("date").value;
  const place = document.getElementById("place").value;
  const chapter = document.getElementById("chapter").value;
  const event_id = 1; // クエリにevent-idを入れておく。それを格納する
  const word = document.getElementById("word").value;
  const last_updater_id = 1; // Cookieから取り出す

  // 入力値を出力
  console.log(date);
  console.log(place);
  console.log(chapter);
  console.log(word);

  // 入力をリクエスト
  const parameter = {
    method: 'POST',
    mode: 'no-cors',
    headers: {
      'Content-Type':'application/json'
    },
    body: JSON.stringify({
      chapter_num: chapter,
      venue: place,
      event_id: event_id,
      content: word,
      last_updater_id: last_updater_id,
      event_date: date,
    })
  };
  const url = 'http://localhost:8080/create-chapter';

  fetch(url, parameter).then((response) => {
    console.log(response);
  })
  .catch(err => {
    console.err(err);
  })
};

const create_chapter_btn = document.getElementById('create-chapter-btn');
create_chapter_btn.onclick = create_chapter;
=======


function create_chapter() {
  // 入力値を取り出す
  const date = document.getElementById("date").value;
  const place = document.getElementById("place").value;
  const chapter = document.getElementById("chapter").value;
  const event_id = 1; // クエリにevent-idを入れておく。それを格納する
  const word = document.getElementById("word").value;
  const last_updater_id = 1; // Cookieから取り出す

  // 入力値を出力
  console.log(date);
  console.log(place);
  console.log(chapter);
  console.log(word);

  // 入力をリクエスト
  const parameter = {
    method: 'POST',
    mode: 'no-cors',
    headers: {
      'Content-Type':'application/json'
    },
    body: JSON.stringify({
      chapter_num: chapter,
      venue: place,
      event_id: event_id,
      content: word,
      last_updater_id: last_updater_id,
      event_date: date,
    })
  };
  const url = 'http://localhost:8080/create-chapter';

  fetch(url, parameter).then((response) => {
    console.log(response);
  })
  .catch(err => {
    console.err(err);
  })

  location.href='../html/chapter-list.html';
};

const create_chapter_btn = document.getElementById('create-chapter-btn');
create_chapter_btn.onclick = create_chapter;
>>>>>>> Stashed changes
