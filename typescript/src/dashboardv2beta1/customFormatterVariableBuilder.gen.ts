// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Custom formatter variable
export class CustomFormatterVariableBuilder implements cog.Builder<dashboardv2beta1.CustomFormatterVariable> {
    protected readonly internal: dashboardv2beta1.CustomFormatterVariable;

    constructor() {
        this.internal = dashboardv2beta1.defaultCustomFormatterVariable();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.CustomFormatterVariable {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    type(type: dashboardv2beta1.VariableType): this {
        this.internal.type = type;
        return this;
    }

    multi(multi: boolean): this {
        this.internal.multi = multi;
        return this;
    }

    includeAll(includeAll: boolean): this {
        this.internal.includeAll = includeAll;
        return this;
    }
}

