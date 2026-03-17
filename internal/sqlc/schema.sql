-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "first_name" text,
    "last_name" text,
    "email" text NOT NULL,
    "phone" text,
    "password" text,
    "code" text,
    "expiry" timestamptz,
    "verified" bool DEFAULT false,
    "user_type" text DEFAULT 'buyer'::text,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS addresses_id_seq;

-- Table Definition
CREATE TABLE "public"."addresses" (
    "id" int8 NOT NULL DEFAULT nextval('addresses_id_seq'::regclass),
    "address_line1" text,
    "address_line2" text,
    "city" text,
    "post_code" int8,
    "country" text,
    "user_id" int8,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "fk_users_address" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id"),
    PRIMARY KEY ("id")
);