package res

type ListTotalDto[T any] struct {
	Contents []T `json:"contents"`
	Total    int `json:"total"`
}

func ToListTotalDto[T any](contents []T, total int) ListTotalDto[T] {
	return ListTotalDto[T]{
		Contents: contents,
		Total:    total,
	}
}

type RoleDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
