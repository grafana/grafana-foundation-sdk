import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class RecordRuleBuilder implements cog.Builder<alerting.RecordRule> {
    protected readonly internal: alerting.RecordRule;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.RecordRule;
    from(from: string): this;
    metric(metric: string): this;
}
