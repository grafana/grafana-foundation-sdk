// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class RuleBuilder implements cog.Builder<alerting.Rule> {
    protected readonly internal: alerting.Rule;

    constructor(title: string) {
        this.internal = alerting.defaultRule();
        if (!(title.length >= 1)) {
            throw new Error("title.length must be >= 1");
        }
        if (!(title.length <= 190)) {
            throw new Error("title.length must be <= 190");
        }
        this.internal.title = title;
    }

    /**
     * Builds the object.
     */
    build(): alerting.Rule {
        return this.internal;
    }

    annotations(annotations: Record<string, string>): this {
        this.internal.annotations = annotations;
        return this;
    }

    condition(condition: string): this {
        this.internal.condition = condition;
        return this;
    }

    queries(data: cog.Builder<alerting.Query>[]): this {
        const dataResources = data.map(builder1 => builder1.build());
        this.internal.data = dataResources;
        return this;
    }

    withQuery(data: cog.Builder<alerting.Query>): this {
        if (!this.internal.data) {
            this.internal.data = [];
        }
        const dataResource = data.build();
        this.internal.data.push(dataResource);
        return this;
    }

    execErrState(execErrState: "OK" | "Alerting" | "Error"): this {
        this.internal.execErrState = execErrState;
        return this;
    }

    folderUID(folderUID: string): this {
        this.internal.folderUID = folderUID;
        return this;
    }

    // The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    // Before this time has elapsed, the rule is only considered to be Pending.
    forDuration(forVal: string): this {
        this.internal.for = forVal;
        return this;
    }

    id(id: number): this {
        this.internal.id = id;
        return this;
    }

    isPaused(isPaused: boolean): this {
        this.internal.isPaused = isPaused;
        return this;
    }

    labels(labels: Record<string, string>): this {
        this.internal.labels = labels;
        return this;
    }

    noDataState(noDataState: "Alerting" | "NoData" | "OK"): this {
        this.internal.noDataState = noDataState;
        return this;
    }

    notificationSettings(notificationSettings: cog.Builder<alerting.NotificationSettings>): this {
        const notificationSettingsResource = notificationSettings.build();
        this.internal.notification_settings = notificationSettingsResource;
        return this;
    }

    orgID(orgID: number): this {
        this.internal.orgID = orgID;
        return this;
    }

    provenance(provenance: alerting.Provenance): this {
        this.internal.provenance = provenance;
        return this;
    }

    ruleGroup(ruleGroup: string): this {
        if (!(ruleGroup.length >= 1)) {
            throw new Error("ruleGroup.length must be >= 1");
        }
        if (!(ruleGroup.length <= 190)) {
            throw new Error("ruleGroup.length must be <= 190");
        }
        this.internal.ruleGroup = ruleGroup;
        return this;
    }

    title(title: string): this {
        if (!(title.length >= 1)) {
            throw new Error("title.length must be >= 1");
        }
        if (!(title.length <= 190)) {
            throw new Error("title.length must be <= 190");
        }
        this.internal.title = title;
        return this;
    }

    uid(uid: string): this {
        if (!(uid.length >= 1)) {
            throw new Error("uid.length must be >= 1");
        }
        if (!(uid.length <= 40)) {
            throw new Error("uid.length must be <= 40");
        }
        this.internal.uid = uid;
        return this;
    }

    updated(updated: string): this {
        this.internal.updated = updated;
        return this;
    }
}
