// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as rolebinding from '../rolebinding';

export class BuiltinRoleRefBuilder implements cog.Builder<rolebinding.BuiltinRoleRef> {
    protected readonly internal: rolebinding.BuiltinRoleRef;

    constructor() {
        this.internal = rolebinding.defaultBuiltinRoleRef();
        this.internal.kind = "BuiltinRole";
    }

    /**
     * Builds the object.
     */
    build(): rolebinding.BuiltinRoleRef {
        return this.internal;
    }

    name(name: "viewer" | "editor" | "admin"): this {
        this.internal.name = name;
        return this;
    }
}
