// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as rolebinding from '../rolebinding';

export class CustomRoleRefBuilder implements cog.OptionsBuilder<rolebinding.CustomRoleRef> {
    private readonly internal: rolebinding.CustomRoleRef;

    constructor() {
        this.internal = rolebinding.defaultCustomRoleRef();
        this.internal.kind = "Role";
    }

    build(): rolebinding.CustomRoleRef {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}
