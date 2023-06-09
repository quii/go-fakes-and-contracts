package adapters

import (
	"context"
	"fmt"
	"github.com/quii/go-fakes-and-contracts/domain/planner/expect"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	startupTimeout = 5 * time.Second
	dockerfileName = "Dockerfile"
)

func StartDockerServer(
	t testing.TB,
	port string,
	binToBuild string,
) {
	t.Helper()

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: newTCDockerfile(binToBuild),
		ExposedPorts:   []string{fmt.Sprintf("%s:%s", port, port)},
		WaitingFor:     wait.ForListeningPort(nat.Port(port)).WithStartupTimeout(startupTimeout),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	expect.NoErr(t, err)
	t.Cleanup(func() {
		expect.NoErr(t, container.Terminate(ctx))
	})
}

func newTCDockerfile(binToBuild string) testcontainers.FromDockerfile {
	return testcontainers.FromDockerfile{
		Context:    "../../.",
		Dockerfile: dockerfileName,
		BuildArgs: map[string]*string{
			"bin_to_build": &binToBuild,
		},
		PrintBuildLog: true,
	}
}
