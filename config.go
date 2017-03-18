package googlecloud

import (
	"github.com/apex/log"
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type googleCloudConfig struct {
	Type                    string        `json:"type" config:"googlecloud.type"`
	ProjectID               string        `json:"project_id" config:"googlecloud.project_id"`
	PrivateKeyID            string        `json:"private_key_id" config:"googlecloud.private_key_id"`
	PrivateKey              string        `json:"private_key" config:"googlecloud.private_key"`
	ClientEmail             string        `json:"client_email" config:"googlecloud.client_email"`
	ClientID                string        `json:"client_id" config:"googlecloud.client_id"`
	AuthURI                 string        `json:"auth_uri" config:"googlecloud.auth_uri"`
	TokenURI                string        `json:"token_uri" config:"googlecloud.token_uri"`
	AuthProviderX509CertURL string        `json:"auth_provider_x509_cert_url" config:"googlecloud.auth_provider_x509_cert_url"`
	ClientX509CertURL       string        `json:"client_x509_cert_url" config:"googlecloud.client_x509_cert_url"`
	done                    chan struct{} `json:"-" config:"-"`
}

var (
	Config = &googleCloudConfig{
		done: make(chan struct{}),
	}
)

func (googleCloudConfig) ConfigName() string {
	return "GoogleCloud"
}

func (a *googleCloudConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *googleCloudConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
}

func (c googleCloudConfig) Wait() {
	<-c.done
}

func (c googleCloudConfig) String() string {
	return pp.Sprintln(c)
}

func (c googleCloudConfig) Debug() {
	log.Debug("GoogleCloud Config = ", c)
}

func init() {
	config.Register(Config)
}
