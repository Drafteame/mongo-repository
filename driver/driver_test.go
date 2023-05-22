package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildConnectionURI(t *testing.T) {
	tests := []struct {
		name string
		conf Config
		exp  string
		err  error
	}{
		{
			name: "ReadPreference",
			conf: Config{
				UserName:        "a",
				Password:        "b",
				ClusterEndpoint: "c",
				DBName:          "d",
				ReadPreference:  "e",
			},
			exp: "mongodb://a:b@c/d?readPreference=e",
		},
		{
			name: "AuthSource",
			conf: Config{
				UserName:        "a",
				Password:        "b",
				ClusterEndpoint: "c",
				DBName:          "d",
				AuthSource:      "e",
			},
			exp: "mongodb://a:b@c/d?authSource=e",
		},
		{
			name: "AuthMechanism",
			conf: Config{
				UserName:        "a",
				Password:        "b",
				ClusterEndpoint: "c",
				DBName:          "d",
				AuthMechanism:   "e",
			},
			exp: "mongodb://a:b@c/d?authMechanism=e",
		},
		{
			name: "ReplicaSet",
			conf: Config{
				UserName:        "a",
				Password:        "b",
				ClusterEndpoint: "c",
				DBName:          "d",
				ReplicaSet:      "e",
			},
			exp: "mongodb://a:b@c/d?replicaSet=e",
		},
		{
			name: "CertPath",
			conf: Config{
				UserName:        "a",
				Password:        "b",
				ClusterEndpoint: "c",
				DBName:          "d",
				CertPath:        "e",
			},
			exp: "mongodb://a:b@c/d?tls=true&sslVerifyCertificate=false",
		},
		{
			name: "CertPath verify certificate",
			conf: Config{
				UserName:             "a",
				Password:             "b",
				ClusterEndpoint:      "c",
				DBName:               "d",
				CertPath:             "e",
				SSLVerifyCertificate: true,
			},
			exp: "mongodb://a:b@c/d?tls=true&sslVerifyCertificate=true",
		},
		{
			name: "AuthSource and AuthMechanism and ReplicaSet",
			conf: Config{
				UserName:        "a",
				Password:        "b",
				ClusterEndpoint: "c",
				DBName:          "d",
				AuthSource:      "e",
				AuthMechanism:   "f",
				ReplicaSet:      "g",
			},
			exp: "mongodb://a:b@c/d?authSource=e&authMechanism=f&replicaSet=g",
		},
		{
			name: "AuthSource and AuthMechanism, ReplicaSet and RetryWrites",
			conf: Config{
				UserName:        "a",
				Password:        "b",
				ClusterEndpoint: "c",
				DBName:          "d",
				AuthSource:      "e",
				AuthMechanism:   "f",
				ReplicaSet:      "g",
				RetryWrites:     "true",
			},
			exp: "mongodb://a:b@c/d?authSource=e&authMechanism=f&replicaSet=g&retryWrites=true",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := buildConnectionURI(&test.conf)

			if test.err != nil {
				assert.Equal(t, test.err.Error(), err.Error())
			}

			assert.Equal(t, test.exp, got)
		})
	}
}
