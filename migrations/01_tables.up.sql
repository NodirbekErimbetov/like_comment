CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY NOT NULL,
    "name" varchar(255) NOT NULL,
    "user_name" varchar(255) UNIQUE NOT NULL,
    "bio" TEXT,
    "media_url" varchar(255),
    "password" varchar(8) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS "posts"(
    "id" UUID PRIMARY KEY NOT NULL,
    "user_id" UUID REFERENCES "users"("id"),
    "title" varchar(255) NOT NULL,
    "body" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP 
);
CREATE TABLE IF NOT EXISTS "likes"(
    "post_id" UUID REFERENCES "posts"("id"),
    "count" bigint
);

CREATE TABLE IF NOT EXISTS "signup" (
    "name" varchar(255) NOT NULL,
    "user_name" varchar(255) NOT NULL,
    "bio" TEXT,
    "media_url" varchar(255),
    "password" varchar(8) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "login" (
    "user_name" varchar(255) NOT NULL,
    "password" varchar(8) NOT NULL
);

CREATE TABLE IF NOT EXISTS "follow" (
    "follower_id" UUID PRIMARY KEY NOT NULL,
    "followed_id" UUID PRIMARY KEY NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);