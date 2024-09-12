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

    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    metric(metric: string): this {
        this.internal.metric = metric;
        return this;
    }
}
