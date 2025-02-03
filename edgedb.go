package edgedbtest

import (
	"context"
	"github.com/edgedb/edgedb-go"
	"testing"
)

func New(t *testing.T, ctx context.Context) *edgedb.Client {
	container, err := RunLatest(ctx)
	if err != nil {
		return nil
	}

	dsn, err := container.DSN(ctx)
	if err != nil {
		t.Fatal(err)
	}

	options := edgedb.Options{
		TLSOptions: edgedb.TLSOptions{
			SecurityMode: edgedb.TLSModeInsecure,
		},
	}
	client, err := edgedb.CreateClientDSN(ctx, dsn, options)
	if err != nil {
		t.Fatal(err)
	}

	err = client.EnsureConnected(ctx)
	if err != nil {
		t.Fatal(err)
	}
	return client
}
