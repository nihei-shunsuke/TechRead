CREATE TABLE IF NOT EXISTS users (
  user_id       INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,
  user_name     VARCHAR(40)     NOT NULL,
  email         VARCHAR(254)     NOT NULL,
  password      VARCHAR(20)     NOT NULL,
  created_at    DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at    DATETIME NOT NULL DEFAULT (DATETIME('now')),
);

CREATE TABLE IF NOT EXISTS comments (
  comment_id    INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,
  chapter_id    INTEGER  NOT NULL,
  comment       VARCHAR(255)     NOT NULL DEFAULT '',
  user_id       INTEGER  NOT NULL,
  created_at    DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at    DATETIME NOT NULL DEFAULT (DATETIME('now')),

  FOREIGN KEY (chapter_id) REFERENCES chapters(chapter_id) ON DELETE CASCADE --外部キー制約
  FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE --外部キー制約
);

CREATE TABLE IF NOT EXISTS chapters (
  chapter_id      INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,
  chapter_num     INTEGER  NOT NULL,
  event_id        INTEGER  NOT NULL,
  event_date      DATETIME NOT NULL DEFAULT '',
  content         TEXT     NOT NULL DEFAULT '',
  last_updater_id INTEGER  NOT NULL,
  created_at      DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at      DATETIME NOT NULL DEFAULT (DATETIME('now')),

  FOREIGN KEY (event_id) REFERENCES chapters(event_id) ON DELETE CASCADE --外部キー制約
  FOREIGN KEY (last_updater_id) REFERENCES users(user_id) ON DELETE CASCADE --外部キー制約
);

CREATE TABLE IF NOT EXISTS events (
  event_id        INTEGER  NOT NULL PRIMARY KEY AUTOINCREMENT,
  event_name      VARCHAR(100) NOT NULL,
  start_date      DATETIME,
  end_date        DATETIME,
  organizer_id    INTEGER  NOT NULL,
  created_at      DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at      DATETIME NOT NULL DEFAULT (DATETIME('now')),

  CREATE INDEX ON events(start_date); --インデックスを貼る
  CREATE INDEX ON events(end_date); --インデックスを貼る
  FOREIGN KEY (organizer_id) REFERENCES users(user_id) --外部キー制約
);

CREATE TABLE IF NOT EXISTS event_groups (
  user_id         INTEGER  NOT,
  event_id        INTEGER  NOT,
  created_at      DATETIME NOT NULL DEFAULT (DATETIME('now')),
  updated_at      DATETIME NOT NULL DEFAULT (DATETIME('now')),

  FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE --外部キー制約
  FOREIGN KEY (event_id) REFERENCES events(event_id) ON DELETE CASCADE --外部キー制約
  PRIMARY KEY(user_id, event_id) --複合主キー
);
