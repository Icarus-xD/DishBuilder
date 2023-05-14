-- name: GetByType :many
SELECT ingredient.title as title, price, ingredient_type.title as type, code 
FROM ingredient
JOIN ingredient_type ON ingredient.type_id=ingredient_type.id
WHERE code = ?;