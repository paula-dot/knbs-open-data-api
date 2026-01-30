-- name: CreateObservation :one
INSERT INTO observations (county_id, indicator_id, year, value, source_document)
VALUES ( $1, $2, $3, $4, $5)
RETURNING *;

-- name: GetDataByCountyAndYear :many
SELECT
    o.value,
    o.year,
    c.name as county_name,
    i.name as indicator_name,
    i.unit
FROM observations o
JOIN counties c ON c.id = o.county_id
JOIN indicators i ON i.id = o.indicator_id
WHERE c.id = $1 AND o.year = $2;