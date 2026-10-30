// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class SimplifiedRoutingBuilder implements cog.Builder<alertingv1beta1.SimplifiedRouting> {
    protected readonly internal: alertingv1beta1.SimplifiedRouting;

    constructor() {
        this.internal = alertingv1beta1.defaultSimplifiedRouting();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.SimplifiedRouting {
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

    groupWait(groupWait: alertingv1beta1.PromDuration): this {
        this.internal.groupWait = groupWait;
        return this;
    }

    groupInterval(groupInterval: alertingv1beta1.PromDuration): this {
        this.internal.groupInterval = groupInterval;
        return this;
    }

    repeatInterval(repeatInterval: alertingv1beta1.PromDuration): this {
        this.internal.repeatInterval = repeatInterval;
        return this;
    }

    muteTimeIntervals(muteTimeIntervals: alertingv1beta1.TimeIntervalRef[]): this {
        this.internal.muteTimeIntervals = muteTimeIntervals;
        return this;
    }

    activeTimeIntervals(activeTimeIntervals: alertingv1beta1.TimeIntervalRef[]): this {
        this.internal.activeTimeIntervals = activeTimeIntervals;
        return this;
    }
}

