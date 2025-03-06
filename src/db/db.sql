CREATE TABLE okaimono.Animes (
	id SMALLINT UNSIGNED auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	tags varchar(50) NULL,
	in_live BOOL NOT NULL,
	next_new_cap TINYINT UNSIGNED DEFAULT 0 NOT NULL,
	max_caps SMALLINT UNSIGNED DEFAULT 12 NOT NULL,
	last_view_cap SMALLINT UNSIGNED DEFAULT 0 NOT NULL,
	prequels varchar(100) NULL,
	sequels varchar(100) NULL,
	movies varchar(100) NULL,
	spin_offs varchar(100) NULL,
	ovas TINYINT UNSIGNED DEFAULT 0 NOT NULL,
	img_path varchar(100) NULL,
	CONSTRAINT Animes_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
