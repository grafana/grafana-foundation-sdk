import * as cog from '../cog';
import * as common from '../common';
export declare class AxisConfigBuilder implements cog.Builder<common.AxisConfig> {
    protected readonly internal: common.AxisConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.AxisConfig;
    axisPlacement(axisPlacement: common.AxisPlacement): this;
    axisColorMode(axisColorMode: common.AxisColorMode): this;
    axisLabel(axisLabel: string): this;
    axisWidth(axisWidth: number): this;
    axisSoftMin(axisSoftMin: number): this;
    axisSoftMax(axisSoftMax: number): this;
    axisGridShow(axisGridShow: boolean): this;
    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this;
    axisCenteredZero(axisCenteredZero: boolean): this;
    axisBorderShow(axisBorderShow: boolean): this;
}
