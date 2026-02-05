import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class RowRepeatOptionsBuilder implements cog.Builder<dashboardv2beta1.RowRepeatOptions> {
    protected readonly internal: dashboardv2beta1.RowRepeatOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowRepeatOptions;
    value(value: string): this;
}
