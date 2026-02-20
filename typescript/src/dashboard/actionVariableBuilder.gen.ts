// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

export class ActionVariableBuilder implements cog.Builder<dashboard.ActionVariable> {
    protected readonly internal: dashboard.ActionVariable;

    constructor() {
        this.internal = dashboard.defaultActionVariable();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.ActionVariable {
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

