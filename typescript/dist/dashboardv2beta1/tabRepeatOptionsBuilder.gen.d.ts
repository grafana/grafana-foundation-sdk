import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TabRepeatOptionsBuilder implements cog.Builder<dashboardv2beta1.TabRepeatOptions> {
    protected readonly internal: dashboardv2beta1.TabRepeatOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabRepeatOptions;
    value(value: string): this;
}
