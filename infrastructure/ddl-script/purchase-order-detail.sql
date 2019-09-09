
CREATE TABLE purchase_order_detail (
	id INTEGER NOT NULL PRIMARY KEY,
	purchase_order_id integer not null,
	product_id integer NOT NULL,
	price real not null,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
