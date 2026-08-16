-- name: CreateBusiness :one
INSERT INTO businesses (
  name, slug, timezone
) VALUES ( $1, $2, $3 )
RETURNING *;

-- name: GetBusinesses :many
SELECT * FROM businesses;

-- name: GetBusinessBySlug :one
SELECT *
  FROM businesses
  WHERE slug=$1;
