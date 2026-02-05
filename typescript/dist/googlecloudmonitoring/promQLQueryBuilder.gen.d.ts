import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';
export declare class PromQLQueryBuilder implements cog.Builder<googlecloudmonitoring.PromQLQuery> {
    protected readonly internal: googlecloudmonitoring.PromQLQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.PromQLQuery;
    projectName(projectName: string): this;
    expr(expr: string): this;
    step(step: string): this;
}
