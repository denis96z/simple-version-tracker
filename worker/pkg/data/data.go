package data

type ExternalProjectInfo struct {
	ID            uint32
	Name          string
	LatestVersion string
	CheckerImage  DockerImageInfo
}

type DockerImageInfo struct {
	Name              string
	Registry          DockerRegistryInfo
	AccessCredentials DockerRegistryCredentials
}

type DockerRegistryInfo struct {
	Host string
}

type DockerRegistryCredentials struct {
	Username string
	Password string
}
