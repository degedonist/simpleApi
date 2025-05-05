CREATE TABLE tasks (
                        id VARCHAR(255) PRIMARY KEY NOT NULL,
                        task VARCHAR(255) NOT NULL,
                        is_done BOOLEAN DEFAULT FALSE,
                        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                        updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                        deleted_at TIMESTAMP DEFAULT NULL
);