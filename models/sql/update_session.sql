UPDATE sessions
SET token_hash = $2
WHERE user_id = $1
RETURNING id;