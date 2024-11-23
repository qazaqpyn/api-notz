-- Users Table
DROP INDEX IF EXISTS idx_users_email;

-- Notes Table
DROP INDEX IF EXISTS idx_notes_created_by;
DROP INDEX IF EXISTS idx_notes_updated_by;
DROP INDEX IF EXISTS idx_notes_deleted_by;
DROP INDEX IF EXISTS idx_notes_created_by_deleted_by;

-- Audio Files Table
DROP INDEX IF EXISTS idx_audio_files_note_id;
DROP INDEX IF EXISTS idx_audio_files_note_id_deleted_at;

-- Tags Table
DROP INDEX IF EXISTS idx_tags_created_by;

-- Folders Table
DROP INDEX IF EXISTS idx_folders_created_by;
DROP INDEX IF EXISTS idx_folders_parent_id;
DROP INDEX IF EXISTS idx_folders_parent_id_created_by;

-- Notes Tags Table
DROP INDEX IF EXISTS idx_notes_tags_note_id;
DROP INDEX IF EXISTS idx_notes_tags_tag_id;

-- Notes Folders Table
DROP INDEX IF EXISTS idx_notes_folders_parent;
DROP INDEX IF EXISTS idx_notes_folders_child;

-- User Activity Table
DROP INDEX IF EXISTS idx_users_activity_user_id;
DROP INDEX IF EXISTS idx_users_activity_created_at;
DROP INDEX IF EXISTS idx_user_activity_user_id_created_at;