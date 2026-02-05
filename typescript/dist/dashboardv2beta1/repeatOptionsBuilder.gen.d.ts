import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class RepeatOptionsBuilder implements cog.Builder<dashboardv2beta1.RepeatOptions> {
    protected readonly internal: dashboardv2beta1.RepeatOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RepeatOptions;
    value(value: string): this;
    direction(direction: "h" | "v"): this;
    maxPerRow(maxPerRow: number): this;
}
