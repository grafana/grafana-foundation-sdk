import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class TimePickerBuilder implements cog.Builder<dashboard.TimePickerConfig> {
    protected readonly internal: dashboard.TimePickerConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.TimePickerConfig;
    hidden(hidden: boolean): this;
    refreshIntervals(refreshIntervals: string[]): this;
    quickRanges(quickRanges: cog.Builder<dashboard.TimeOption>[]): this;
    nowDelay(nowDelay: string): this;
}
