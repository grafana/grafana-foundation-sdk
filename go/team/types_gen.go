// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package team

type Team struct {
	// Name of the team.
	Name string `json:"name"`
	// Email of the team.
	Email *string `json:"email,omitempty"`
}

func (resource Team) Equals(other Team) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Email == nil && other.Email != nil || resource.Email != nil && other.Email == nil {
		return false
	}

	if resource.Email != nil {
		if *resource.Email != *other.Email {
			return false
		}
	}

	return true
}
