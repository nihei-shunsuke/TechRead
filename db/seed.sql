INSERT INTO users (
  user_name, email, password, created_at, updated_at
) VALUES (
  '木下真菜', 'hoge@hoge.com', 'password', now(), now()
);

INSERT INTO events (
  event_name, start_date, organizer_id, created_at, updated_at
) VALUES (
  'リーダブルコード輪読会', now(), 1, now(), now()
);

INSERT INTO chapters (
  chapter_num, event_id, event_date, content, last_updater_id, created_at, updated_at
) VALUES (
  1, 1, now(), '理解しやすいコード', 1, now(), now()
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