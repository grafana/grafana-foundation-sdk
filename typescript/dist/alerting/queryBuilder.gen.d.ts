import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class QueryBuilder implements cog.Builder<alerting.Query> {
    protected readonly internal: alerting.Query;
    constructor(refId: string);
    /**
     * Builds the object.
     */
    build(): alerting.Query;
    datasourceUid(datasourceUid: string): this;
    model(model: cog.Builder<cog.Dataquery>): this;
    queryType(queryType: string): this;
    refId(refId: string): this;
    relativeTimeRange(relativeTimeRange: alerting.RelativeTimeRange): this;
}
