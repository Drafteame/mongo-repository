package driver

import (
	"github.com/Drafteame/mgorepo/internal/env"
)

// Config is the configuration for the driver.
type Config struct {
	SSLVerifyCertificate bool
	ConnectTimeout       int
	QueryTimeout         int
	MinPoolSize          int
	MaxPoolSize          int
	URI                  string
	ReadPreference       string
	RetryWrites          string
	AuthSource           string
	AuthMechanism        string
	ReplicaSet           string
	Username             string
	Password             string
	ClusterEndpoint      string
	CertPath             string
	DBName               string
}

func DefaultConfig() Config {
	return Config{
		SSLVerifyCertificate: env.GetBool(MongoSSLVerifyEnv),
		ConnectTimeout:       env.GetInt(MongoConnectTimeoutEnv),
		QueryTimeout:         env.GetInt(MongoQueryTimeoutEnv),
		URI:                  env.GetString(MongoURIEnv),
		ReadPreference:       env.GetString(MongoReadPreferenceEnv),
		RetryWrites:          env.GetString(MongoRetryWritesEnv),
		AuthSource:           env.GetString(MongoAuthSourceEnv),
		AuthMechanism:        env.GetString(MongoAuthMechanismEnv),
		ReplicaSet:           env.GetString(MongoReplicaSetEnv),
		Username:             env.GetString(MongoUsernameEnv),
		Password:             env.GetString(MongoPasswordEnv),
		ClusterEndpoint:      env.GetString(MongoClusterEndpointEnv),
		CertPath:             env.GetString(MongoCertPathEnv),
		DBName:               env.GetString(MongoDBNameEnv),
		MinPoolSize:          env.GetInt(MongoMinPoolSizeEnv),
		MaxPoolSize:          env.GetInt(MongoMaxPoolSizeEnv),
	}
}
