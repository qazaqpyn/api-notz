package model

type Folder struct {
	Id        string  `json:"id,omitempty"`
	Name      string  `json:"name"`
	ParentId  string  `json:"parent_id"`
	IsRoot    bool    `json:"is_root"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeleteAt  *string `json:"delete_at,omitempty"`
	CreatedBy string  `json:"created_by"`
}

type UpdateFolder struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}
