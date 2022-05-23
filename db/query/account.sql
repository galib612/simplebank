/* this comment instruct how tp generate the Golang function signature for this query.
    :one specify the no of object returns.
 */

-- name: CreateAccount :one 
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;