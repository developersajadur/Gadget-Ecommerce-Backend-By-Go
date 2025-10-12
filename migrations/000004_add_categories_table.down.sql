ALTER TABLE products DROP CONSTRAINT IF EXISTS fk_products_category;
DROP TABLE IF EXISTS categories;
DROP EXTENSION IF EXISTS "uuid-ossp";
