import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class NotificationPolicyBuilder implements cog.Builder<alerting.NotificationPolicy> {
    protected readonly internal: alerting.NotificationPolicy;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.NotificationPolicy;
    activeTimeIntervals(activeTimeIntervals: string[]): this;
    continueVal(continueVal: boolean): this;
    groupBy(groupBy: string[]): this;
    groupInterval(groupInterval: string): this;
    groupWait(groupWait: string): this;
    match(match: Record<string, string>): this;
    matchRe(matchRe: alerting.MatchRegexps): this;
    matchers(matchers: alerting.Matchers): this;
    muteTimeIntervals(muteTimeIntervals: string[]): this;
    objectMatchers(objectMatchers: alerting.ObjectMatchers): this;
    provenance(provenance: alerting.Provenance): this;
    receiver(receiver: string): this;
    repeatInterval(repeatInterval: string): this;
    routes(routes: cog.Builder<alerting.NotificationPolicy>[]): this;
}
