package driver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Driver struct {
	client *mongo.Client
	dbName string
}

func New(ctx context.Context) (*Driver, error) {
	return NewWithConfig(ctx, DefaultConfig())
}

func NewWithConfig(ctx context.Context, config Config) (*Driver, error) {
	uri, err := buildConnectionURI(&config)
	if err != nil {
		return nil, err
	}

	client, errConn := buildConnection(
		ctx,
		uri,
		config.CertPath,
		config.MinPoolSize,
		config.MaxPoolSize,
	)

	if errConn != nil {
		return nil, errConn
	}

	return &Driver{
		client: client,
		dbName: config.DBName,
	}, nil
}

func (d *Driver) Client() *mongo.Client {
	return d.client
}

func (d *Driver) DbName() string {
	return d.dbName
}

func (d *Driver) Close() error {
	return d.client.Disconnect(context.Background())
}
