-- +goose Up
-- +goose StatementBegin
CREATE TABLE "public"."names" (
	"nconst"              varchar(255),
	"primary_name"        varchar(255),
	"birth_year"          varchar(4),
	"death_year"          varchar(4) DEFAULT '',
	"primary_professions" varchar[],
	"known_for_titles"    varchar[]
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE "public"."names";

-- +goose StatementEnd
