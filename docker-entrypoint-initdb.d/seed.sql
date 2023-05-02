<<<<<<< Updated upstream
INSERT INTO users (
  user_name, email, password, created_at, updated_at
) VALUES (
  '木下真菜', 'hoge@hoge.com', 'password', now(), now()
);

INSERT INTO events (
  event_name, book_name, start_date, organizer_id, created_at, updated_at
) VALUES (
  'リーダブルコード輪読会', 'リーダブルコード ―より良いコードを書くためのシンプルで実践的なテクニック',now(), 1, now(), now()
);

INSERT INTO chapters (
  chapter_num, venue, event_id, event_date, content, last_updater_id, created_at, updated_at
) VALUES (
  1, 'オンライン', 1, now(), '理解しやすいコード', 1, now(), now()
);

INSERT INTO comments (
  chapter_id, comment, user_id, created_at, updated_at
) VALUES (
  1, '読みやすいコード大事！', 1, now(), now()
);

INSERT INTO event_groups (
  user_id, event_id, created_at, updated_at
) VALUES (
  1, 1, now(), now()
);
=======
CREATE DATABASE techread;
USE techread;

INSERT INTO users (
  user_name, email, password, created_at, updated_at
) VALUES (
  '木下真菜', 'hoge@hoge.com', 'password', now(), now()
);

INSERT INTO events (
  event_name, book_name, start_date, organizer_id, created_at, updated_at
) VALUES (
  'リーダブルコード輪読会', 'リーダブルコード ―より良いコードを書くためのシンプルで実践的なテクニック',now(), 1, now(), now()
);

INSERT INTO chapters (
  chapter_num, venue, event_id, event_date, content, last_updater_id, created_at, updated_at
) VALUES (
  1, 'オンライン', 1, date(now()), '理解しやすいコード', 1, now(), now()
);

INSERT INTO comments (
  chapter_id, comment, user_id, created_at, updated_at
) VALUES (
  1, '読みやすいコード大事！', 1, now(), now()
);

INSERT INTO event_groups (
  user_id, event_id, created_at, updated_at
) VALUES (
  1, 1, now(), now()
);
>>>>>>> Stashed changes
