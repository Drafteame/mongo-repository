package driver

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func buildConnection(ctx context.Context, uri, certPath string, minSize, maxSize *int) (*mongo.Client, error) {
	tlsConfig, err := getCustomTLSConfig(certPath)
	if err != nil {
		return nil, err
	}

	opt := options.Client().ApplyURI(uri)

	if minSize != nil {
		opt = opt.SetMinPoolSize(uint64(*minSize))
	}

	if maxSize != nil {
		opt = opt.SetMaxPoolSize(uint64(*maxSize))
	}

	if tlsConfig != nil {
		opt = opt.SetTLSConfig(tlsConfig)
	}

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

// buildConnectionURI builds a connection URI from a configuration
func buildConnectionURI(config *Config) (string, error) {
	if config.URI != "" {
		dbName, err := getDBNameFromURI(config.URI)
		if err != nil {
			return "", err
		}

		config.DBName = dbName

		return config.URI, nil
	}

	if !validateMinimumConfig(config) {
		return "", errors.New("mongo-driver: missing required configuration")
	}

	template := "mongodb://%s:%s@%s/%s"
	flag := false
	token := "?"

	type params struct {
		check    string
		property []string
		value    string
	}

	paramsList := []params{
		{
			check:    config.ReadPreference,
			property: []string{config.ReadPreference},
			value:    "%s%sreadPreference=%s",
		},
		{
			check:    config.AuthSource,
			property: []string{config.AuthSource},
			value:    "%s%sauthSource=%s",
		},
		{
			check:    config.AuthMechanism,
			property: []string{config.AuthMechanism},
			value:    "%s%sauthMechanism=%s",
		},
		{
			check:    config.ReplicaSet,
			property: []string{config.ReplicaSet},
			value:    "%s%sreplicaSet=%s",
		},
		{
			check:    config.RetryWrites,
			property: []string{config.RetryWrites},
			value:    "%s%sretryWrites=%s",
		},
		{
			check:    config.CertPath,
			property: []string{},
			value:    "%s%s" + fmt.Sprintf("tls=true&sslVerifyCertificate=%t", config.SSLVerifyCertificate),
		},
	}

	for _, param := range paramsList {
		if param.check == "" {
			continue
		}

		replaces := []any{template, token}

		for _, p := range param.property {
			replaces = append(replaces, p)
		}

		template = fmt.Sprintf(param.value, replaces...)

		if !flag {
			token = "&"
			flag = true
		}
	}

	return fmt.Sprintf(template, config.UserName, config.Password, config.ClusterEndpoint, config.DBName), nil
}

func validateMinimumConfig(config *Config) bool {
	if config.UserName == "" || config.Password == "" || config.ClusterEndpoint == "" || config.DBName == "" {
		return false
	}

	return true
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	if caFile == "" {
		return nil, nil
	}

	tlsConfig := new(tls.Config)
	certs, err := os.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("mongo-driver: failed parsing pem file")
	}

	return tlsConfig, nil
}

func getDBNameFromURI(uri string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", errors.Join(errors.New("mongo-driver: error getting db name from uri"), err)
	}

	if parsed.Path == "" {
		return "", errors.New("mongo-driver: no db name specified in uri")
	}

	return parsed.Path[1:], nil
}
