CREATE TABLE users (
    id SERIAL PRIMARY KEY,               -- Unique user ID
    username VARCHAR(255) NOT NULL,      -- User's chosen username
    email VARCHAR(255) UNIQUE NOT NULL,   -- User's email address
    password_hash VARCHAR(255) NOT NULL,  -- Hashed password
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- When the user account was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Last time the user details were updated
);

CREATE TABLE links (
    id SERIAL PRIMARY KEY,               -- Unique link ID
    user_id INTEGER REFERENCES users(id), -- Foreign key to the users table
    long_url TEXT NOT NULL,              -- The original URL (long URL)
    short_code VARCHAR(6) UNIQUE NOT NULL, -- The unique short code (e.g. 'abc123')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- When the URL was shortened
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- When the shortened URL details were updated
);

CREATE TABLE redirect_events (
    id SERIAL PRIMARY KEY,               -- Unique redirect entry ID
    url_id INTEGER REFERENCES links(id),  -- Foreign key to the links table
    user_ip VARCHAR(45),                 -- IP address of the user who clicked the short URL
    user_agent TEXT,                     -- User agent (browser/OS info)
    referrer TEXT,                       -- Referring URL (if any)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp of the redirect event
);