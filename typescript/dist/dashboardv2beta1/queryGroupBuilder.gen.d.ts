import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class QueryGroupBuilder implements cog.Builder<dashboardv2beta1.QueryGroupKind> {
    protected readonly internal: dashboardv2beta1.QueryGroupKind;
    private nextQueryId;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.QueryGroupKind;
    targets(queries: cog.Builder<dashboardv2beta1.PanelQueryKind>[]): this;
    target(querie: cog.Builder<dashboardv2beta1.PanelQueryKind>): this;
    transformations(transformations: cog.Builder<dashboardv2beta1.TransformationKind>[]): this;
    transformation(transformation: cog.Builder<dashboardv2beta1.TransformationKind>): this;
    queryOptions(queryOptions: cog.Builder<dashboardv2beta1.QueryOptionsSpec>): this;
}
