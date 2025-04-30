-- Drop triggers first
DROP TRIGGER IF EXISTS update_habit_checks_updated_at ON habit_checks;
DROP TRIGGER IF EXISTS update_habits_updated_at ON habits;
DROP TRIGGER IF EXISTS update_habit_lists_updated_at ON habit_lists;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables in correct order
DROP TABLE IF EXISTS habit_checks;
DROP TABLE IF EXISTS habits;
DROP TABLE IF EXISTS habit_lists;
DROP TABLE IF EXISTS users;