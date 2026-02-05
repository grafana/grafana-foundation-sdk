import * as common from '../common';
export declare enum PointShape {
    Circle = "circle",
    Square = "square"
}
export declare const defaultPointShape: () => PointShape;
export declare enum SeriesMapping {
    Auto = "auto",
    Manual = "manual"
}
export declare const defaultSeriesMapping: () => SeriesMapping;
export declare enum XYShowMode {
    Points = "points",
    Lines = "lines",
    PointsAndLines = "points+lines"
}
export declare const defaultXYShowMode: () => XYShowMode;
export interface MatcherConfig {
    id: string;
    options?: any;
}
export declare const defaultMatcherConfig: () => MatcherConfig;
export interface FieldConfig {
    show?: XYShowMode;
    pointSize?: {
        fixed?: number;
        min?: number;
        max?: number;
    };
    pointShape?: PointShape;
    pointStrokeWidth?: number;
    fillOpacity?: number;
    lineWidth?: number;
    hideFrom?: common.HideSeriesConfig;
    axisPlacement?: common.AxisPlacement;
    axisColorMode?: common.AxisColorMode;
    axisLabel?: string;
    axisWidth?: number;
    axisSoftMin?: number;
    axisSoftMax?: number;
    axisGridShow?: boolean;
    scaleDistribution?: common.ScaleDistributionConfig;
    axisCenteredZero?: boolean;
    lineStyle?: common.LineStyle;
    axisBorderShow?: boolean;
}
export declare const defaultFieldConfig: () => FieldConfig;
export interface XYSeriesConfig {
    name?: {
        fixed?: string;
    };
    frame?: {
        matcher: MatcherConfig;
    };
    x?: {
        matcher: MatcherConfig;
    };
    y?: {
        matcher: MatcherConfig;
    };
    color?: {
        matcher: MatcherConfig;
    };
    size?: {
        matcher: MatcherConfig;
    };
}
export declare const defaultXYSeriesConfig: () => XYSeriesConfig;
export interface Options {
    mapping: SeriesMapping;
    legend: common.VizLegendOptions;
    tooltip: common.VizTooltipOptions;
    series: XYSeriesConfig[];
}
export declare const defaultOptions: () => Options;
