CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY NOT NULL,
    "first_name" varchar(255) NOT NULL,
    "last_name" varchar(255) NOT NULL,
    "email" varchar(255) UNIQUE NOT NULL,
    "password" varchar(8) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS "post"(
    "id" UUID PRIMARY KEY NOT NULL,
    "title" varchar(255) NOT NULL,
    "body" TEXT NOT NULL,
    "user_id" UUID REFERENCES "users"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
);

INSERT INTO "post" ("id", "title", "body", "user_id") VALUES ('ab167112-bb09-45f3-ac54-07c21889ab39','Just','Just a body','720db164-6365-4288-8dbb-d98d1aca95cb');