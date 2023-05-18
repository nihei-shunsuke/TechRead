

const create_chapter = async () => {
  // 入力値を取り出す
  const dateString = await document.getElementById("date").value;
  const place = await document.getElementById("place").value;
  const chapterStr = await document.getElementById("chapter").value;
  const event_id = await 1; // クエリにevent-idを入れておく。それを格納する
  const word = await document.getElementById("word").value;

  // 入力値を出力
  const date = await new Date(dateString);
  const chapter = await Number(chapterStr);
  console.log(date);
  console.log(place);
  console.log(chapter);
  console.log(word);

  // 入力をリクエスト
  const parameter = await {
    method: 'POST',
    headers: {
      'Content-Type':'application/json'
    },
    body: JSON.stringify({
      chapter_num: chapter,
      venue: place,
      event_id: event_id,
      content: word,
      event_date: date,
    })
  };

  const url = 'http://localhost:8080/create-chapter';

  fetch(url, parameter).then((response) => {
    console.log(response);
  })
  .catch(err => {
    console.err(err);
  });

  const response = await fetch("http://localhost:8080/create-chapter", parameter);
  const res = await response.json();
  console.log(res);

  if(res.res_state == 'success'){
    location.href='../html/chapter-list.html';
  };
};

const create_chapter_btn = document.getElementById('create-chapter-btn');
create_chapter_btn.onclick = create_chapter;
