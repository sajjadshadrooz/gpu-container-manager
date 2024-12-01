package models

type ContainerRequest struct {
    Name       string   `json:"name"`
    Image      string   `json:"image"`
    GPUCount   int      `json:"gpu_count"`
    EnvVars    []string `json:"env_vars"`
    Command    []string `json:"command"`
}