import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class MuteTimingBuilder implements cog.Builder<alerting.MuteTiming> {
    protected readonly internal: alerting.MuteTiming;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.MuteTiming;
    name(name: string): this;
    timeIntervals(timeIntervals: cog.Builder<alerting.TimeInterval>[]): this;
}
