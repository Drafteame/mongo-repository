package driver

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/orlangure/gnomock"
	gnomokmongo "github.com/orlangure/gnomock/preset/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestDriver struct {
	Driver
	testServer *gnomock.Container
}

func NewTest(t *testing.T) (*TestDriver, error) {
	md := &TestDriver{}

	serverPreset := gnomokmongo.Preset()
	server, startErr := gnomock.Start(serverPreset)
	if startErr != nil {
		return nil, startErr
	}

	uri := fmt.Sprintf("mongodb://%s:%v/%s", server.Host, server.Port("default"), "test")
	rc, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := rc.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	dbName := uri[strings.LastIndex(uri, "/")+1:]

	md.client = rc
	md.dbName = dbName
	md.testServer = server

	t.Cleanup(func() {
		if errClose := md.Close(); errClose != nil {
			t.Fatal(errClose)
		}
	})

	return md, nil
}

func (gm *TestDriver) Close() error {
	if err := gm.client.Disconnect(context.Background()); err != nil {
		return err
	}

	return gnomock.Stop(gm.testServer)
}
