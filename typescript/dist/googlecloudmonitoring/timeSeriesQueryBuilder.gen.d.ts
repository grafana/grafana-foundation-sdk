import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';
export declare class TimeSeriesQueryBuilder implements cog.Builder<googlecloudmonitoring.TimeSeriesQuery> {
    protected readonly internal: googlecloudmonitoring.TimeSeriesQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.TimeSeriesQuery;
    projectName(projectName: string): this;
    query(query: string): this;
    graphPeriod(graphPeriod: string): this;
}
