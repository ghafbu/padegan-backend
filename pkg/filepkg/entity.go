package filepkg

type FileEntity struct {
	FullName    string `json:"fullName"`
	ContentType string `json:"content_type"`
	Format      string `json:"format"`
	Name        string `json:"name"`
	Hash        string `json:"hash"`
	Size        int64  `json:"size"`
}
