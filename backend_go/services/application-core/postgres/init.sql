BEGIN;

CREATE TABLE IF NOT EXISTS mints (
  asset_key TEXT NOT NULL PRIMARY KEY,
  tx_signature TEXT NOT NULL,
  program_key TEXT NOT NULL,
  block_time TIMESTAMPTZ NOT NULL,
  collection_key TEXT NOT NULL,
  asset_owner_key TEXT NOT NULL,
  asset_name TEXT NOT NULL,
  asset_symbol TEXT NOT NULL,
  asset_image TEXT NOT NULL,
  asset_description TEXT NOT NULL,
  asset_attributes JSONB[] NOT NULL,
  mint_price NUMERIC NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX mints_collection_key_idx ON mints(collection_key);
CREATE INDEX mints_block_time_idx ON mints(block_time);
CREATE INDEX mint_created_at_idx ON mints(created_at);

CREATE TYPE LAUNCHPAD AS ENUM ('lmnft', 'truffle');

CREATE TYPE COLLECTION_TYPE AS ENUM ('cnft', 'mplx', 'core');

CREATE TABLE IF NOT EXISTS collections (
  collection_key TEXT NOT NULL PRIMARY KEY,
  candy_machine TEXT,
  name TEXT NOT NULL,
  symbol TEXT NOT NULL,
  image TEXT NOT NULL,
  size INTEGER NOT NULL,
  type COLLECTION_TYPE NOT NULL,
  total_minted INTEGER NOT NULL,
  mint_price NUMERIC NOT NULL,
  launchpad LAUNCHPAD NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TYPE COLLECTION_LINK_TYPE AS ENUM ('mint', 'discord', 'twitter');

CREATE TABLE IF NOT EXISTS collection_links (
  collection_key TEXT NOT NULL,
  type COLLECTION_LINK_TYPE NOT NULL,
  uri TEXT NOT NULL,
  PRIMARY KEY (collection_key, type),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX collections_updated_at_idx ON collections(updated_at);
CREATE INDEX collections_created_at_idx ON collections(created_at);

CREATE TABLE IF NOT EXISTS auth_init (
  address TEXT NOT NULL PRIMARY KEY,
  nonce TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX auth_init_created_at_idx ON auth_init(created_at);

CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  address TEXT NOT NULL UNIQUE,
  name TEXT,
  image TEXT,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS refcodes (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  code TEXT NOT NULL UNIQUE,
  max_usages INT NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_refcode (
  user_id BIGINT NOT NULL REFERENCES users(id) PRIMARY KEY,
  refcode_id BIGINT NOT NULL REFERENCES refcodes(id)
);

CREATE OR REPLACE FUNCTION verify_user_refcode()
RETURNS TRIGGER AS
$$
  DECLARE
    currentUsages BIGINT;
    maxUsages INT;
  BEGIN
    SELECT max_usages INTO maxUsages
    FROM refcodes WHERE id = NEW.refcode_id;

    SELECT count(*) INTO currentUsages
    FROM user_refcode WHERE user_refcode.refcode_id = NEW.refcode_id;

    IF currentUsages > maxUsages THEN
      RAISE EXCEPTION 'usages exceeded';
    END IF;

    RETURN NEW;
  END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER claim_ref_code_trigger
AFTER UPDATE OR INSERT ON user_refcode
FOR EACH ROW EXECUTE PROCEDURE verify_user_refcode();

COMMIT;
