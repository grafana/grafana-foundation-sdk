// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class TimeIntervalBuilder implements cog.Builder<alerting.TimeInterval> {
    protected readonly internal: alerting.TimeInterval;

    constructor() {
        this.internal = alerting.defaultTimeInterval();
    }

    /**
     * Builds the object.
     */
    build(): alerting.TimeInterval {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    timeIntervals(timeIntervals: cog.Builder<alerting.TimeIntervalItem>[]): this {
        const timeIntervalsResources = timeIntervals.map(builder1 => builder1.build());
        this.internal.time_intervals = timeIntervalsResources;
        return this;
    }
}
