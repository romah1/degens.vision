package sql

const SelectTopCollections = `
	SELECT
		top_collections.collection_key,
		top_collections.mints_count_period,
		collections_with_count.mints_count_total,
		collections_with_count.name,
		collections_with_count.symbol,
		collections_with_count.image,
		collections_with_count.size,
		collections_with_count.launchpad,
		collections_with_count.mint_price
	FROM
		(
			SELECT mints.collection_key, count(*) AS mints_count_period FROM mints
			WHERE mints.created_at > NOW() - $1 * interval '1 minutes'
			GROUP BY mints.collection_key
		) top_collections
			JOIN
		(
			SELECT
				collections.collection_key,
				collections.name,
				collections.symbol,
				collections.image,
				collections.size,
				collections.launchpad,
				collections.mint_price,
				GREATEST(count(*), collections.total_minted) AS mints_count_total FROM
					collections JOIN mints ON collections.collection_key = mints.collection_key
			GROUP BY (collections.collection_key)
		) collections_with_count
			ON top_collections.collection_key = collections_with_count.collection_key
			ORDER BY top_collections.mints_count_period DESC;
`

const SelectTopHolders = `
	SELECT
		mints.asset_owner_key,
		count(*) AS assets_amount
	FROM
		mints
	WHERE mints.collection_key = $1
	GROUP BY mints.asset_owner_key
	ORDER BY count(*) DESC;
`

const SelectCollection = `
	SELECT
		collections.collection_key,
		collections.candy_machine,
		collections.name,
		collections.symbol,
		collections.image,
		collections.size,
		collections.mint_price,
		GREATEST(collection_with_count.mints_count, collections.total_minted),
		collections.launchpad,
		collection_with_count.attributes_amount,
		collection_with_unique_minters.unique_minters,
		collections.created_at,
		collections.type
	FROM
		(
			SELECT
				collection_key,
				max(COALESCE(array_length(asset_attributes, 1), 0)) AS attributes_amount,
				count(*) AS mints_count
				FROM mints
			WHERE collection_key = $1
			GROUP BY collection_key
		) collection_with_count
		JOIN collections ON collection_with_count.collection_key=collections.collection_key
		JOIN (
			SELECT
				tmp.collection_key,
				count(*) AS unique_minters
			FROM (SELECT DISTINCT collection_key, asset_owner_key FROM mints WHERE collection_key = $1) tmp
			GROUP BY tmp.collection_key
		) collection_with_unique_minters ON collections.collection_key=collection_with_unique_minters.collection_key
	WHERE collections.collection_key = $1;
`

const SelectCollectionLinks = `
		SELECT type, uri FROM collection_links
		WHERE collection_key = $1;
`

const SelectMintsDistribution = `
	SELECT owner_counts.mints_amount, count(*) AS holders_amount FROM
	(
			SELECT
					asset_owner_key,
					count(*) AS mints_amount
			FROM mints
			WHERE collection_key = $1
			GROUP BY asset_owner_key
	) owner_counts
	GROUP BY owner_counts.mints_amount;
`

const UpsertAuthInit = `
		INSERT INTO auth_init (address, nonce)
		VALUES ($1, $2)
		ON CONFLICT (address) DO UPDATE SET
			nonce = EXCLUDED.nonce;
`

const SelectNonceFromAuthInit = `
		SELECT nonce FROM auth_init
		WHERE address = $1;
`

const DeleteAuthInit = `
		DELETE FROM auth_init
		WHERE address = $1;
`

const InitUser = `
		INSERT INTO users(address)
		VALUES ($1)
		ON CONFLICT (address) DO UPDATE SET
			address = EXCLUDED.address
		RETURNING users.id;
`

const UpsertUser = `
		INSERT INTO users (address, name, image, updated_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT DO UPDATE SET
			address = EXCLUDED.address,
			name = EXCLUDED.name
			image = EXCLUDED.image
			updated_at = NOW();
`

const SelectUser = `
		SELECT address, name, image FROM users
		WHERE address = '$1';
`

const SearchCollections = `
		SELECT collection_key, name, symbol, image FROM collections
		WHERE LOWER(name) LIKE $1 OR collection_key LIKE $1
		LIMIT 10;
`

const SelectTopCollectionsHoldersMinted = `
    SELECT
			top_collections.collection_key,
			top_collections.mints_count,
			collections.name,
			collections.symbol,
			collections.image,
			collections.size
		FROM
		(
			SELECT collection_key, count(*) AS mints_count FROM
			(
				SELECT asset_owner_key FROM mints
				WHERE collection_key = $1
			) holders JOIN mints ON mints.asset_owner_key = holders.asset_owner_key
			GROUP BY collection_key
		) top_collections JOIN collections ON top_collections.collection_key = collections.collection_key
		ORDER BY top_collections.mints_count DESC;
`

const SelectInitialMintsFeed = `
	  SELECT
			mints.tx_signature,
			mints.program_key,
			mints.block_time,
			mints.collection_key,
			mints.asset_key,
			mints.asset_owner_key,
			mints.asset_name,
			mints.asset_symbol,
			mints.asset_image,
			mints.mint_price,
			mints.asset_description,
			mints.asset_attributes,
			collections.collection_key,
			collections.name,
			collections.symbol,
			collections.image,
			collections.size,
			collections.launchpad
		FROM mints JOIN collections ON mints.collection_key = collections.collection_key
		ORDER BY mints.created_at DESC
		LIMIT 20;
`

const SelectHasRefcode = `
  SELECT EXISTS(SELECT FROM user_refcode WHERE user_id = $1);
`

const ClaimRefcode = `
  INSERT INTO user_refcode (user_id, refcode_id) VALUES ($1, $2);
`

const SelectRefcodeId = `
	SELECT id FROM refcodes where code = $1;
`

const AuthWithRefcode = `
SET TRANSACTION REPEATABLE READ;

WITH _user_id AS (
	INSERT INTO users (address) VALUES ($1)
	ON CONFLICT (address) DO NOTHING
	RETURNING users.id
)
WITH _refcode_id AS (
	SELECT refcodes.id FROM refcodes
	WHERE refcodes.code = $2
)
INSERT INTO user_refcode (user_id, refcode_id)
VALUES (_user_id, _refcode_id);

UPDATE refcodes SET
	cur_usages = cur_usages + 1
WHERE refcodes.code = $2;
`

const InsertRefcode = `
	INSERT INTO refcodes (code, max_usages) VALUES ($1, $2);
`

const SelectRefcodes = `
	SELECT
		refcodes.code,
		refcodes.max_usages,
		COALESCE(refcode_count.cur_usages, 0)
	FROM refcodes LEFT JOIN
		(
			SELECT refcode_id, count(*) AS cur_usages FROM user_refcode
			GROUP BY refcode_id
		) refcode_count
		ON refcodes.id = refcode_count.refcode_id;
`
