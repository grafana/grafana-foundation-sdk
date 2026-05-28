// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class TimeRangeOptionBuilder implements cog.Builder<dashboardv2.TimeRangeOption> {
    protected readonly internal: dashboardv2.TimeRangeOption;

    constructor() {
        this.internal = dashboardv2.defaultTimeRangeOption();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.TimeRangeOption {
        return this.internal;
    }

    display(display: string): this {
        this.internal.display = display;
        return this;
    }

    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    to(to: string): this {
        this.internal.to = to;
        return this;
    }
}

