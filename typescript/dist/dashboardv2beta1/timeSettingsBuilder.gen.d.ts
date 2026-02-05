import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TimeSettingsBuilder implements cog.Builder<dashboardv2beta1.TimeSettingsSpec> {
    protected readonly internal: dashboardv2beta1.TimeSettingsSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TimeSettingsSpec;
    timezone(timezone: string): this;
    from(from: string): this;
    to(to: string): this;
    autoRefresh(autoRefresh: string): this;
    autoRefreshIntervals(autoRefreshIntervals: string[]): this;
    quickRanges(quickRanges: cog.Builder<dashboardv2beta1.TimeRangeOption>[]): this;
    hideTimepicker(hideTimepicker: boolean): this;
    weekStart(weekStart: "saturday" | "monday" | "sunday"): this;
    fiscalYearStartMonth(fiscalYearStartMonth: number): this;
    nowDelay(nowDelay: string): this;
}
