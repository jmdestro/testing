// ===================================================================================
// AssetChain v0.5 -  Chaincode definition for test
// ===================================================================================

package main

// ----- Tickets ----- //
type Test struct {
	ObjectType         string  `json:"docType"`     //field for couchdb
	Status             string  `json:"status"`
}
