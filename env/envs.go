package env

import "github.com/Drafteame/mgorepo/logger"

const (
	MongoURIEnv             = "MONGO_URI"
	MongoReadPreferenceEnv  = "MONGO_READ_PREFERENCE"
	MongoAuthSourceEnv      = "MONGO_AUTH_SOURCE"
	MongoAuthMechanismEnv   = "MONGO_AUTH_MECHANISM"
	MongoReplicaSetEnv      = "MONGO_REPLICA_SET"
	MongoRetryWritesEnv     = "MONGO_RETRY_WRITES"
	MongoSSLVerifyEnv       = "MONGO_SSL_VERIFY_CERTIFICATE"
	MongoMinPoolSizeEnv     = "MONGO_MIN_POOL_SIZE"
	MongoMaxPoolSizeEnv     = "MONGO_MAX_POOL_SIZE"
	MongoConnectTimeoutEnv  = "MONGO_CONNECT_TIMEOUT"
	MongoQueryTimeoutEnv    = "MONGO_QUERY_TIMEOUT"
	MongoUsernameEnv        = "MONGO_USERNAME"
	MongoPasswordEnv        = "MONGO_PASSWORD"
	MongoClusterEndpointEnv = "MONGO_CLUSTER_ENDPOINT"
	MongoCertPathEnv        = "MONGO_CERT_PATH"
	MongoDBNameEnv          = "MONGO_DB_NAME"
	MongoLogLevelEnv        = "MONGO_LOG_LEVEL"

	MongoURIDefault             = ""
	MongoReadPreferenceDefault  = "primary"
	MongoAuthSourceDefault      = ""
	MongoAuthMechanismDefault   = ""
	MongoReplicaSetDefault      = ""
	MongoRetryWritesDefault     = "false"
	MongoSSLVerifyDefault       = false
	MongoMinPoolSizeDefault     = 1
	MongoMaxPoolSizeDefault     = 10
	MongoConnectTimeoutDefault  = 5
	MongoQueryTimeoutDefault    = 30
	MongoUsernameDefault        = ""
	MongoPasswordDefault        = ""
	MongoClusterEndpointDefault = ""
	MongoCertPathDefault        = ""
	MongoDBNameDefault          = ""
	MongoLogLevelDefault        = logger.LevelNone
)

var defaultEnvs = map[string]any{
	MongoURIEnv:             MongoURIDefault,
	MongoReadPreferenceEnv:  MongoReadPreferenceDefault,
	MongoAuthSourceEnv:      MongoAuthSourceDefault,
	MongoAuthMechanismEnv:   MongoAuthMechanismDefault,
	MongoReplicaSetEnv:      MongoReplicaSetDefault,
	MongoRetryWritesEnv:     MongoRetryWritesDefault,
	MongoSSLVerifyEnv:       MongoSSLVerifyDefault,
	MongoMinPoolSizeEnv:     MongoMinPoolSizeDefault,
	MongoMaxPoolSizeEnv:     MongoMaxPoolSizeDefault,
	MongoConnectTimeoutEnv:  MongoConnectTimeoutDefault,
	MongoQueryTimeoutEnv:    MongoQueryTimeoutDefault,
	MongoUsernameEnv:        MongoUsernameDefault,
	MongoPasswordEnv:        MongoPasswordDefault,
	MongoClusterEndpointEnv: MongoClusterEndpointDefault,
	MongoCertPathEnv:        MongoCertPathDefault,
	MongoDBNameEnv:          MongoDBNameDefault,
	MongoLogLevelEnv:        MongoLogLevelDefault,
}
