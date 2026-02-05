import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class AverageBuilder implements cog.Builder<elasticsearch.Average> {
    protected readonly internal: elasticsearch.Average;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Average;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        script?: elasticsearch.InlineScript;
        missing?: string;
    }): this;
    hide(hide: boolean): this;
}
