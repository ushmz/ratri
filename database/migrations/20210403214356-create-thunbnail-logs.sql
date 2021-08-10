
-- +migrate Up
CREATE TABLE IF NOT EXISTS behavior_thumb_logs  (
    id VARCHAR(200) NOT NULL,
    uid VARCHAR(200) NOT NULL,
    url VARCHAR(512) NOT NULL,
    time_on_thumb INT(11) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS behavior_thumb_logs;