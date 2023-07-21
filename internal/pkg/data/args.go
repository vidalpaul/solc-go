package data

import "encoding/json"

type input struct {
	Lang     lang           `json:"language"`
	Sources  map[string]src `json:"sources"`
	Settings *Settings      `json:"settings"`
}

type src struct {
	Keccak256 string   `json:"keccak256,omitempty"`
	Content   string   `json:"content,omitempty"`
	URLS      []string `json:"urls,omitempty"`
}

// Settings for the compilation.
type Settings struct {
	lang            lang                           `json:"-"`
	Remappings      []string                       `json:"remappings,omitempty"`
	Optimizer       *Optimizer                     `json:"optimizer"`
	ViaIR           bool                           `json:"viaIR,omitempty"`
	EVMVersion      EVMVersion                     `json:"evmVersion"`
	OutputSelection map[string]map[string][]string `json:"outputSelection"`
}

type Optimizer struct {
	Enabled bool   `json:"enabled"`
	Runs    uint64 `json:"runs"`
}

type output struct {
	Errors    []error_                       `json:"errors"`
	Sources   map[string]srcOut              `json:"sources"`
	Contracts map[string]map[string]contract `json:"contracts"`
}

type sourceLocation struct {
	File  string `json:"file"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type srcOut struct {
	ID        int             `json:"id"`
	AST       json.RawMessage `json:"ast"`
	LegacyAST json.RawMessage `json:"legacyAST"`
}

type contract struct {
	ABI      []json.RawMessage `json:"abi"`
	Metadata string            `json:"metadata"`
	UserDoc  json.RawMessage   `json:"userdoc"`
	DevDoc   json.RawMessage   `json:"devdoc"`
	IR       string            `json:"ir"`
	EVM      evm               `json:"evm"`
}

type evm struct {
	Assembly          string                       `json:"assembly"`
	LegacyAssembly    json.RawMessage              `json:"legacyAssembly"`
	Bytecode          bytecode                     `json:"bytecode"`
	DeployedBytecode  bytecode                     `json:"deployedBytecode"`
	MethodIdentifiers map[string]string            `json:"methodIdentifiers"`
	GasEstimates      map[string]map[string]string `json:"gasEstimates"`
}

type bytecode struct {
	Object         hexBytes                              `json:"object"`
	Opcodes        string                                `json:"opcodes"`
	SourceMap      string                                `json:"sourceMap"`
	LinkReferences map[string]map[string][]linkReference `json:"linkReferences"`
}

type linkReference struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
