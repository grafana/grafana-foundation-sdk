import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class CountBuilder implements cog.Builder<elasticsearch.Count> {
    protected readonly internal: elasticsearch.Count;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Count;
    id(id: string): this;
    hide(hide: boolean): this;
}
