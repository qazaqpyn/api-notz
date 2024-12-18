package model

type AudioFiles struct {
	Id        string  `json:"id,omitempty"`
	FileName  string  `json:"file_name"`
	FilePath  string  `json:"file_path"`
	NoteId    string  `json:"note_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeleteAt  *string `json:"delete_at,omitempty"`
}
