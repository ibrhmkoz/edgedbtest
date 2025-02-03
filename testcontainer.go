package edgedbtest

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"net"
)

const (
	defaultUser = "edgedb"
	defaultPass = "edgedb"
	defaultDB   = "main"
)

type EdgeDBContainer struct {
	testcontainers.Container
}

func (c *EdgeDBContainer) DSN(ctx context.Context) (string, error) {
	containerPort, err := c.MappedPort(ctx, "5656/tcp")
	if err != nil {
		return "", err
	}

	host, err := c.Host(ctx)
	if err != nil {
		return "", err
	}

	connStr := fmt.Sprintf("edgedb://%s:%s@%s/%s", defaultUser, defaultPass, net.JoinHostPort(host, containerPort.Port()), defaultDB)
	return connStr, nil
}

func RunLatest(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*EdgeDBContainer, error) {
	return Run(ctx, "edgedb/edgedb:latest", opts...)
}

// Run creates an instance of the SurrealDB container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*EdgeDBContainer, error) {
	req := testcontainers.ContainerRequest{
		Image: img,
		Env: map[string]string{
			"EDGEDB_SERVER_SECURITY": "insecure_dev_mode",
			"EDGEDB_SERVER_USER":     defaultUser,
			"EDGEDB_SERVER_PASSWORD": defaultPass,
		},
		ExposedPorts: []string{"5656/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForLog("Serving on "),
		),
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	for _, opt := range opts {
		if err := opt.Customize(&genericContainerReq); err != nil {
			return nil, fmt.Errorf("customize: %w", err)
		}
	}

	container, err := testcontainers.GenericContainer(ctx, genericContainerReq)
	var c *EdgeDBContainer
	if container != nil {
		c = &EdgeDBContainer{Container: container}
	}

	if err != nil {
		return c, fmt.Errorf("generic container: %w", err)
	}

	return c, nil
}
