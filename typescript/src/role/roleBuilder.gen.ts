// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as role from '../role';

export class RoleBuilder implements cog.Builder<role.Role> {
    protected readonly internal: role.Role;

    constructor() {
        this.internal = role.defaultRole();
    }

    /**
     * Builds the object.
     */
    build(): role.Role {
        return this.internal;
    }

    // The role identifier `managed:builtins:editor:permissions`
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Optional display
    displayName(displayName: string): this {
        this.internal.displayName = displayName;
        return this;
    }

    // Name of the team.
    groupName(groupName: string): this {
        this.internal.groupName = groupName;
        return this;
    }

    // Role description
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // Do not show this role
    hidden(hidden: boolean): this {
        this.internal.hidden = hidden;
        return this;
    }
}
