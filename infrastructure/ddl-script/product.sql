
CREATE TABLE product (
	id INTEGER NOT NULL PRIMARY KEY,
	brand_id integer NOT NULL,
	color_id integer NOT NULL,
	size_id integer NOT NULL,
	name text (50) not null,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
