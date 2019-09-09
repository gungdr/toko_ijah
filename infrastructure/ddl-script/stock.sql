
CREATE TABLE "stock" (
	id INTEGER NOT NULL PRIMARY KEY,
	receipt_order_id integer not null,
	product_id integer not null,
	price real not null,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
