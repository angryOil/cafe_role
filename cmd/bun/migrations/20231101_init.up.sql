SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."cafe_role"
(
    id          SERIAL PRIMARY KEY,
    cafe_id     int         not null,
    name        VARCHAR(50) NOT NULL,
    description VARCHAR(2000),
    created_at  timestamptz
);


create unique index cr_cafe_id_name_unique on cafe_role (cafe_id, name);
