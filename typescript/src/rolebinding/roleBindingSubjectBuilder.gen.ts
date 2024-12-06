// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as rolebinding from '../rolebinding';

export class RoleBindingSubjectBuilder implements cog.Builder<rolebinding.RoleBindingSubject> {
    protected readonly internal: rolebinding.RoleBindingSubject;

    constructor() {
        this.internal = rolebinding.defaultRoleBindingSubject();
    }

    /**
     * Builds the object.
     */
    build(): rolebinding.RoleBindingSubject {
        return this.internal;
    }

    kind(kind: "Team" | "User"): this {
        this.internal.kind = kind;
        return this;
    }

    // The team/user identifier name
    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}
