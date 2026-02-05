import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class MatcherBuilder implements cog.Builder<alerting.Matcher> {
    protected readonly internal: alerting.Matcher;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.Matcher;
    name(name: string): this;
    type(type: alerting.MatchType): this;
    value(value: string): this;
}
