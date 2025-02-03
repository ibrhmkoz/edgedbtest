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

func TestSelectFromCustomType(t *testing.T) {
	ctx := context.Background()
	client := edgedbtest.New(t, ctx)

	var result struct {
		Name string `edgedb:"name"`
		Age  int32  `edgedb:"age"`
	}

	err := client.QuerySingle(ctx, `
		with user := (insert User {
		  name := "John Doe",
		  age  := 25
		}), select user {
		  name,
		  age
		};
	`, &result)

	if err != nil {
		t.Fatal(err)
	}

	if result.Name != "John Doe" {
		t.Fatalf("expected John Doe, got %s", result.Name)
	}

	if result.Age != 25 {
		t.Fatalf("expected 25, got %d", result.Age)
	}
}
