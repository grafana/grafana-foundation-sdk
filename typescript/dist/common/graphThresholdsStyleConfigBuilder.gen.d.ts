import * as cog from '../cog';
import * as common from '../common';
export declare class GraphThresholdsStyleConfigBuilder implements cog.Builder<common.GraphThresholdsStyleConfig> {
    protected readonly internal: common.GraphThresholdsStyleConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.GraphThresholdsStyleConfig;
    mode(mode: common.GraphThresholdsStyleMode): this;
}
