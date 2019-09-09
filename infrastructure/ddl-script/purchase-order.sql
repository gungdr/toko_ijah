
CREATE TABLE purchase_order (
	id INTEGER NOT NULL PRIMARY KEY,
	date timestampt not null,
	notes text,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
