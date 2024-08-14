create type user_roles as enum('admin', 'chief', 'user');

CREATE TABLE if not exists users (
    id UUID PRIMARY KEY default gen_random_uuid(),
    phone_number VARCHAR(100) UNIQUE NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    birthday date,
    user_role user_roles default 'user' NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE if not exists workers_of_branches (
    id UUID PRIMARY KEY default gen_random_uuid(),
    user_id UUID not null references users(id),
    branch_id UUID not null, -- [ref: > branches.id, not null] 
    created_at timestamp default now(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE if not exists refresh_tokens (
    id UUID PRIMARY KEY default gen_random_uuid(),
    user_id UUID references users(id),
    refresh_token text not null,
    expires_at TIMESTAMP not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
