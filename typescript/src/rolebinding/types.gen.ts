// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface RoleBinding {
	// The role we are discussing
	role: BuiltinRoleRef | CustomRoleRef;
	// The team or user that has the specified role
	subject: RoleBindingSubject;
}

export const defaultRoleBinding = (): RoleBinding => ({
	role: defaultBuiltinRoleRef(),
	subject: defaultRoleBindingSubject(),
});

export interface CustomRoleRef {
	kind: "Role";
	name: string;
}

export const defaultCustomRoleRef = (): CustomRoleRef => ({
	kind: "Role",
	name: "",
});

export interface BuiltinRoleRef {
	kind: "BuiltinRole";
	name: "viewer" | "editor" | "admin";
}

export const defaultBuiltinRoleRef = (): BuiltinRoleRef => ({
	kind: "BuiltinRole",
	name: "viewer",
});

export interface RoleBindingSubject {
	kind: "Team" | "User";
	// The team/user identifier name
	name: string;
}

export const defaultRoleBindingSubject = (): RoleBindingSubject => ({
	kind: "Team",
	name: "",
});

