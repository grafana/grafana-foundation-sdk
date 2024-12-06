// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as rolebinding from '../rolebinding';

export class RoleBindingBuilder implements cog.Builder<rolebinding.RoleBinding> {
    protected readonly internal: rolebinding.RoleBinding;

    constructor() {
        this.internal = rolebinding.defaultRoleBinding();
    }

    /**
     * Builds the object.
     */
    build(): rolebinding.RoleBinding {
        return this.internal;
    }

    // The role we are discussing
    role(role: cog.Builder<rolebinding.BuiltinRoleRef> | cog.Builder<rolebinding.CustomRoleRef>): this {
        const roleResource = role.build();
        this.internal.role = roleResource;
        return this;
    }

    // The team or user that has the specified role
    subject(subject: cog.Builder<rolebinding.RoleBindingSubject>): this {
        const subjectResource = subject.build();
        this.internal.subject = subjectResource;
        return this;
    }
}
