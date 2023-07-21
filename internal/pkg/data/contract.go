package data

// Contract represents a compiled contract.
type Contract struct {
	Code       []byte // The bytecode of the contract after deployment.
	DeployCode []byte // The bytecode to deploy the contract.
}
