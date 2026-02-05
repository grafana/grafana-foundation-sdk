import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ValueMappingResultBuilder implements cog.Builder<dashboardv2beta1.ValueMappingResult> {
    protected readonly internal: dashboardv2beta1.ValueMappingResult;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ValueMappingResult;
    text(text: string): this;
    color(color: string): this;
    icon(icon: string): this;
    index(index: number): this;
}
