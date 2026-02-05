// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TabRepeatOptionsBuilder implements cog.Builder<dashboardv2beta1.TabRepeatOptions> {
    protected readonly internal: dashboardv2beta1.TabRepeatOptions;

    constructor() {
        this.internal = dashboardv2beta1.defaultTabRepeatOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabRepeatOptions {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

