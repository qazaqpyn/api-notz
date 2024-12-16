package model

type Note struct {
	Id         string `json:"id,omitempty"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Summary    string `json:"summary"`
	Transcript string `json:"transcript"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeleteAt   string `json:"delete_at,omitempty"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
	DeletedBy  string `json:"deleted_by,omitempty"`
}

type UpdateNoteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNoteTagsInput struct {
	Added   []string
	Deleted []string
}
