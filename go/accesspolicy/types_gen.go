// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

type AccessPolicy struct {
	// The scope where these policies should apply
	Scope ResourceRef `json:"scope"`
	// The role that must apply this policy
	Role RoleRef `json:"role"`
	// The set of rules to apply.  Note that * is required to modify
	// access policy rules, and that "none" will reject all actions
	Rules []AccessRule `json:"rules"`
}

func (resource AccessPolicy) Equals(other AccessPolicy) bool {
	if !resource.Scope.Equals(other.Scope) {
		return false
	}
	if !resource.Role.Equals(other.Role) {
		return false
	}

	if len(resource.Rules) != len(other.Rules) {
		return false
	}

	for i1 := range resource.Rules {
		if !resource.Rules[i1].Equals(other.Rules[i1]) {
			return false
		}
	}

	return true
}

type RoleRef struct {
	// Policies can apply to roles, teams, or users
	// Applying policies to individual users is supported, but discouraged
	Kind  RoleRefKind `json:"kind"`
	Name  string      `json:"name"`
	Xname string      `json:"xname"`
}

func (resource RoleRef) Equals(other RoleRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.Xname != other.Xname {
		return false
	}

	return true
}

type ResourceRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

func (resource ResourceRef) Equals(other ResourceRef) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Name != other.Name {
		return false
	}

	return true
}

type AccessRule struct {
	// The kind this rule applies to (dashboards, alert, etc)
	Kind string `json:"kind"`
	// READ, WRITE, CREATE, DELETE, ...
	// should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
	Verb string `json:"verb"`
	// Specific sub-elements like "alert.rules" or "dashboard.permissions"????
	Target *string `json:"target,omitempty"`
}

func (resource AccessRule) Equals(other AccessRule) bool {
	if resource.Kind != other.Kind {
		return false
	}
	if resource.Verb != other.Verb {
		return false
	}
	if resource.Target == nil && other.Target != nil || resource.Target != nil && other.Target == nil {
		return false
	}

	if resource.Target != nil {
		if *resource.Target != *other.Target {
			return false
		}
	}

	return true
}

type RoleRefKind string

const (
	RoleRefKindRole        RoleRefKind = "Role"
	RoleRefKindBuiltinRole RoleRefKind = "BuiltinRole"
	RoleRefKindTeam        RoleRefKind = "Team"
	RoleRefKindUser        RoleRefKind = "User"
)
