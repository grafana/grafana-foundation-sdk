import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class CumulativeSumBuilder implements cog.Builder<elasticsearch.CumulativeSum> {
    protected readonly internal: elasticsearch.CumulativeSum;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.CumulativeSum;
    pipelineAgg(pipelineAgg: string): this;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        format?: string;
    }): this;
    hide(hide: boolean): this;
}
