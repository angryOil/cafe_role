package res

import "cafe_role/internal/domain"

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

func ToRoleDtoList(domains []domain.Role) []RoleDto {
	results := make([]RoleDto, len(domains))
	for i, d := range domains {
		results[i] = RoleDto{
			Id:          d.Id,
			Name:        d.Name,
			Description: d.Description,
		}
	}
	return results
}
