import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class ContactPointBuilder implements cog.Builder<alerting.ContactPoint> {
    protected readonly internal: alerting.ContactPoint;
    constructor();
    /**
     * Builds the object.
     */
    build(): alerting.ContactPoint;
    disableResolveMessage(disableResolveMessage: boolean): this;
    name(name: string): this;
    provenance(provenance: string): this;
    settings(settings: alerting.Json): this;
    type(type: "alertmanager" | "dingding" | "discord" | "email" | "googlechat" | "kafka" | "line" | "opsgenie" | "pagerduty" | "pushover" | "sensugo" | "slack" | "teams" | "telegram" | "threema" | "victorops" | "webhook" | "wecom"): this;
    uid(uid: string): this;
}
