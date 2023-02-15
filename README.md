# TechRead

## 使用ライブラリ
- gorilla: Routing, Middleware,Session
- gorm: ORM

## 起動方法


|  エンドポイント  |ハンドラ|  メソッド  | 説明 |
| ---- | ---- |---- | ---- |
|  /login | LoginHandler |  POST  |ログイン機能|
|  /sign-up | SignUpHandler |  POST  | サインアップ機能 |
|  /mypage | FetchProfileHandler |  GET  | 自分のプロフィール情報取得 |
|  /edit-profile | EditProfileHandler |  POST  | 自分のプロフィール更新 |
|  /create-event | CreateEventHandler |  POST  |イベント作成機能 |
|  /edit-event | EditEventHandler |  POST  |イベント編集機能|
|  /event-list  | EventListHandler |  GET  |イベント一覧取得|
|  /event-detail | EventDetailHandler |  GET  |イベント詳細|
|  /create-abstract | CreateAbstractHandler |  POST  |議事録作成|
|  /edit-abstract | EditAbstractHandler |  POST  |議事録編集|


## LoginHandler
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| user_id | string | 10 | ○ | ユーザーの独自ID |
| password | string | 20 | ○ | ユーザーのパスワード |

### Response(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| user_id | string | 10 | ○ | ユーザーの独自ID |
| login_flg | bool |  | ○ | ログイン成功or失敗 |

## SignUpHandler
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| user_id | string | 10 | ○ | ユーザーの独自ID |
| user_name | string | 40 | ○ | ユーザーの名前 |
| password | string | 20 | ○ | ユーザーのパスワード |
| email | string | 254 | ○ | eメールアドレス |

### Response(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| login_flg | bool |  | ○ | 新規登録成功or失敗 |

## CreateEventHandler
### Request(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| user_id | string | 10 | ○ | ユーザーの独自ID |
| user_name | string | 40 | ○ | ユーザーの名前 |
| password | string | 20 | ○ | ユーザーのパスワード |
| email | string | 254 | ○ | eメールアドレス |

### Response(POST)
| JSON Key | 型 | サイズ | 必須 | 値の説明 |
|:-----------|:-----------|:-----------|:-----------|:-----------|
| login_flg | bool |  | ○ | 新規登録成功or失敗 |