package storage

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/jackc/pgx/v4"

	"github.com/denis96z/simple-version-tracker/worker/pkg/data"
)

type Storage struct {
	conf Config
	conn *pgx.Conn
}

func NewStorage(conf Config) *Storage {
	return &Storage{
		conf: conf,
	}
}

func (s *Storage) Init(ctx context.Context) error {
	conn, err := pgx.Connect(
		ctx, fmt.Sprintf(
			"postgres://%s:%s@%s/%s",
			s.conf.Username, s.conf.Password,
			net.JoinHostPort(
				s.conf.Host, strconv.FormatUint(uint64(s.conf.Port), 10),
			),
			s.conf.Database,
		),
	)
	if err != nil {
		return fmt.Errorf(
			"failed to establish database connection: %w", err,
		)
	}

	s.conn = conn
	return nil
}

func (s *Storage) Finit(ctx context.Context) error {
	if err := s.conn.Close(ctx); err != nil {
		return fmt.Errorf(
			"failed to close database connection: %w", err,
		)
	}
	return nil
}

func (s *Storage) SelectExternalProjectsToBeCheckedForUpdate(ctx context.Context) ([]data.ExternalProjectInfo, error) {
	rows, err := s.conn.Query(
		ctx, `SELECT xp.id, xp.name, xp.latest_version, r.host, c.username, c.password, img.name FROM external_project xp JOIN external_project_version_check chk ON xp.id = chk.external_project_id JOIN docker_image img ON chk.script_docker_image_id = img.id JOIN docker_registry r ON img.registry_id = r.id JOIN docker_registry_credentials c ON img.access_credentials_id = c.id WHERE NOW() > (chk.last_check_ts + (chk.check_interval_seconds * INTERVAL '1 second'))`,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to select external projects to be checked for update: %w", err,
		)
	}

	defer rows.Close()

	arr := make([]data.ExternalProjectInfo, 0)
	for rows.Next() {
		var info data.ExternalProjectInfo
		err = rows.Scan(
			&info.ID,
			&info.Name,
			&info.LatestVersion,
			&info.CheckerImage.Registry.Host,
			&info.CheckerImage.AccessCredentials.Username,
			&info.CheckerImage.AccessCredentials.Password,
			&info.CheckerImage.Name,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to scan external project info: %w", err,
			)
		}
		arr = append(
			arr, info,
		)
	}

	return arr, nil
}
