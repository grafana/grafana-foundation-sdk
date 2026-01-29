// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TimeRangeOptionBuilder implements cog.Builder<dashboardv2beta1.TimeRangeOption> {
    protected readonly internal: dashboardv2beta1.TimeRangeOption;

    constructor() {
        this.internal = dashboardv2beta1.defaultTimeRangeOption();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TimeRangeOption {
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

