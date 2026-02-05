import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConditionalRenderingDataBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingDataKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingDataKind;
    spec(spec: cog.Builder<dashboardv2beta1.ConditionalRenderingDataSpec>): this;
    value(value: boolean): this;
}
