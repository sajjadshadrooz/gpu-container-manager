package handlers

import (
    "context"
    "encoding/json"
    "fmt"
    "gpu-container-manager/models"
    "net/http"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

func CreateContainer(w http.ResponseWriter, r *http.Request) {
    var req models.ContainerRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    ctx := context.Background()
    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: req.Image,
        Env:   req.EnvVars,
        Cmd:   req.Command,
        Labels: map[string]string{
            "gpu_count": fmt.Sprintf("%d", req.GPUCount),
        },
    }, nil, nil, nil, req.Name)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"container_id": resp.ID})
}

func UpdateContainer(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Updating container is not supported in Docker directly. Consider recreating the container with desired configurations.", http.StatusNotImplemented)
}

func DeleteContainer(w http.ResponseWriter, r *http.Request) {
    containerID := r.URL.Query().Get("id")
    if containerID == "" {
        http.Error(w, "Container ID is required", http.StatusBadRequest)
        return
    }

    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    ctx := context.Background()
    if err := cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true}); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Container deleted successfully"})
}
