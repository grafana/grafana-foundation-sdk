import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConditionalRenderingDataSpecBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingDataSpec> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingDataSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingDataSpec;
    value(value: boolean): this;
}
