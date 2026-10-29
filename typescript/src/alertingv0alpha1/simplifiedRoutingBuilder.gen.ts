// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class SimplifiedRoutingBuilder implements cog.Builder<alertingv0alpha1.SimplifiedRouting> {
    protected readonly internal: alertingv0alpha1.SimplifiedRouting;

    constructor() {
        this.internal = alertingv0alpha1.defaultSimplifiedRouting();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.SimplifiedRouting {
        return this.internal;
    }

    receiver(receiver: string): this {
        this.internal.receiver = receiver;
        return this;
    }

    groupBy(groupBy: string[]): this {
        this.internal.groupBy = groupBy;
        return this;
    }

    groupWait(groupWait: alertingv0alpha1.PromDuration): this {
        this.internal.groupWait = groupWait;
        return this;
    }

    groupInterval(groupInterval: alertingv0alpha1.PromDuration): this {
        this.internal.groupInterval = groupInterval;
        return this;
    }

    repeatInterval(repeatInterval: alertingv0alpha1.PromDuration): this {
        this.internal.repeatInterval = repeatInterval;
        return this;
    }

    muteTimeIntervals(muteTimeIntervals: alertingv0alpha1.TimeIntervalRef[]): this {
        this.internal.muteTimeIntervals = muteTimeIntervals;
        return this;
    }

    activeTimeIntervals(activeTimeIntervals: alertingv0alpha1.TimeIntervalRef[]): this {
        this.internal.activeTimeIntervals = activeTimeIntervals;
        return this;
    }
}

