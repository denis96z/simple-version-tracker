CREATE TABLE project (
    id SERIAL PRIMARY KEY,
    name VARCHAR(127) NOT NULL UNIQUE
);

CREATE TABLE external_project (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    latest_version VARCHAR(64) NOT NULL
);

CREATE TABLE project_external_project_version (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES project(id),
    external_project_id INTEGER NOT NULL REFERENCES external_project(id),
    current_version VARCHAR(64) NOT NULL,
    UNIQUE (project_id, external_project_id)
);

CREATE TABLE external_project_version_check (
    id SERIAL PRIMARY KEY,
    external_project_id INTEGER NOT NULL UNIQUE REFERENCES external_project(id),
    check_interval_seconds INTEGER NOT NULL,
    last_check_ts TIMESTAMP NOT NULL DEFAULT NOW(),
    script_docker_image VARCHAR(255) NOT NULL UNIQUE
);
