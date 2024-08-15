create type user_roles as enum('super_admin', 'admin', 'chief', 'user');

CREATE TABLE if not exists users (
    id           UUID PRIMARY KEY default gen_random_uuid(),
    phone_number VARCHAR(100) UNIQUE NOT NULL,
    full_name    VARCHAR(100) NOT NULL,
    birthday     date,
    user_role    user_roles default 'user' NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP,
    deleted_at   TIMESTAMP
);

CREATE TABLE if not exists workers_of_branches (
    id         UUID PRIMARY KEY default gen_random_uuid(),
    user_id    UUID not null references users(id),
    worker_id  serial not null,
    branch_id  UUID not null,
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at timestamp
);

alter sequence workers_of_branches_worker_id_seq restart with 1000;

CREATE TABLE if not exists refresh_tokens (
    id            UUID PRIMARY KEY default gen_random_uuid(),
    user_id       UUID references users(id),
    refresh_token text not null,
    expires_at    TIMESTAMP not null,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE if not exists user_locations (
    id               UUID PRIMARY KEY default gen_random_uuid(),
    user_id          UUID not null references users(id),
    address          varchar not null,
    home_number      varchar,
    floor_number     int,
    apartment_number int,
    padyezd_number   int,
    latitute         double precision,
    longitute        double precision,
    created_at       timestamp default CURRENT_TIMESTAMP,
    updated_at       timestamp,
    deleted_at       timestamp
);