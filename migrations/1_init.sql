-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE "scc" (
    "id" integer primary key,
    "username" text unique ,
    "email" text unique ,
    "password" text
);
-- +migrate StatementEnd
-- +migrate Down
DROP TABLE "scc"