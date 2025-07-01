CREATE TYPE "article_status" AS ENUM (
    'draft',
    'published'
);

CREATE TYPE "user_role" AS ENUM (
    'admin',
    'editor',
    'reporter',
    'reader'
);

CREATE TABLE "articles" (
    "id" uuid PRIMARY KEY,
    "title" varchar(256) NOT NULL,
    "body" text NOT NULL,
    "status" article_status DEFAULT 'draft',
    "category_id" uuid,
    "author_id" uuid,
    "publish_at" timestamp,
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "users" (
    "id" uuid PRIMARY KEY,
    "name" varchar(32),
    "email" varchar(320) UNIQUE NOT NULL,
    "role" user_role DEFAULT 'reader',
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "categories" (
    "id" uuid PRIMARY KEY,
    "name" varchar(100) UNIQUE NOT NULL,
    "slug" varchar(100) UNIQUE NOT NULL,
    "description" text,
    "parent_id" uuid,
    "is_active" bool DEFAULT 'true',
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "article_views" (
    "id" uuid PRIMARY KEY,
    "article_id" uuid,
    "viewed_at" timestamp,
    "user_id" uuid
);

CREATE TABLE "comments" (
    "id" uuid PRIMARY KEY,
    "article_id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "parent_id" uuid,
    "body" text NOT NULL,
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "tags" (
    "id" uuid PRIMARY KEY,
    "name" varchar(50) UNIQUE NOT NULL,
    "slug" varchar(50) UNIQUE NOT NULL,
    "created_at" timestamp DEFAULT now(),
    "updated_at" timestamp DEFAULT now(),
    "deleted_at" timestamp
);

CREATE TABLE "article_tags" (
    "article_id" uuid NOT NULL,
    "tag_id" uuid NOT NULL
);

CREATE INDEX ON "articles" ("category_id");

CREATE INDEX ON "articles" ("author_id");

CREATE INDEX ON "articles" ("status");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "categories" ("slug");

CREATE INDEX ON "comments" ("article_id");

CREATE INDEX ON "comments" ("user_id");

CREATE INDEX ON "article_tags" ("article_id", "tag_id");

ALTER TABLE "articles" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "articles" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "article_views" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");

ALTER TABLE "article_views" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("parent_id") REFERENCES "comments" ("id");

ALTER TABLE "article_tags" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");

ALTER TABLE "article_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

ALTER TABLE "categories" ADD FOREIGN KEY ("parent_id") REFERENCES "categories" ("id") ON DELETE SET NULL;
