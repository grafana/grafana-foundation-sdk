// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ActionVariableBuilder implements cog.Builder<dashboardv2.ActionVariable> {
    protected readonly internal: dashboardv2.ActionVariable;

    constructor() {
        this.internal = dashboardv2.defaultActionVariable();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ActionVariable {
        return this.internal;
    }

    key(key: string): this {
        this.internal.key = key;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}

