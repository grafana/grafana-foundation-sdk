// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class IntervalTriggerBuilder implements cog.Builder<alertingv1beta1.IntervalTrigger> {
    protected readonly internal: alertingv1beta1.IntervalTrigger;

    constructor() {
        this.internal = alertingv1beta1.defaultIntervalTrigger();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.IntervalTrigger {
        return this.internal;
    }

    interval(interval: alertingv1beta1.PromDuration): this {
        this.internal.interval = interval;
        return this;
    }
}

