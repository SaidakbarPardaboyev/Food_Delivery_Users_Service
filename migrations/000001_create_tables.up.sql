create type user_roles as enum('admin', 'chief', 'user');

CREATE TABLE users (
    id UUID PRIMARY KEY default gen_random_uuid(),
    phone_number VARCHAR(100) UNIQUE NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    birthday date,
    user_role user_roles default 'user' NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY default gen_random_uuid(),
    user_id UUID references users(id),
    refresh_token text not null,
    expires_at TIMESTAMP not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
