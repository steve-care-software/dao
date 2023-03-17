package applications

import (
	application_blockchain "github.com/steve-care-software/fungible-unit-pow-blockchains/applications"
)

// Application represents the application
type Application interface {
	Blockchain() application_blockchain.Application
}
