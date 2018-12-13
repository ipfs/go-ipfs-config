package config

type CID struct {
	Base       string // default multibase to use in output for API and command line
	OutputVer1 bool   // upgrades CIDv0 to CIDv1 in output for most commands
	                  // option only used when Base is not the empty string

	//Version int // default Cid for newly created CIDs for command that
	              // support the --cid-version or similar command
	//

	//DefaultVersion     int // default CID version for all newly created CID
	//DefaultBase string int // default multibase for all CID output
}
