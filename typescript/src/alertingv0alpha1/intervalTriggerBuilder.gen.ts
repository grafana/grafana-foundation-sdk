// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class IntervalTriggerBuilder implements cog.Builder<alertingv0alpha1.IntervalTrigger> {
    protected readonly internal: alertingv0alpha1.IntervalTrigger;

    constructor() {
        this.internal = alertingv0alpha1.defaultIntervalTrigger();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.IntervalTrigger {
        return this.internal;
    }

    interval(interval: alertingv0alpha1.PromDuration): this {
        this.internal.interval = interval;
        return this;
    }
}

