import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class DerivativeBuilder implements cog.Builder<elasticsearch.Derivative> {
    protected readonly internal: elasticsearch.Derivative;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Derivative;
    pipelineAgg(pipelineAgg: string): this;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        unit?: string;
    }): this;
    hide(hide: boolean): this;
}
