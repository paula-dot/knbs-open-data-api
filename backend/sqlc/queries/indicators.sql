-- name: CreateIndicator :one
INSERT INTO indicators (code, name, description, unit, source)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListIndicators :many
SELECT * FROM indicators
ORDER BY name;

-- name: GetIndicatorByCode :one
SELECT * FROM indicators
WHERE code = $1
LIMIT 1;