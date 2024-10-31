// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class MuteTimingBuilder implements cog.Builder<alerting.MuteTiming> {
    protected readonly internal: alerting.MuteTiming;

    constructor() {
        this.internal = alerting.defaultMuteTiming();
    }

    /**
     * Builds the object.
     */
    build(): alerting.MuteTiming {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    timeIntervals(timeIntervals: cog.Builder<alerting.TimeInterval>[]): this {
        const timeIntervalsResources = timeIntervals.map(builder1 => builder1.build());
        this.internal.time_intervals = timeIntervalsResources;
        return this;
    }
}
