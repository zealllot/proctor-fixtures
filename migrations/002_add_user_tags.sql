CREATE TABLE user_tags (
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  tag TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (user_id, tag)
);

CREATE INDEX idx_user_tags_tag ON user_tags(tag);
