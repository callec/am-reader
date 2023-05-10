CREATE TABLE IF NOT EXISTS magazines (
	id TEXT DEFAULT (uuid()) NOT NULL,
	inserted INTEGER DEFAULT (unixepoch()) NOT NULL,
	number INT UNIQUE NOT NULL,
	date INTEGER NOT NULL,
	location TEXT NOT NULL,
	PRIMARY KEY (id)
);
