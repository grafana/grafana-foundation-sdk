import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class QueryOptionsSpecBuilder implements cog.Builder<dashboardv2beta1.QueryOptionsSpec> {
    protected readonly internal: dashboardv2beta1.QueryOptionsSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.QueryOptionsSpec;
    timeFrom(timeFrom: string): this;
    maxDataPoints(maxDataPoints: number): this;
    timeShift(timeShift: string): this;
    queryCachingTTL(queryCachingTTL: number): this;
    interval(interval: string): this;
    cacheTimeout(cacheTimeout: string): this;
    hideTimeOverride(hideTimeOverride: boolean): this;
    timeCompare(timeCompare: string): this;
}
