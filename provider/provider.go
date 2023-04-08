package provider

// Provider is the interface that wraps the basic Register and Boot methods.
type Provider interface {
	// Register is called when the application is registering the provider.
	Register()

	// Boot is called when the application is booting the provider.
	Boot()
}

// UnimplementedProvider is a default implementation of the Provider interface.
// It is used to avoid implementing all methods of the Provider interface.
//
// Example:
// package provider
//
//	type MyProvider struct {
//		*provider.UnimplementedProvider
//	}
//
//	var _ provider.Provider = (*MyProvider)(nil)
//
//	func (p *MyProvider) Register() {
//		// do something
//	}
type UnimplementedProvider struct {
}

// Register is a default implementation of the Provider interface.
func (u *UnimplementedProvider) Register() {
}

// Boot is a default implementation of the Provider interface.
func (u *UnimplementedProvider) Boot() {
}
