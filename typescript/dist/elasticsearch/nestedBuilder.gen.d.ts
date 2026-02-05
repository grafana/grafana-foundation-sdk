import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
export declare class NestedBuilder implements cog.Builder<elasticsearch.Nested> {
    protected readonly internal: elasticsearch.Nested;
    constructor();
    /**
     * Builds the object.
     */
    build(): elasticsearch.Nested;
    field(field: string): this;
    id(id: string): this;
    settings(settings: any): this;
}
