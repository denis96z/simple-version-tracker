CREATE TABLE project (
    id SERIAL PRIMARY KEY,
    name VARCHAR(127) NOT NULL UNIQUE
);

CREATE TABLE external_project (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    latest_version VARCHAR(64) NOT NULL
);

CREATE TABLE docker_registry (
    id SERIAL PRIMARY KEY,
    host VARCHAR(127) NOT NULL UNIQUE
);

CREATE TABLE docker_registry_credentials (
    id SERIAL PRIMARY KEY,
    registry_id INTEGER NOT NULL REFERENCES docker_registry(id),
    username VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    UNIQUE (registry_id, username)
);

CREATE TABLE docker_image (
    id SERIAL PRIMARY KEY,
    name VARCHAR(127) NOT NULL,
    registry_id INTEGER NOT NULL REFERENCES docker_registry(id),
    access_credentials_id INTEGER NOT NULL REFERENCES docker_registry_credentials(id),
    UNIQUE (name, registry_id)
);

CREATE TABLE external_project_version_check (
    id SERIAL PRIMARY KEY,
    external_project_id INTEGER NOT NULL UNIQUE REFERENCES external_project(id),
    last_check_ts TIMESTAMP NOT NULL DEFAULT NOW(),
    check_interval_seconds INTEGER NOT NULL,
    script_docker_image_id INTEGER NOT NULL REFERENCES docker_image(id)
);

CREATE TABLE project_external_project_version (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL REFERENCES project(id),
    external_project_id INTEGER NOT NULL REFERENCES external_project(id),
    current_version VARCHAR(64) NOT NULL,
    UNIQUE (project_id, external_project_id)
);
