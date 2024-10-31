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

    continueVal(continueVal: boolean): this {
        this.internal.continue = continueVal;
        return this;
    }

    groupBy(groupBy: string[]): this {
        this.internal.group_by = groupBy;
        return this;
    }

    groupInterval(groupInterval: string): this {
        this.internal.group_interval = groupInterval;
        return this;
    }

    groupWait(groupWait: string): this {
        this.internal.group_wait = groupWait;
        return this;
    }

    // Deprecated. Remove before v1.0 release.
    match(match: Record<string, string>): this {
        this.internal.match = match;
        return this;
    }

    matchRe(matchRe: alerting.MatchRegexps): this {
        this.internal.match_re = matchRe;
        return this;
    }

    // Matchers is a slice of Matchers that is sortable, implements Stringer, and
    // provides a Matches method to match a LabelSet against all Matchers in the
    // slice. Note that some users of Matchers might require it to be sorted.
    matchers(matchers: alerting.Matchers): this {
        this.internal.matchers = matchers;
        return this;
    }

    muteTimeIntervals(muteTimeIntervals: string[]): this {
        this.internal.mute_time_intervals = muteTimeIntervals;
        return this;
    }

    // Matchers is a slice of Matchers that is sortable, implements Stringer, and
    // provides a Matches method to match a LabelSet against all Matchers in the
    // slice. Note that some users of Matchers might require it to be sorted.
    objectMatchers(objectMatchers: alerting.ObjectMatchers): this {
        this.internal.object_matchers = objectMatchers;
        return this;
    }

    provenance(provenance: alerting.Provenance): this {
        this.internal.provenance = provenance;
        return this;
    }

    receiver(receiver: string): this {
        this.internal.receiver = receiver;
        return this;
    }

    repeatInterval(repeatInterval: string): this {
        this.internal.repeat_interval = repeatInterval;
        return this;
    }

    routes(routes: cog.Builder<alerting.NotificationPolicy>[]): this {
        const routesResources = routes.map(builder1 => builder1.build());
        this.internal.routes = routesResources;
        return this;
    }
}
