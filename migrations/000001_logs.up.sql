-- Create logs table
CREATE TABLE logs
(
    id              SERIAL PRIMARY KEY,
    user_id         VARCHAR(36),
    video_id        VARCHAR(36),
    watchtime       INTEGER
);

-- Add indexes to logs table
CREATE INDEX ON logs USING btree (user_id);
CREATE INDEX ON logs USING btree (video_id);
