package driver

import "github.com/Drafteame/mgorepo/env"

// Config is the configuration for the driver.
type Config struct {
	SSLVerifyCertificate bool
	ConnectTimeout       int
	QueryTimeout         int
	MinPoolSize          *int
	MaxPoolSize          *int
	URI                  string
	ReadPreference       string
	RetryWrites          string
	AuthSource           string
	AuthMechanism        string
	ReplicaSet           string
	UserName             string
	Password             string
	ClusterEndpoint      string
	CertPath             string
	DBName               string
}

func DefaultConfig() Config {
	var maxPoolSize, minPoolSize *int

	if val := env.GetInt(env.MongoMaxPoolSizeEnv); val != 0 {
		maxPoolSize = &val
	}

	if val := env.GetInt(env.MongoMinPoolSizeEnv); val != 0 {
		minPoolSize = &val
	}

	return Config{
		SSLVerifyCertificate: env.GetBool(env.MongoSSLVerifyEnv),
		ConnectTimeout:       env.GetInt(env.MongoConnectTimeoutEnv),
		QueryTimeout:         env.GetInt(env.MongoQueryTimeoutEnv),
		URI:                  env.GetString(env.MongoURIEnv),
		ReadPreference:       env.GetString(env.MongoReadPreferenceEnv),
		RetryWrites:          env.GetString(env.MongoRetryWritesEnv),
		AuthSource:           env.GetString(env.MongoAuthSourceEnv),
		AuthMechanism:        env.GetString(env.MongoAuthMechanismEnv),
		ReplicaSet:           env.GetString(env.MongoReplicaSetEnv),
		UserName:             env.GetString(env.MongoUsernameEnv),
		Password:             env.GetString(env.MongoPasswordEnv),
		ClusterEndpoint:      env.GetString(env.MongoClusterEndpointEnv),
		CertPath:             env.GetString(env.MongoCertPathEnv),
		DBName:               env.GetString(env.MongoDBNameEnv),
		MinPoolSize:          minPoolSize,
		MaxPoolSize:          maxPoolSize,
	}
}
