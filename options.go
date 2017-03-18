package googlecloud

type Options struct {
	Type                    string `json:"type,omitempty"`
	ProjectID               string `json:"project_id,omitempty"`
	PrivateKeyID            string `json:"private_key_id,omitempty"`
	PrivateKey              string `json:"private_key,omitempty"`
	ClientEmail             string `json:"client_email,omitempty"`
	ClientID                string `json:"client_id,omitempty"`
	AuthURI                 string `json:"auth_uri,omitempty"`
	TokenURI                string `json:"token_uri,omitempty"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`
	ClientX509CertURL       string `json:"client_x509_cert_url,omitempty"`
}

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	o := &Options{
		Type:                    Config.Type,
		ProjectID:               Config.ProjectID,
		PrivateKeyID:            Config.PrivateKeyID,
		PrivateKey:              Config.PrivateKey,
		ClientEmail:             Config.ClientEmail,
		ClientID:                Config.ClientID,
		AuthURI:                 Config.AuthURI,
		TokenURI:                Config.TokenURI,
		AuthProviderX509CertURL: Config.AuthProviderX509CertURL,
		ClientX509CertURL:       Config.ClientX509CertURL,
	}

	for _, f := range opts {
		f(o)
	}

	return o
}

func Type(s string) Option {
	return func(o *Options) {
		o.Type = s
	}
}

func ProjectID(s string) Option {
	return func(o *Options) {
		o.ProjectID = s
	}
}

func PrivateKeyID(s string) Option {
	return func(o *Options) {
		o.PrivateKeyID = s
	}
}

func PrivateKey(s string) Option {
	return func(o *Options) {
		o.PrivateKey = s
	}
}

func ClientEmail(s string) Option {
	return func(o *Options) {
		o.ClientEmail = s
	}
}

func ClientID(s string) Option {
	return func(o *Options) {
		o.ClientID = s
	}
}

func AuthURI(s string) Option {
	return func(o *Options) {
		o.AuthURI = s
	}
}

func TokenURI(s string) Option {
	return func(o *Options) {
		o.TokenURI = s
	}
}

func AuthProviderX509CertURL(s string) Option {
	return func(o *Options) {
		o.AuthProviderX509CertURL = s
	}
}

func ClientX509CertURL(s string) Option {
	return func(o *Options) {
		o.ClientX509CertURL = s
	}
}
