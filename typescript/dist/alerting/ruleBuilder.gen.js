"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RuleBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class RuleBuilder {
    constructor(title) {
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
    build() {
        return this.internal;
    }
    annotations(annotations) {
        this.internal.annotations = annotations;
        return this;
    }
    condition(condition) {
        this.internal.condition = condition;
        return this;
    }
    queries(data) {
        const dataResources = data.map(builder1 => builder1.build());
        this.internal.data = dataResources;
        return this;
    }
    withQuery(data) {
        if (!this.internal.data) {
            this.internal.data = [];
        }
        const dataResource = data.build();
        this.internal.data.push(dataResource);
        return this;
    }
    execErrState(execErrState) {
        this.internal.execErrState = execErrState;
        return this;
    }
    folderUID(folderUID) {
        this.internal.folderUID = folderUID;
        return this;
    }
    // The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    // Before this time has elapsed, the rule is only considered to be Pending.
    forDuration(forVal) {
        this.internal.for = forVal;
        return this;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    isPaused(isPaused) {
        this.internal.isPaused = isPaused;
        return this;
    }
    labels(labels) {
        this.internal.labels = labels;
        return this;
    }
    noDataState(noDataState) {
        this.internal.noDataState = noDataState;
        return this;
    }
    notificationSettings(notificationSettings) {
        const notificationSettingsResource = notificationSettings.build();
        this.internal.notification_settings = notificationSettingsResource;
        return this;
    }
    orgID(orgID) {
        this.internal.orgID = orgID;
        return this;
    }
    provenance(provenance) {
        this.internal.provenance = provenance;
        return this;
    }
    record(record) {
        const recordResource = record.build();
        this.internal.record = recordResource;
        return this;
    }
    ruleGroup(ruleGroup) {
        if (!(ruleGroup.length >= 1)) {
            throw new Error("ruleGroup.length must be >= 1");
        }
        if (!(ruleGroup.length <= 190)) {
            throw new Error("ruleGroup.length must be <= 190");
        }
        this.internal.ruleGroup = ruleGroup;
        return this;
    }
    title(title) {
        if (!(title.length >= 1)) {
            throw new Error("title.length must be >= 1");
        }
        if (!(title.length <= 190)) {
            throw new Error("title.length must be <= 190");
        }
        this.internal.title = title;
        return this;
    }
    uid(uid) {
        if (!(uid.length >= 1)) {
            throw new Error("uid.length must be >= 1");
        }
        if (!(uid.length <= 40)) {
            throw new Error("uid.length must be <= 40");
        }
        this.internal.uid = uid;
        return this;
    }
    updated(updated) {
        this.internal.updated = updated;
        return this;
    }
}
exports.RuleBuilder = RuleBuilder;
//# sourceMappingURL=ruleBuilder.gen.js.map