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
