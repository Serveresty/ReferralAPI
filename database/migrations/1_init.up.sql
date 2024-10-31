CREATE TABLE IF NOT EXISTS "users_data" (
    "user_id" SERIAL PRIMARY KEY,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "referrer_id" INT REFERENCES "users_data"("user_id") ON DELETE SET NULL
);

-- CREATE TABLE IF NOT EXISTS "users_ref_tokens" (
--     "user_id" INT PRIMARY KEY REFERENCES "users_data"("user_id") ON DELETE CASCADE,
--     "ref_token" VARCHAR(255) UNIQUE NOT NULL,
--     "expired_at" TIMESTAMP NOT NULL
-- );

CREATE TABLE IF NOT EXISTS "ref_tokens" (
    "ref_token" VARCHAR(255) PRIMARY KEY,
    "user_id" INT UNIQUE NOT NULL REFERENCES "users_data"("user_id") ON DELETE CASCADE,
    "expired_at" TIMESTAMP NOT NULL
);