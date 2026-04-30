// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class RowRepeatOptionsBuilder implements cog.Builder<dashboardv2.RowRepeatOptions> {
    protected readonly internal: dashboardv2.RowRepeatOptions;

    constructor() {
        this.internal = dashboardv2.defaultRowRepeatOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.RowRepeatOptions {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

