package credentials

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineerr"
)

var (
	// ErrNoValidProvidersFoundInChain Is returned when there are no valid
	// providers in the ChainProvider.
	//
	// This has been deprecated. For verbose error messaging set
	// volcengine.Config.CredentialsChainVerboseErrors to true.
	ErrNoValidProvidersFoundInChain = volcengineerr.New("NoCredentialProviders",
		`no valid providers in chain. Deprecated.
	For verbose messaging see volcengine.Config.CredentialsChainVerboseErrors`,
		nil)
)

// A ChainProvider will search for a provider which returns credentials
// and cache that provider until Retrieve is called again.
//
// The ChainProvider provides a way of chaining multiple providers together
// which will pick the first available using priority order of the Providers
// in the list.
//
// If none of the Providers retrieve valid credentials Value, ChainProvider's
// Retrieve() will return the error ErrNoValidProvidersFoundInChain.
//
// If a Provider is found which returns valid credentials Value ChainProvider
// will cache that Provider for all calls to IsExpired(), until Retrieve is
// called again.
//
// Example of ChainProvider to be used with an EnvProvider and EC2RoleProvider.
// In this example EnvProvider will first check if any credentials are available
// via the environment variables. If there are none ChainProvider will check
// the next Provider in the list, EC2RoleProvider in this case. If EC2RoleProvider
// does not return any credentials ChainProvider will return the error
type ChainProvider struct {
	Providers     []Provider
	curr          Provider
	VerboseErrors bool
}

// NewChainCredentials returns a pointer to a new Credentials object
// wrapping a chain of providers.
func NewChainCredentials(providers []Provider) *Credentials {
	return NewCredentials(&ChainProvider{
		Providers: append([]Provider{}, providers...),
	})
}

// Retrieve returns the credentials value or error if no provider returned
// without error.
//
// If a provider is found it will be cached and any calls to IsExpired()
// will return the expired state of the cached provider.
func (c *ChainProvider) Retrieve() (Value, error) {
	var errs []error
	for _, p := range c.Providers {
		creds, err := p.Retrieve()
		if err == nil {
			c.curr = p
			return creds, nil
		}
		errs = append(errs, err)
	}
	c.curr = nil

	var err error
	err = ErrNoValidProvidersFoundInChain
	if c.VerboseErrors {
		err = volcengineerr.NewBatchError("NoCredentialProviders", "no valid providers in chain", errs)
	}
	return Value{}, err
}

// IsExpired will returned the expired state of the currently cached provider
// if there is one.  If there is no current provider, true will be returned.
func (c *ChainProvider) IsExpired() bool {
	if c.curr != nil {
		return c.curr.IsExpired()
	}

	return true
}
