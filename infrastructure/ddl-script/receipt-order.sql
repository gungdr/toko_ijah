
CREATE TABLE receipt_order (
	id INTEGER NOT NULL PRIMARY KEY,
	purchase_order_id integer not null,
	receipt_number text null,
	date timestamp not null,
	notes text not null,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
