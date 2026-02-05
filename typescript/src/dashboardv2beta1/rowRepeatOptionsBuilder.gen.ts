// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class RowRepeatOptionsBuilder implements cog.Builder<dashboardv2beta1.RowRepeatOptions> {
    protected readonly internal: dashboardv2beta1.RowRepeatOptions;

    constructor() {
        this.internal = dashboardv2beta1.defaultRowRepeatOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowRepeatOptions {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

