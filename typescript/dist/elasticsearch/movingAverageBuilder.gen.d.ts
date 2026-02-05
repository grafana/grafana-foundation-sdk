import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class MovingAverageBuilder implements cog.Builder<elasticsearch.MovingAverage> {
    protected readonly internal: elasticsearch.MovingAverage;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingAverage;
    pipelineAgg(pipelineAgg: string): this;
    field(field: string): this;
    id(id: string): this;
    settings(settings: Record<string, any>): this;
    hide(hide: boolean): this;
}
