// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as accesspolicy from '../accesspolicy';

export class RoleRefBuilder implements cog.Builder<accesspolicy.RoleRef> {
    protected readonly internal: accesspolicy.RoleRef;

    constructor() {
        this.internal = accesspolicy.defaultRoleRef();
    }

    /**
     * Builds the object.
     */
    build(): accesspolicy.RoleRef {
        return this.internal;
    }

    // Policies can apply to roles, teams, or users
    // Applying policies to individual users is supported, but discouraged
    kind(kind: "Role" | "BuiltinRole" | "Team" | "User"): this {
        this.internal.kind = kind;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    xname(xname: string): this {
        this.internal.xname = xname;
        return this;
    }
}
