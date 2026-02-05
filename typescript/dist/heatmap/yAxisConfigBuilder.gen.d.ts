import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';
export declare class YAxisConfigBuilder implements cog.Builder<heatmap.YAxisConfig> {
    protected readonly internal: heatmap.YAxisConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.YAxisConfig;
    unit(unit: string): this;
    reverse(reverse: boolean): this;
    decimals(decimals: number): this;
    min(min: number): this;
    axisPlacement(axisPlacement: common.AxisPlacement): this;
    axisColorMode(axisColorMode: common.AxisColorMode): this;
    axisLabel(axisLabel: string): this;
    axisWidth(axisWidth: number): this;
    axisSoftMin(axisSoftMin: number): this;
    axisSoftMax(axisSoftMax: number): this;
    axisGridShow(axisGridShow: boolean): this;
    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this;
    axisCenteredZero(axisCenteredZero: boolean): this;
    max(max: number): this;
    axisBorderShow(axisBorderShow: boolean): this;
}
