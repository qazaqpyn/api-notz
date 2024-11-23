-- Users Table
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Notes Table
CREATE INDEX IF NOT EXISTS idx_notes_created_by ON notes(created_by);
CREATE INDEX IF NOT EXISTS idx_notes_updated_by ON notes(updated_by);
CREATE INDEX IF NOT EXISTS idx_notes_deleted_by ON notes(deleted_by);
CREATE INDEX IF NOT EXISTS idx_notes_created_by_deleted_by ON notes(created_by, deleted_by);

-- Audio Files Table
CREATE INDEX IF NOT EXISTS idx_audio_files_note_id ON audio_files(note_id);
CREATE INDEX IF NOT EXISTS idx_audio_files_note_id_deleted_at ON audio_files(note_id, deleted_at);


-- Tags Table
CREATE INDEX IF NOT EXISTS idx_tags_created_by ON tags(created_by);

-- Folders Table
CREATE INDEX IF NOT EXISTS idx_folders_created_by ON folders(created_by);
CREATE INDEX IF NOT EXISTS idx_folders_parent_id ON folders(parent_id);
CREATE INDEX IF NOT EXISTS idx_folders_parent_id_created_by ON folders(parent_id, created_by);


-- Notes Tags Table
CREATE INDEX IF NOT EXISTS idx_notes_tags_note_id ON notes_tags(note_id);
CREATE INDEX IF NOT EXISTS idx_notes_tags_tag_id ON notes_tags(tag_id);

-- User Activity Table
CREATE INDEX IF NOT EXISTS idx_users_activity_user_id ON users_activity(user_id);
CREATE INDEX IF NOT EXISTS idx_users_activity_created_at ON users_activity(created_at);
CREATE INDEX IF NOT EXISTS idx_user_activity_user_id_created_at ON users_activity(user_id, created_at);