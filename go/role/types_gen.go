// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package role

type Role struct {
	// The role identifier `managed:builtins:editor:permissions`
	Name string `json:"name"`
	// Optional display
	DisplayName *string `json:"displayName,omitempty"`
	// Name of the team.
	GroupName *string `json:"groupName,omitempty"`
	// Role description
	Description *string `json:"description,omitempty"`
	// Do not show this role
	Hidden bool `json:"hidden"`
}

func (resource Role) Equals(other Role) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.DisplayName == nil && other.DisplayName != nil || resource.DisplayName != nil && other.DisplayName == nil {
		return false
	}

	if resource.DisplayName != nil {
		if *resource.DisplayName != *other.DisplayName {
			return false
		}
	}
	if resource.GroupName == nil && other.GroupName != nil || resource.GroupName != nil && other.GroupName == nil {
		return false
	}

	if resource.GroupName != nil {
		if *resource.GroupName != *other.GroupName {
			return false
		}
	}
	if resource.Description == nil && other.Description != nil || resource.Description != nil && other.Description == nil {
		return false
	}

	if resource.Description != nil {
		if *resource.Description != *other.Description {
			return false
		}
	}
	if resource.Hidden != other.Hidden {
		return false
	}

	return true
}
