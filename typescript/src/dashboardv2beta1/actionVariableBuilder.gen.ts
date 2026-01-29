// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ActionVariableBuilder implements cog.Builder<dashboardv2beta1.ActionVariable> {
    protected readonly internal: dashboardv2beta1.ActionVariable;

    constructor() {
        this.internal = dashboardv2beta1.defaultActionVariable();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ActionVariable {
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

