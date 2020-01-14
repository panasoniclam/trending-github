-- +migrate Up
CREATE TABLE "scc" (
    id integer primary key,
    username text unique ,
    email text unique ,
    password text
);

-- +migrate Down
DROP TABLE "scc"