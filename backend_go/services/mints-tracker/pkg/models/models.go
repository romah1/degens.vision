package models

type Launchpad = string

const (
	LaunchpadLmnft   Launchpad = "lmnft"
	LaunchpadTruffle Launchpad = "truffle"
)

type CollectionType = string

const (
	CollectionTypeCNFT CollectionType = "cnft"
	CollectionTypeMPLX CollectionType = "mplx"
	CollectionTypeCore CollectionType = "core"
)

type CollectionLinkType = string

const (
	CollectionLinkTypeMint    string = "mint"
	CollectionLinkTypeDiscord string = "discord"
	CollectionLinkTypeTwitter string = "twitter"
)
