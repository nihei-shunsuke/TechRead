# TechRead

## 使用ライブラリ
- gorilla: Routing, Middleware,Session

## 起動方法
### フロントエンド側
- Cookieに関わる処理が必要な場合、localhostを立ち上げる。Cookieを扱うには、サーバを立てないとできないらしい。
- 方法はなんでもいいが、手軽なのはVSCodeの[Live Server](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer)という拡張機能。
- ステータスバーのGo Liveボタンを押すと勝手にlocalhostが立ち上がります。

### サーバーサイド側
#### 初回 or Dockerfileに変更があった場合
- `docker-compose up --build`
#### 2回目以降
- `docker-compose up`
#### DBの初期化をするときは
- `docker-compose down --volume`
#### DockerのDBにアクセスするときは
- `docker exec -it db-for-techread bash`
- `mysql -u ユーザー名 -p`

## API設計
|  エンドポイント  |ハンドラ|  メソッド  | 説明 |
| ---- | ---- |---- | ---- |
|  /login | LoginHandler |  POST  |ログイン機能|
|  /sign-up | SignUpHandler |  POST  | サインアップ機能 |
|  /mypage | FetchProfileHandler |  GET  | 自分のプロフィール情報取得 |
|  /edit-profile | EditProfileHandler |  POST  | 自分のプロフィール更新 |
|  /create-event | CreateEventHandler |  POST  |イベント作成機能 |
|  /edit-event | EditEventHandler |  POST  |イベント編集機能|
|  /event-list  | EventListHandler |  GET  |イベント一覧取得|
|  /chapter-list | ChapterHandler |  GET  |議事録一覧|
|  /create-chapter | CreateChapterHandler |  POST  |議事録作成|
|  /edit-chapter | EditChapterHandler |  POST  |議事録編集|

### Cookieについて
- user_id: ログインしてるユーザID
/login, /sign-upのレスポンスで来たuser_idをフロントエンド側でCookieに入れて、リクエストの度に送る。

## /login
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| email | string | 254 | ○ | eメールアドレス |
| password | string | 20 | ○ | ユーザーのパスワード |

### Response
#### 成功の場合
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存成功ステータス(success) |
| user_id | string | 10 | ○ | ユーザーID |

#### 失敗の場合
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存失敗ステータス(failed) |
| user_id | string | 10 | ○ | 保存失敗したときは0 |

## /sign-up
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| user_name | string | 40 | ○ | ユーザーの名前 |
| email | string | 254 | ○ | eメールアドレス |
| password | string | 20 | ○ | ユーザーのパスワード |

### Response(POST)
#### 成功の場合
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存成功ステータス(success) |
| user_id | string | 10 | ○ | ユーザーID |

#### 失敗の場合(同じEmailがDBにあるときなど)

| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存失敗ステータス(failed) |
| user_id | string | 10 | ○ | 保存失敗したときは0 |

## /create-event
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| event_name | string | 40 | ○ | イベントの名前 |
| book_name | string | 40 | ○ | 書籍名 |
| member_list | string | 40 | ○ | 参加メンバー |

### Response(POST)
#### 成功の場合
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存成功ステータス(success) |
| event_id | int | 10 | ○ | イベントID |

#### 失敗の場合(同じEmailがDBにあるときなど)

| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存失敗ステータス(failed) |
| event_id | int | 10 | ○ | イベントID(保存失敗したときは0) |

## /edit-event
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| event_name | string | 40 | ○ | イベントの名前 |
| book_name | string | 40 | ○ | 書籍名 |
| member_list | string | 40 | ○ | 参加メンバー |

### Response(POST)
#### 成功の場合
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存成功ステータス(success) |
| event_id | int | 10 | ○ | イベントID |

#### 失敗の場合(同じEmailがDBにあるときなど)

| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存失敗ステータス(failed) |
| event_id | int | 10 | ○ | イベントID(保存失敗したときは0) |

## /event-list
ログインしているユーザが参加しているイベントを返す

### Request(GET)
なし

### Response
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | ステータス |
| event_id | int | 10 | ○ | イベントID |
| book_name | string | 10 | ○ | 書籍名 |
| member_count | string | 10 | ○ | 参加人数 |
レスポンス例
```
{
  "res_state": "success",
  "events": [
    {
      "event_id": 1,
      "book_name": "book1",
      "member_count": 3,
    },
    {
      "event_id": 2,
      "book_name": "book2",
      "member_count": 5,
    },
  ]
}
```

## /create-chapter
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| event_id | int | 40 | ○ | イベントID |
| event_date | date | 40 | ○ | 開催日 |
| venue | string | 254 | ○ | 開催場所 |
| chapter_num | int | 20 | ○ | 章 |
| content | string | 20 | ○ | 内容 |

### Response(POST)
#### 成功の場合
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存成功ステータス(success) |
| event_id | int | 10 | ○ | イベントID |

#### 失敗の場合(同じEmailがDBにあるときなど)

| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存失敗ステータス(failed) |
| event_id | int | 10 | ○ | イベントID(保存失敗した時は0) |


## /chapter-list
対象のイベントの議事録一覧ページの情報を返す

### Request(GET)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| event_id | int | 10 | ○ | イベントID |

### Response
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | ステータス |
| event_date | date | 40 | ○ | 開催日 |
| venue | string | 254 | ○ | 開催場所 |
| chapter_num | int | 20 | ○ | 章 |
| content | int | 20 | ○ | 内容 |
レスポンス例
```
{
  "res_state": "success",
  "event_id": 1,
  "event_name": "リーダブルコード輪読会",
  "chapters": [
    {
      "chapter_id": 1,
      "event_date": "2022-11-02",
      "venue": "オンライン",
      "chapter_num": 1,
      "content": "綺麗に書くことの重要性が書かれた章。~~~~~~~aaaaaaaaa",
    },
    {
      "chapter_id": 2,
      "event_date": "2022-11-03",
      "venue": "スマプロのルーム",
      "chapter_num": 2,
      "content": "1行ごとの改善が大事ということを書いた章。~~~~~aaaaaaaa",
    },
  ]
}
```

## /edit-chapter
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| event_date | date | 40 | ○ | 開催日 |
| venue | string | 254 | ○ | 開催場所 |
| chapter_num | int | 20 | ○ | 章 |
| content | int | 20 | ○ | 内容 |

### Response(POST)
#### 成功の場合
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存成功ステータス(success) |
| event_id | int | 10 | ○ | イベントID |

#### 失敗の場合(同じEmailがDBにあるときなど)

| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| res_state | string | 10 | ○ | 保存失敗ステータス(failed) |
| event_id | int | 10 | ○ | イベントID(保存失敗した時は0) |
