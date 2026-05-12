// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class TabRepeatOptionsBuilder implements cog.Builder<dashboardv2.TabRepeatOptions> {
    protected readonly internal: dashboardv2.TabRepeatOptions;

    constructor() {
        this.internal = dashboardv2.defaultTabRepeatOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.TabRepeatOptions {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

