CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) DEFAULT 'user' NOT NULL,
    last_login_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);



CREATE TABLE IF NOT EXISTS habit_lists (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    is_archived BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_user_list_name UNIQUE (user_id, name)
);

CREATE INDEX idx_habit_lists_user_id ON habit_lists(user_id);

CREATE TABLE IF NOT EXISTS habits (
    id SERIAL PRIMARY KEY,
    habit_list_id INTEGER NOT NULL REFERENCES habit_lists(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    default_emoji VARCHAR(10) DEFAULT '✅' NOT NULL,
    is_archived BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_list_habit_name UNIQUE (habit_list_id, name)
);

CREATE INDEX idx_habits_habit_list_id ON habits(habit_list_id);

CREATE TABLE IF NOT EXISTS habit_checks (
    id SERIAL PRIMARY KEY,
    habit_id INTEGER NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    emoji VARCHAR(10) DEFAULT '✅' NOT NULL,
    note TEXT NULL,
    is_completed BOOLEAN DEFAULT TRUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_habit_date UNIQUE (habit_id, date)
);

CREATE INDEX idx_habit_checks_habit_id ON habit_checks(habit_id);
CREATE INDEX idx_habit_checks_date ON habit_checks(date);
