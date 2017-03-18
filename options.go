package googlecloud

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Unknwon/com"
	"github.com/rai-project/config"
	"github.com/rai-project/utils"
)

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

func decrypt(s string) string {
	if utils.IsEncryptedString(s) {
		c, err := utils.DecryptStringBase64(config.App.Secret, s)
		if err == nil {
			return c
		}
	}
	return s
}

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

func FromDefaultConfigurationFile() Option {
	return func(o *Options) {
		filename, err := DefaultConfigurationSource()
		if err != nil {
			return
		}
		FromConfigurationFile(filename)(o)
	}
}

func FromConfigurationFile(filename string) Option {
	return func(o *Options) {
		if !com.IsFile(filename) {
			return
		}
		buf, err := ioutil.ReadFile(filename)
		if err != nil {
			return
		}

		var d Options
		if err := json.Unmarshal(buf, &d); err != nil {
			return
		}
		*o = d
	}
}

func Type(s string) Option {
	return func(o *Options) {
		o.Type = s
	}
}

func ProjectID(s string) Option {
	return func(o *Options) {
		o.ProjectID = decrypt(s)
	}
}

func PrivateKeyID(s string) Option {
	return func(o *Options) {
		o.PrivateKeyID = decrypt(s)
	}
}

func PrivateKey(s string) Option {
	return func(o *Options) {
		o.PrivateKey = decrypt(s)
	}
}

func ClientEmail(s string) Option {
	return func(o *Options) {
		o.ClientEmail = decrypt(s)
	}
}

func ClientID(s string) Option {
	return func(o *Options) {
		o.ClientID = decrypt(s)
	}
}

func AuthURI(s string) Option {
	return func(o *Options) {
		o.AuthURI = decrypt(s)
	}
}

func TokenURI(s string) Option {
	return func(o *Options) {
		o.TokenURI = decrypt(s)
	}
}

func AuthProviderX509CertURL(s string) Option {
	return func(o *Options) {
		o.AuthProviderX509CertURL = decrypt(s)
	}
}

func ClientX509CertURL(s string) Option {
	return func(o *Options) {
		o.ClientX509CertURL = decrypt(s)
	}
}

func (o *Options) Bytes() []byte {
	buf, err := json.Marshal(o)
	if err != nil {
		return []byte{}
	}
	return buf
}
