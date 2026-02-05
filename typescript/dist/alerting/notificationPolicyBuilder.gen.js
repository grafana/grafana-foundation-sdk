"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.NotificationPolicyBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
class NotificationPolicyBuilder {
    constructor() {
        this.internal = alerting.defaultNotificationPolicy();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    activeTimeIntervals(activeTimeIntervals) {
        this.internal.active_time_intervals = activeTimeIntervals;
        return this;
    }
    continueVal(continueVal) {
        this.internal.continue = continueVal;
        return this;
    }
    groupBy(groupBy) {
        this.internal.group_by = groupBy;
        return this;
    }
    groupInterval(groupInterval) {
        this.internal.group_interval = groupInterval;
        return this;
    }
    groupWait(groupWait) {
        this.internal.group_wait = groupWait;
        return this;
    }
    // Deprecated. Remove before v1.0 release.
    match(match) {
        this.internal.match = match;
        return this;
    }
    matchRe(matchRe) {
        this.internal.match_re = matchRe;
        return this;
    }
    // Matchers is a slice of Matchers that is sortable, implements Stringer, and
    // provides a Matches method to match a LabelSet against all Matchers in the
    // slice. Note that some users of Matchers might require it to be sorted.
    matchers(matchers) {
        this.internal.matchers = matchers;
        return this;
    }
    muteTimeIntervals(muteTimeIntervals) {
        this.internal.mute_time_intervals = muteTimeIntervals;
        return this;
    }
    objectMatchers(objectMatchers) {
        this.internal.object_matchers = objectMatchers;
        return this;
    }
    provenance(provenance) {
        this.internal.provenance = provenance;
        return this;
    }
    receiver(receiver) {
        this.internal.receiver = receiver;
        return this;
    }
    repeatInterval(repeatInterval) {
        this.internal.repeat_interval = repeatInterval;
        return this;
    }
    routes(routes) {
        const routesResources = routes.map(builder1 => builder1.build());
        this.internal.routes = routesResources;
        return this;
    }
}
exports.NotificationPolicyBuilder = NotificationPolicyBuilder;
//# sourceMappingURL=notificationPolicyBuilder.gen.js.map