CREATE TABLE IF NOT EXISTS magazines (
	id TEXT NOT NULL UNIQUE,
	inserted INTEGER DEFAULT (unixepoch()) NOT NULL,
	number INT UNIQUE NOT NULL,
	date INTEGER NOT NULL,
	location TEXT NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id TEXT NOT NULL UNIQUE,
    pwd TEXT NOT NULL,
    created INTEGER DEFAULT (unixepoch()) NOT NULL,
    lastonline INTEGER DEFAULT (unixepoch()) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS unames (
    uid TEXT NOT NULL,
    uname TEXT NOT NULL UNIQUE,
    FOREIGN KEY (uid) REFERENCES users(id),
    PRIMARY KEY (uid, uname)
);
