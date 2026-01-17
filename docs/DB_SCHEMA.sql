-- Royal Vault database schema

CREATE TABLE users (
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE vaults (
    id                  SERIAL PRIMARY KEY,
    user_id             INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    encrypted_vault_key TEXT NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE folders (
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name        VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE credentials (
    id                  SERIAL PRIMARY KEY,
    vault_id            INTEGER NOT NULL REFERENCES vaults(id) ON DELETE CASCADE,
    folder_id           INTEGER REFERENCES folders(id) ON DELETE SET NULL,
    title               VARCHAR(255) NOT NULL,
    url                 TEXT,
    encrypted_username  TEXT NOT NULL,
    encrypted_password  TEXT NOT NULL,
    encrypted_notes     TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE security_questions (
    id                  SERIAL PRIMARY KEY,
    credential_id       INTEGER NOT NULL REFERENCES credentials(id) ON DELETE CASCADE,
    encrypted_question  TEXT NOT NULL,
    encrypted_answer    TEXT NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);