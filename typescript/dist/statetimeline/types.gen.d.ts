import * as common from '../common';
export interface Options {
    showValue: common.VisibilityMode;
    rowHeight: number;
    mergeValues?: boolean;
    alignValue?: common.TimelineValueAlignment;
    legend: common.VizLegendOptions;
    tooltip: common.VizTooltipOptions;
    timezone?: common.TimeZone[];
    perPage?: number;
}
export declare const defaultOptions: () => Options;
export interface FieldConfig {
    lineWidth?: number;
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
    fillOpacity?: number;
    axisBorderShow?: boolean;
}
export declare const defaultFieldConfig: () => FieldConfig;
