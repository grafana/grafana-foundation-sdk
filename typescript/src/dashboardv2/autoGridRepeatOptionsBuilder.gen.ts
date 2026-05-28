// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class AutoGridRepeatOptionsBuilder implements cog.Builder<dashboardv2.AutoGridRepeatOptions> {
    protected readonly internal: dashboardv2.AutoGridRepeatOptions;

    constructor() {
        this.internal = dashboardv2.defaultAutoGridRepeatOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.AutoGridRepeatOptions {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

