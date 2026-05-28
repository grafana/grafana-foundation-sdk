// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class RepeatOptionsBuilder implements cog.Builder<dashboardv2.RepeatOptions> {
    protected readonly internal: dashboardv2.RepeatOptions;

    constructor() {
        this.internal = dashboardv2.defaultRepeatOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.RepeatOptions {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    direction(direction: "h" | "v"): this {
        this.internal.direction = direction;
        return this;
    }

    maxPerRow(maxPerRow: number): this {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
}

