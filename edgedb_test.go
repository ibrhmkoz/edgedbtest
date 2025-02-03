package edgedbtest_test

import (
	"context"
	"edgedbtest"
	"testing"
)

func TestSelect(t *testing.T) {
	ctx := context.Background()
	client := edgedbtest.New(t, ctx)

	var result int64
	err := client.QuerySingle(ctx, `select 1`, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result != 1 {
		t.Fatalf("expected 1, got %d", result)
	}
}
