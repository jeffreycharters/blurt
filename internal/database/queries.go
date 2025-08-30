package database

const v1 = `
	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS user (
		username TEXT NOT NULL PRIMARY KEY,
		created TEXT NOT NULL,

		CHECK(
        	typeof("username") = "text" AND
        	length("username") <= 14
    	)
	);

	CREATE TABLE IF NOT EXISTS blurt (
		id TEXT NOT NULL PRIMARY KEY,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		created TEXT NOT NULL,

		CHECK(
        	typeof("content") = "text" AND
        	length("content") <= 14
    	),

		FOREIGN KEY (author) REFERENCES user(username) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS lik (
		username TEXT NOT NULL,
		blurt_id TEXT NOT NULL,
		PRIMARY KEY (username, blurt_id),
		FOREIGN KEY (username) REFERENCES user(username) ON DELETE CASCADE,
		FOREIGN KEY (blurt_id) REFERENCES blurt(id) ON DELETE CASCADE
	);
`
