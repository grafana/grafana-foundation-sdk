// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class RelativeTimeRangeBuilder implements cog.Builder<alertingv0alpha1.RelativeTimeRange> {
    protected readonly internal: alertingv0alpha1.RelativeTimeRange;

    constructor() {
        this.internal = alertingv0alpha1.defaultRelativeTimeRange();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.RelativeTimeRange {
        return this.internal;
    }

    from(from: alertingv0alpha1.PromDurationWMillis): this {
        this.internal.from = from;
        return this;
    }

    to(to: alertingv0alpha1.PromDurationWMillis): this {
        this.internal.to = to;
        return this;
    }
}

