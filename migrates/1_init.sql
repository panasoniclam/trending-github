-- +migrate Up
CREATE TABLE "users" (
  "user_id" text PRIMARY KEY,
  "full_name" text,
  "email" text UNIQUE,
  "password" text,
  "role" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "repos" (
  "name" text PRIMARY KEY,
  "description" text,
  "url" text,
  "color" text,
  "lang" text ,
  "fork" text,
  "stars" text,
  "stars_today" text,
  "build_by" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "bookmarks" (
  "bid" text PRIMARY KEY,
  "user_id" text,
  "repo_name" text UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "khachhang" (
    "makh" text PRIMARY KEY,
    "tenkh" text NOT NULL ,
    "diachi" text NOT NULL,
    "dienthoai" text UNIQUE
);

CREATE TABLE "phieugiaohang" (
    "magiao" text PRIMARY KEY,
    "ngaygiao" TIMESTAMPTZ NOT NULL,
    "madat" text UNIQUE NOT NULL
);
CREATE TABLE "chitietgiaohang" (
    "magiao" text NOT NULL UNIQUE ,
    "mahh" text NOT NULL UNIQUE ,
    "slgiao" integer NOT NULL,
    "dongiagiao" integer NOT NULL
);


CREATE TABLE "hanghoa" (
    "mahh" text PRIMARY KEY,
    "tenhh" text NOT NULL ,
    "dvt" text,
    "slcon" integer,
    "dongiahh" float
);

CREATE TABLE "dondathang" (
    "madat" text PRIMARY KEY,
    "ngaydat" TIMESTAMPTZ NOT NULL ,
    "makh" text NOT NULL ,
    "tinhtrang" text
);
CREATE TABLE "chitietdathang" (
    "madat" text PRIMARY KEY,
    "makh" text NOT NULL,
    "sldat" integer
);
--
CREATE TABLE "lichsugia" (
    "mahh" text NOT NULL,
    "ngayhl" TIMESTAMPTZ NOT NULL,
    "dongia" float
);
--
ALTER TABLE "phieugiaohang" ADD FOREIGN KEY ("madat") REFERENCES "chitietdathang" ("madat");
ALTER TABLE "chitietgiaohang" ADD FOREIGN KEY ("magiao") REFERENCES "phieugiaohang" ("magiao");
ALTER TABLE "chitietgiaohang" ADD FOREIGN KEY ("mahh") REFERENCES "hanghoa" ("mahh");
ALTER TABLE "dondathang" ADD FOREIGN KEY ("makh") REFERENCES "khachhang" ("makh");
ALTER TABLE "lichsugia" ADD FOREIGN KEY ("mahh")  REFERENCES "hanghoa" ("mahh");
ALTER TABLE "chitietdathang" ADD FOREIGN KEY ("makh") REFERENCES "khachhang" ("makh");



ALTER TABLE "bookmarks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
-- p0EE6bmdpH5cpLeu
--https://cloud.google.com/vpc/docs/using-firewalls#creating_firewall_rules
ALTER TABLE "bookmarks" ADD FOREIGN KEY ("repo_name") REFERENCES "repos" ("name");

-- +migrate Down
DROP TABLE bookmarks;
DROP TABLE users;
DROP TABLE repos;


DROP TABLE lichsugia;
DROP TABLE chitietgiaohang ;
DROP TABLE dondathang;
DROP TABLE khachhang;
DROP TABLE chitietdathang;
DROP TABLE phieugiaohang;
DROP TABLE hanghoa;

ALTER TABLE "phieugiaohang" DROP CONSTRAINT "madat";