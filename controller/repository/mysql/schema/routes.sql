CREATE TABLE IF NOT EXISTS routes (
	proxy_id VARCHAR(512) CHARACTER SET ascii,
	env_id VARCHAR(512) CHARACTER SET ascii,
	destination VARCHAR(512) CHARACTER SET ascii NOT NULL,
	PRIMARY KEY (proxy_id, env_id)
)
