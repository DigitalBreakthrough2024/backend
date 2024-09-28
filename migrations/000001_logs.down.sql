-- Drop indexes from logs table
DROP INDEX IF EXISTS logs_user_id_idx;
DROP INDEX IF EXISTS logs_video_id_idx;

-- Drop logs table
DROP TABLE IF EXISTS logs;
