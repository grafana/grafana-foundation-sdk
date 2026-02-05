import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConditionalRenderingTimeRangeSizeBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind;
    spec(spec: cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec>): this;
    value(value: string): this;
}
