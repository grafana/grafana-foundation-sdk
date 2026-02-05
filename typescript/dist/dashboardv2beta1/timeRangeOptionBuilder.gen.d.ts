import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TimeRangeOptionBuilder implements cog.Builder<dashboardv2beta1.TimeRangeOption> {
    protected readonly internal: dashboardv2beta1.TimeRangeOption;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TimeRangeOption;
    display(display: string): this;
    from(from: string): this;
    to(to: string): this;
}
