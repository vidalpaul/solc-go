package data

// lang represents the language of the source code.
type lang string

const (
	langSolidity lang = "Solidity"
	langYul      lang = "Yul"
)

// EVMVersion represents the EVM version to compile for.
type EVMVersion string

const (
	EVMVersionShanghai   EVMVersion = "shanghai"
	EVMVersionParis      EVMVersion = "paris"
	EVMVersionLondon     EVMVersion = "london"
	EVMVersionBerlin     EVMVersion = "berlin"
	EVMVersionIstanbul   EVMVersion = "istanbul"
	EVMVersionPetersburg EVMVersion = "petersburg"
	EVMVersionByzantium  EVMVersion = "byzantium"
)
