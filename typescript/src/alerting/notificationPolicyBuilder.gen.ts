// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
export class NotificationPolicyBuilder implements cog.Builder<alerting.NotificationPolicy> {
    protected readonly internal: alerting.NotificationPolicy;

    constructor() {
        this.internal = alerting.defaultNotificationPolicy();
    }

    build(): alerting.NotificationPolicy {
        return this.internal;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    continueVal(continueVal: boolean): this {
        this.internal.continue = continueVal;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    groupBy(groupBy: string[]): this {
        this.internal.group_by = groupBy;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    groupInterval(groupInterval: string): this {
        this.internal.group_interval = groupInterval;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    groupWait(groupWait: string): this {
        this.internal.group_wait = groupWait;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    match(match: Record<string, string>): this {
        this.internal.match = match;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    matchRe(matchRe: alerting.MatchRegexps): this {
        this.internal.match_re = matchRe;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    matchers(matchers: alerting.Matchers): this {
        this.internal.matchers = matchers;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    muteTimeIntervals(muteTimeIntervals: string[]): this {
        this.internal.mute_time_intervals = muteTimeIntervals;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    objectMatchers(objectMatchers: alerting.ObjectMatchers): this {
        this.internal.object_matchers = objectMatchers;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    provenance(provenance: alerting.Provenance): this {
        this.internal.provenance = provenance;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    receiver(receiver: string): this {
        this.internal.receiver = receiver;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    repeatInterval(repeatInterval: string): this {
        this.internal.repeat_interval = repeatInterval;
        return this;
    }

    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    routes(routes: cog.Builder<alerting.NotificationPolicy>[]): this {
        const routesResources = routes.map(builder1 => builder1.build());
        this.internal.routes = routesResources;
        return this;
    }
}
