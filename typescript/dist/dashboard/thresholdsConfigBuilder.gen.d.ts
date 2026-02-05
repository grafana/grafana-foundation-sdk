import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class ThresholdsConfigBuilder implements cog.Builder<dashboard.ThresholdsConfig> {
    protected readonly internal: dashboard.ThresholdsConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.ThresholdsConfig;
    mode(mode: dashboard.ThresholdsMode): this;
    steps(steps: dashboard.Threshold[]): this;
}
