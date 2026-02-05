import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class RuleGroupBuilder implements cog.Builder<alerting.RuleGroup> {
    protected readonly internal: alerting.RuleGroup;
    constructor(title: string);
    /**
     * Builds the object.
     */
    build(): alerting.RuleGroup;
    folderUid(folderUid: string): this;
    interval(interval: alerting.Duration): this;
    rules(rules: cog.Builder<alerting.Rule>[]): this;
    withRule(rule: cog.Builder<alerting.Rule>): this;
    title(title: string): this;
}
