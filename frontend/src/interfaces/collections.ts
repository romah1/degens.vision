export interface TopMitsResponse {
    collections: CollectionInfo[];
}

export interface CollectionInfoResponse {
    collection: CollectionFullInfo;
    mints_distribution: {
        mints_amount: number;
        holders_amount: number;
    }[];
}

interface LinkInfo {
    type: "mint" | "twitter" | "magiceden" | "tensor" | "sniper" | "hyperspace";
    uri: string;
}

interface CollectionFullInfo {
    collection_key: string;
    image: string;
    launchpad: string;
    mint_price: number;
    mints_count: number;
    name: string;
    size: number;
    symbol: string;
    date: string;
    attributes_amount: number;
    unique_minters_amount: number;
    type: string;
    links: LinkInfo[];
    market_links: LinkInfo[];
}

export interface CollectionInfo {
    collection_key: string;
    image: string;
    mints_count_period: number;
    name: string;
    symbol: string;
    size: number;
    mints_count_total?: number;
    mint_price: number;
    launchpad: string;
}

export interface LiveMintsItem {
    collection: {
        collection_key: string;
        image: string;
        name: string;
        size: number;
        symbol: string;
        launchpad: string;
    };
    mint: {
        asset_image: string;
        asset_description: string;
        asset_attributes: {
            value: string;
            trait_type: string;
        }[];
        asset_key: string;
        asset_name: string;
        asset_owner_key: string;
        asset_symbol: string;
        block_time: string;
        collection_key: string;
        program_key: string;
        tx_signature: string;
        mint_price: number;
    }
}
