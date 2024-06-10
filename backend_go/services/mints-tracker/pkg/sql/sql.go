package sql

const UpsertMintQuery = `
	INSERT INTO mints (
		tx_signature,
		program_key,
		block_time,
		collection_key,
		asset_key,
		asset_owner_key,
		asset_name,
		asset_symbol,
		asset_image,
		asset_description,
		asset_attributes,
		mint_price
	)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	ON CONFLICT (asset_key) DO NOTHING;
`

const UpsertCollectionQuery = `
	INSERT INTO collections (
		collection_key,
		candy_machine,
		name,
		symbol,
		image,
		size,
		type,
		mint_price,
		total_minted,
		launchpad,
		updated_at
	)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())
	ON CONFLICT (collection_key) DO UPDATE SET
		size = EXCLUDED.size,
		mint_price = EXCLUDED.mint_price,
		candy_machine = EXCLUDED.candy_machine,
		updated_at = NOW()
	;
`

const SelectCollectionForUpdate = `
	UPDATE collections SET
		updated_at = NOW()
	WHERE collections.collection_key IN (
		SELECT collections.collection_key FROM collections
			JOIN (
				SELECT collection_key, MAX(created_at) AS created_at FROM mints
				GROUP BY collection_key
			) latest_mints ON collections.collection_key = latest_mints.collection_key
		WHERE
			(collections.size != collections.total_minted OR collections.size = 0) AND
			NOW() - interval '10 seconds' > collections.updated_at AND
			latest_mints.created_at > NOW() - interval '24 hours'
		ORDER BY collections.updated_at DESC
		LIMIT 1
	)
	RETURNING
		collections.collection_key,
		collections.candy_machine,
		collections.launchpad,
		collections.type,
		collections.name
	;
`

const UpdateCollection = `
	UPDATE collections SET
		total_minted = $2,
		size = $3
	WHERE collection_key = $1;
`

const UpsertCollectionLink = `
	INSERT INTO collection_links (collection_key, type, uri)
	VALUES ($1, $2, $3)
	ON CONFLICT (collection_key, type) DO UPDATE SET
		uri = EXCLUDED.uri,
		updated_at = NOW();
`
