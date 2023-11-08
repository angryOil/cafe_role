package req

type CreateDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PatchDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
