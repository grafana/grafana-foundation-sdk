import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConditionalRenderingGroupBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingGroupKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingGroupKind;
    spec(spec: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupSpec>): this;
    visibility(visibility: "show" | "hide"): this;
    condition(condition: "and" | "or"): this;
    items(items: (cog.Builder<dashboardv2beta1.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind>)[]): this;
}
