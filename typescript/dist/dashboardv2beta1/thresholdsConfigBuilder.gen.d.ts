import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ThresholdsConfigBuilder implements cog.Builder<dashboardv2beta1.ThresholdsConfig> {
    protected readonly internal: dashboardv2beta1.ThresholdsConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ThresholdsConfig;
    mode(mode: dashboardv2beta1.ThresholdsMode): this;
    steps(steps: dashboardv2beta1.Threshold[]): this;
}
