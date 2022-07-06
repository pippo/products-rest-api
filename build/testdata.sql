USE products_rest_api;

REPLACE INTO products(id, sku, pname, category, price) VALUES(1, "000001", "BV Lean leather ankle boots", "boots", 89000);
REPLACE INTO products(id, sku, pname, category, price) VALUES(2, "000002", "Naima embellished suede sandals", "sandals", 79500);
REPLACE INTO products(id, sku, pname, category, price) VALUES(3, "000003", "Nathane leather sneakers", "snickers", 59000);
REPLACE INTO products(id, sku, pname, category, price) VALUES(4, "000004", "Newer BV Lean leather ankle boots", "boots", 99000);
REPLACE INTO products(id, sku, pname, category, price) VALUES(5, "000005", "Newer Naima embellished suede sandals", "sandals", 89500);
REPLACE INTO products(id, sku, pname, category, price) VALUES(6, "000006", "Nathane leather sneakers enchanced", "snickers", 85000);

-- 

REPLACE INTO discounts(id, sku, category, percent) VALUES(1, "000001", NULL, 10);
REPLACE INTO discounts(id, sku, category, percent) VALUES(2, "000002", NULL, 3);
REPLACE INTO discounts(id, sku, category, percent) VALUES(3, NULL, "sandals", 5);