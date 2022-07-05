CREATE TABLE users(
        user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL,
        created_by UUID REFERENCES users(user_id),
        created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
        archived_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
