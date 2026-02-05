import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class SumBuilder implements cog.Builder<elasticsearch.Sum> {
    protected readonly internal: elasticsearch.Sum;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Sum;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        script?: elasticsearch.InlineScript;
        missing?: string;
    }): this;
    hide(hide: boolean): this;
}
