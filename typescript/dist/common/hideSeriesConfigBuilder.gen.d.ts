import * as cog from '../cog';
import * as common from '../common';
export declare class HideSeriesConfigBuilder implements cog.Builder<common.HideSeriesConfig> {
    protected readonly internal: common.HideSeriesConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.HideSeriesConfig;
    tooltip(tooltip: boolean): this;
    legend(legend: boolean): this;
    viz(viz: boolean): this;
}
