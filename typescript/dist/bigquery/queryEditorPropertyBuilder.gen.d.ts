import * as cog from '../cog';
import * as bigquery from '../bigquery';
export declare class QueryEditorPropertyBuilder implements cog.Builder<bigquery.QueryEditorProperty> {
    protected readonly internal: bigquery.QueryEditorProperty;
    constructor();
    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorProperty;
    name(name: string): this;
}
