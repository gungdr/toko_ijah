
CREATE TABLE sales_detail (
	id INTEGER NOT NULL PRIMARY KEY,
	sales_id integer not null,
	stock_id integer not null,
	price real not null,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
