import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
export declare class LogGroupBuilder implements cog.Builder<cloudwatch.LogGroup> {
    protected readonly internal: cloudwatch.LogGroup;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.LogGroup;
    arn(arn: string): this;
    name(name: string): this;
    accountId(accountId: string): this;
    accountLabel(accountLabel: string): this;
}
