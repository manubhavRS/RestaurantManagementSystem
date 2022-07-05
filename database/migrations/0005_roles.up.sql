CREATE TABLE user_roles(
        user_id UUID NOT NULL REFERENCES users(user_id),
        name TEXT NOT NULL,
        role TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        archived_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
