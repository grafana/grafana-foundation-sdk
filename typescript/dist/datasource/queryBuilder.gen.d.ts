import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
    panelId(panelId: number): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    withTransforms(withTransforms: boolean): this;
}
