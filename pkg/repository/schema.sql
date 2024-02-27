CREATE TABLE "Person" (
    "id" serial primary key,
    "name" varchar not null unique,
    "age" int check(age > 0)
);