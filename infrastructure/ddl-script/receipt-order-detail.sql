
CREATE TABLE receipt_order_detail (
	id INTEGER NOT NULL PRIMARY KEY,
	receipt_order_id integer Not NULL,
	product_id integer not null,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
