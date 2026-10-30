// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class RelativeTimeRangeBuilder implements cog.Builder<alertingv1beta1.RelativeTimeRange> {
    protected readonly internal: alertingv1beta1.RelativeTimeRange;

    constructor() {
        this.internal = alertingv1beta1.defaultRelativeTimeRange();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.RelativeTimeRange {
        return this.internal;
    }

    from(from: alertingv1beta1.PromDurationWMillis): this {
        this.internal.from = from;
        return this;
    }

    to(to: alertingv1beta1.PromDurationWMillis): this {
        this.internal.to = to;
        return this;
    }
}

