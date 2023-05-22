-- name: GetMagazine :one
SELECT * FROM magazines
WHERE id = ? LIMIT 1;

-- name: GetMagazineByNumber :one
SELECT * FROM magazines
WHERE number = ? LIMIT 1;

-- name: ListMagazines :many
SELECT * FROM magazines
ORDER BY date
LIMIT ? OFFSET ?;

-- name: AddMagazine :execresult
INSERT INTO magazines (
	number,
	date,
	location
) VALUES (
	?, ?, ?
);

-- name: RemoveMagazine :exec
DELETE FROM magazines
WHERE id = ?;

-- name: GetUid :one
SELECT uid FROM unames
WHERE uname = ? LIMIT 1;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: RegisterUser :one
INSERT INTO users (
    pwd
) VALUES (
    ?
) RETURNING id;

-- name: AddUName :execresult
INSERT INTO unames (
    uid,
    uname
) VALUES (
    ?, ?
);
