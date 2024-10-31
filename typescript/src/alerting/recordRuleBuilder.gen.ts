// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class RecordRuleBuilder implements cog.Builder<alerting.RecordRule> {
    protected readonly internal: alerting.RecordRule;

    constructor() {
        this.internal = alerting.defaultRecordRule();
    }

    build(): alerting.RecordRule {
        return this.internal;
    }

    // Which expression node should be used as the input for the recorded metric.
    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    // Name of the recorded metric.
    metric(metric: string): this {
        this.internal.metric = metric;
        return this;
    }
}
