package checker

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/denis96z/simple-version-tracker/worker/pkg/data"
)

type Checker struct {
	conf Config
}

func NewChecker(conf Config) *Checker {
	return &Checker{
		conf: conf,
	}
}

func (c *Checker) GetLatestVersion(ctx context.Context, img data.DockerImageInfo) (string, error) {
	name := imgName(img)
	confFilePath := path.Join(
		c.conf.Docker.ConfigDirectoryPath, sha256Sum(name)+".json",
	)

	cmd := exec.CommandContext(
		ctx, c.conf.Docker.BinaryPath, "--config", confFilePath,
		"login", img.Registry.Host, "-u", img.AccessCredentials.Username, "-p", img.AccessCredentials.Password,
	)

	var esb strings.Builder
	cmd.Stderr = &esb

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(
			"failed to login [registry = %q, username = %q] (%s): %w",
			img.Registry.Host, img.AccessCredentials.Username, esb.String(), err,
		)
	}

	cmd = exec.CommandContext(
		ctx, c.conf.Docker.BinaryPath, "--config", confFilePath,
		"pull", name,
	)

	esb.Reset()
	cmd.Stderr = &esb

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(
			"failed to pull image [image = %q, username = %q] (%s): %w",
			name, img.AccessCredentials.Username, esb.String(), err,
		)
	}

	cmd = exec.CommandContext(
		ctx, c.conf.Docker.BinaryPath, "--config", confFilePath,
		"run", "--rm", name,
	)

	esb.Reset()
	cmd.Stderr = &esb

	var sb strings.Builder
	cmd.Stdout = &sb

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf(
			"failed to check version [image = %q] (%s): %w",
			name, esb.String(), err,
		)
	}

	return sb.String(), nil
}

func imgName(img data.DockerImageInfo) string {
	return img.Registry.Host + "/" + img.Name
}

func sha256Sum(imgName string) string {
	b := sha256.Sum256([]byte(imgName))
	return string(b[:])
}
