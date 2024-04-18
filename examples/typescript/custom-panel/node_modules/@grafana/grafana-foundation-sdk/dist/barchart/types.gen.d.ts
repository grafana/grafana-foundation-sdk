import * as common from '../common';
export interface Options {
    xField?: string;
    colorByField?: string;
    orientation: common.VizOrientation;
    barRadius?: number;
    xTickLabelRotation: number;
    xTickLabelMaxLength: number;
    xTickLabelSpacing?: number;
    stacking: common.StackingMode;
    showValue: common.VisibilityMode;
    barWidth: number;
    groupWidth: number;
    legend: common.VizLegendOptions;
    tooltip: common.VizTooltipOptions;
    text?: common.VizTextDisplayOptions;
    fullHighlight: boolean;
}
export declare const defaultOptions: () => Options;
export interface FieldConfig {
    lineWidth?: number;
    fillOpacity?: number;
    gradientMode?: common.GraphGradientMode;
    axisPlacement?: common.AxisPlacement;
    axisColorMode?: common.AxisColorMode;
    axisLabel?: string;
    axisWidth?: number;
    axisSoftMin?: number;
    axisSoftMax?: number;
    axisGridShow?: boolean;
    scaleDistribution?: common.ScaleDistributionConfig;
    axisCenteredZero?: boolean;
    hideFrom?: common.HideSeriesConfig;
    thresholdsStyle?: common.GraphThresholdsStyleConfig;
    axisBorderShow?: boolean;
}
export declare const defaultFieldConfig: () => FieldConfig;
