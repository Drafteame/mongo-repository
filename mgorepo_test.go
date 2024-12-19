package mgorepo

import (
	"context"
	"strings"
	"testing"

	"github.com/Drafteame/mgorepo/driver"
)

func getTestDriver(t *testing.T) *driver.Driver {
	t.Helper()

	dbName := getTestDB(t)

	d, driverErr := driver.NewWithConfig(context.TODO(), driver.Config{
		URI:            "mongodb://root:root@localhost:27017/" + dbName + "?authSource=admin",
		ReadPreference: "primary",
		DBName:         dbName,
	})

	if driverErr != nil {
		t.Fatal(driverErr)
	}

	t.Cleanup(func() {
		cleanup(t, d)
	})

	return d
}

func cleanup(t *testing.T, d *driver.Driver) {
	t.Helper()

	if err := d.Client().Database(d.DbName()).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}

	if err := d.Close(); err != nil {
		t.Fatal(err)
	}
}

func getTestDB(t *testing.T) string {
	t.Helper()

	slug := t.Name()

	replace := []string{"/", " "}

	for _, rep := range replace {
		slug = strings.ReplaceAll(slug, rep, "")
	}

	slug = "test_" + slug

	// trim to 60 chars
	if len(slug) > 60 {
		slug = slug[:60]
	}
	return slug
}
