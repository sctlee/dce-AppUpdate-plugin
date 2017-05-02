package update

import (
    "context"

    "github.com/sctlee/app_update/client"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/swarm"
)

// SwarmStackNamespaceLabel is app name
const SwarmStackNamespaceLabel = "com.docker.stack.namespace"

// ServiceUpdateSpec ...
type ServiceUpdateSpec struct {
    ID      string      `json:"ID"`
    Image   string      `json:"Image"`
}

// GetApp ...
func GetApp(appName string) ([]swarm.Service, error) {
    ctx := context.Background()

    cli := client.GetDockerClient()
    services, err := cli.ServiceList(ctx, types.ServiceListOptions{})
    if err != nil {
        return nil, err
    }

    filteredServices := make([]swarm.Service, 0)
    for _, service := range services {
        if ns, ok := service.Spec.Labels[SwarmStackNamespaceLabel]; ok {
            if ns == appName {
                filteredServices = append(filteredServices, service)
            }
        }
    }

    return filteredServices, nil
}

// UpdateApp ...
func UpdateApp(appName string, servicesSpec []ServiceUpdateSpec) error {
    return nil
}
