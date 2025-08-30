package database

const v1 = `
	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS user (
		id TEXT NOT NULL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		created TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

		CHECK(
        	typeof("username") = "text" AND
        	length("username") <= 14
    	)
	);

	CREATE TABLE IF NOT EXISTS blurt (
		id TEXT NOT NULL PRIMARY KEY,
		content TEXT NOT NULL,
		author_id TEXT NOT NULL,
		created TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

		CHECK(
        	typeof("content") = "text" AND
        	length("content") <= 14
    	),

		FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS lik (
		user_id TEXT NOT NULL,
		blurt_id TEXT NOT NULL,
		PRIMARY KEY (user_id, blurt_id),
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (blurt_id) REFERENCES blurt(id) ON DELETE CASCADE
	);
`
