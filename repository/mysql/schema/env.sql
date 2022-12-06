CREATE TABLE IF NOT EXISTS envs (
	env_id VARCHAR(512) CHARACTER SET ascii,
	destination VARCHAR(512) CHARACTER SET ascii,
	PRIMARY KEY (env_id, destination)
);
