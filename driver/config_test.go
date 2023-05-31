package driver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	t.Run("get default config", func(t *testing.T) {
		conf := DefaultConfig()

		assert.Equal(t, defaultEnvs[MongoURIEnv], conf.URI)
		assert.Equal(t, defaultEnvs[MongoReadPreferenceEnv], conf.ReadPreference)
		assert.Equal(t, defaultEnvs[MongoRetryWritesEnv], conf.RetryWrites)
		assert.Equal(t, defaultEnvs[MongoAuthSourceEnv], conf.AuthSource)
		assert.Equal(t, defaultEnvs[MongoAuthMechanismEnv], conf.AuthMechanism)
		assert.Equal(t, defaultEnvs[MongoReplicaSetEnv], conf.ReplicaSet)
		assert.Equal(t, defaultEnvs[MongoUsernameEnv], conf.Username)
		assert.Equal(t, defaultEnvs[MongoPasswordEnv], conf.Password)
		assert.Equal(t, defaultEnvs[MongoClusterEndpointEnv], conf.ClusterEndpoint)
		assert.Equal(t, defaultEnvs[MongoCertPathEnv], conf.CertPath)
		assert.Equal(t, defaultEnvs[MongoDBNameEnv], conf.DBName)
		assert.Equal(t, defaultEnvs[MongoMaxPoolSizeEnv], conf.MaxPoolSize)
		assert.Equal(t, defaultEnvs[MongoMinPoolSizeEnv], conf.MinPoolSize)
		assert.Equal(t, defaultEnvs[MongoSSLVerifyEnv], conf.SSLVerifyCertificate)
		assert.Equal(t, defaultEnvs[MongoConnectTimeoutEnv], conf.ConnectTimeout)
		assert.Equal(t, defaultEnvs[MongoQueryTimeoutEnv], conf.QueryTimeout)
	})

	t.Run("get default config with envs", func(t *testing.T) {
		t.Setenv(MongoURIEnv, "mongodb://localhost:27017")
		t.Setenv(MongoReadPreferenceEnv, "primary")
		t.Setenv(MongoRetryWritesEnv, "true")
		t.Setenv(MongoAuthSourceEnv, "admin")
		t.Setenv(MongoAuthMechanismEnv, "SCRAM-SHA-1")
		t.Setenv(MongoReplicaSetEnv, "rs0")
		t.Setenv(MongoUsernameEnv, "root")
		t.Setenv(MongoPasswordEnv, "root")
		t.Setenv(MongoClusterEndpointEnv, "localhost:27017")
		t.Setenv(MongoCertPathEnv, "certs/mongodb.pem")
		t.Setenv(MongoDBNameEnv, "test")
		t.Setenv(MongoMaxPoolSizeEnv, "100")
		t.Setenv(MongoMinPoolSizeEnv, "10")
		t.Setenv(MongoSSLVerifyEnv, "true")
		t.Setenv(MongoConnectTimeoutEnv, "1000")
		t.Setenv(MongoQueryTimeoutEnv, "1000")

		conf := DefaultConfig()

		assert.Equal(t, "mongodb://localhost:27017", conf.URI)
		assert.Equal(t, "primary", conf.ReadPreference)
		assert.Equal(t, "true", conf.RetryWrites)
		assert.Equal(t, "admin", conf.AuthSource)
		assert.Equal(t, "SCRAM-SHA-1", conf.AuthMechanism)
		assert.Equal(t, "rs0", conf.ReplicaSet)
		assert.Equal(t, "root", conf.Username)
		assert.Equal(t, "root", conf.Password)
		assert.Equal(t, "localhost:27017", conf.ClusterEndpoint)
		assert.Equal(t, "certs/mongodb.pem", conf.CertPath)
		assert.Equal(t, "test", conf.DBName)
		assert.Equal(t, 100, conf.MaxPoolSize)
		assert.Equal(t, 10, conf.MinPoolSize)
		assert.Equal(t, true, conf.SSLVerifyCertificate)
		assert.Equal(t, 1000, conf.ConnectTimeout)
		assert.Equal(t, 1000, conf.QueryTimeout)
	})
}
