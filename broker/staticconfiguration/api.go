package staticconfiguration

import (
	"time"
)

type BaseConfig struct {
	ServicePort                       uint16        `yaml:"service_port"`
	StatusPort                        uint16        `yaml:"status_port"`
	TLSCertFilename                   string        `yaml:"tls_cert_filename"`
	TLSKeyFilename                    string        `yaml:"tls_key_filename"`
	AccountConfigurationUrl           string        `yaml:"account_configuration_url"`
	AccountConfigurationCheckInterval time.Duration `yaml:"account_configuration_check_interval"`
}

type StaticConfiguration struct {
	Base BaseConfig
}
