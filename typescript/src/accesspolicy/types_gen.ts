// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface AccessPolicy {
	// The scope where these policies should apply
	scope: ResourceRef;
	// The role that must apply this policy
	role: RoleRef;
	// The set of rules to apply.  Note that * is required to modify
	// access policy rules, and that "none" will reject all actions
	rules: AccessRule[];
}

export const defaultAccessPolicy = (): AccessPolicy => ({
	scope: defaultResourceRef(),
	role: defaultRoleRef(),
	rules: [],
});

export interface RoleRef {
	// Policies can apply to roles, teams, or users
	// Applying policies to individual users is supported, but discouraged
	kind: "Role" | "BuiltinRole" | "Team" | "User";
	name: string;
	xname: string;
}

export const defaultRoleRef = (): RoleRef => ({
	kind: "Role",
	name: "",
	xname: "",
});

export interface ResourceRef {
	kind: string;
	name: string;
}

export const defaultResourceRef = (): ResourceRef => ({
	kind: "",
	name: "",
});

export interface AccessRule {
	// The kind this rule applies to (dashboards, alert, etc)
	kind: "*" | string;
	// READ, WRITE, CREATE, DELETE, ...
	// should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
	verb: "*" | "none" | string;
	// Specific sub-elements like "alert.rules" or "dashboard.permissions"????
	target?: string;
}

export const defaultAccessRule = (): AccessRule => ({
	kind: "*",
	verb: "*",
});

