import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class DataQueryKindBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    group(group: string): this;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
    spec(spec: any): this;
}
