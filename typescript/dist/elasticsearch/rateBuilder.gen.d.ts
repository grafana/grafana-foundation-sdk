import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class RateBuilder implements cog.Builder<elasticsearch.Rate> {
    protected readonly internal: elasticsearch.Rate;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Rate;
    field(field: string): this;
    id(id: string): this;
    settings(settings: {
        unit?: string;
        mode?: string;
    }): this;
    hide(hide: boolean): this;
}
