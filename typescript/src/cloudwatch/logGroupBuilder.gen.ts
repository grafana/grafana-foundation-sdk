// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class LogGroupBuilder implements cog.Builder<cloudwatch.LogGroup> {
    protected readonly internal: cloudwatch.LogGroup;

    constructor() {
        this.internal = cloudwatch.defaultLogGroup();
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.LogGroup {
        return this.internal;
    }

    // ARN of the log group
    arn(arn: string): this {
        this.internal.arn = arn;
        return this;
    }

    // Name of the log group
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // AccountId of the log group
    accountId(accountId: string): this {
        this.internal.accountId = accountId;
        return this;
    }

    // Label of the log group
    accountLabel(accountLabel: string): this {
        this.internal.accountLabel = accountLabel;
        return this;
    }
}
