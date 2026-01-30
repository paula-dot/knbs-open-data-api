-- name: CreateCounty :one
INSERT INTO counties (id, name, code, former_province, area_sq_km)
VALUES ( $1, $2, $3, $4, $5)
RETURNING *;

-- name: GetCounty :one
SELECT * FROM counties
WHERE id = $1
LIMIT 1;

-- name: ListCounties :many
SELECT * FROM counties
ORDER BY id;
