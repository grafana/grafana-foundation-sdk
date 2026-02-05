import * as cog from '../cog';
import * as alerting from '../alerting';
export declare class RuleBuilder implements cog.Builder<alerting.Rule> {
    protected readonly internal: alerting.Rule;
    constructor(title: string);
    /**
     * Builds the object.
     */
    build(): alerting.Rule;
    annotations(annotations: Record<string, string>): this;
    condition(condition: string): this;
    queries(data: cog.Builder<alerting.Query>[]): this;
    withQuery(data: cog.Builder<alerting.Query>): this;
    execErrState(execErrState: "OK" | "Alerting" | "Error"): this;
    folderUID(folderUID: string): this;
    forDuration(forVal: string): this;
    id(id: number): this;
    isPaused(isPaused: boolean): this;
    labels(labels: Record<string, string>): this;
    noDataState(noDataState: "Alerting" | "NoData" | "OK"): this;
    notificationSettings(notificationSettings: cog.Builder<alerting.NotificationSettings>): this;
    orgID(orgID: number): this;
    provenance(provenance: alerting.Provenance): this;
    record(record: cog.Builder<alerting.RecordRule>): this;
    ruleGroup(ruleGroup: string): this;
    title(title: string): this;
    uid(uid: string): this;
    updated(updated: string): this;
}
