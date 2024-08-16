INSERT INTO users (
    phone_number, 
    full_name, 
    birthday, 
    user_role, 
    created_at, 
    updated_at, 
    deleted_at
) VALUES (
    '+998937587909', 
    'Ko''palov Muhammadjon', 
    '2005-11-01', 
    'super_admin', 
    CURRENT_TIMESTAMP, 
    NULL, 
    NULL
),
(
    '+998997693731', 
    'Saidakbar Pardaboyev', 
    '2005-06-11', 
    'super_admin', 
    CURRENT_TIMESTAMP, 
    NULL, 
    NULL
);

INSERT INTO users (
    id,
    phone_number, 
    full_name, 
    birthday, 
    user_role, 
    created_at, 
    updated_at, 
    deleted_at
) VALUES (
    '024fcac1-b077-4d5c-b22a-671e01904025',
    '+9989976937345', 
    'test', 
    '2005-06-11', 
    'user',
    CURRENT_TIMESTAMP,
    NULL,
    NULL
);
