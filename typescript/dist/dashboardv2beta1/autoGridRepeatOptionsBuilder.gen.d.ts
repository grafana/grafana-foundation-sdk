import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AutoGridRepeatOptionsBuilder implements cog.Builder<dashboardv2beta1.AutoGridRepeatOptions> {
    protected readonly internal: dashboardv2beta1.AutoGridRepeatOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridRepeatOptions;
    value(value: string): this;
}
