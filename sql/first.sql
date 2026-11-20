SET search_path TO marketplace;

SELECT id, email, name 
FROM users 
ORDER BY id ASC;