// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package synclet

import (
	"context"

	"github.com/google/go-cloud/wire"
	"github.com/windmilleng/tilt/internal/build"
	"github.com/windmilleng/tilt/internal/k8s"
)

func WireSynclet(ctx context.Context, env k8s.Env) (*Synclet, error) {
	wire.Build(
		build.DefaultDockerClient,
		wire.Bind(new(build.DockerClient), new(build.DockerCli)),

		NewSynclet,
	)
	return nil, nil
}
