package edgedbtest_test

import (
	"context"
	"fmt"
	"github.com/ibrhmkoz/edgedbtest"
	"github.com/testcontainers/testcontainers-go"
	"log"
)

func ExampleRun() {
	ctx := context.Background()

	edgeDBContainer, err := edgedbtest.RunLatest(ctx)
	defer func() {
		if err := testcontainers.TerminateContainer(edgeDBContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	state, err := edgeDBContainer.State(ctx)
	if err != nil {
		log.Printf("failed to get container state: %s", err)
		return
	}

	fmt.Println(state.Running)
	// Output: true
}
