
CREATE TABLE sales (
	id INTEGER NOT NULL PRIMARY KEY,
	date timestamp not null,
	notes text not null,
	is_void integer NOT NULL DEFAULT 0,
	created_on TIMESTAMP NOT NULL,
	modified_on TIMESTAMP
);
