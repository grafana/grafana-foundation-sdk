// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


// Auto is "table" in the UI
export enum SeriesMapping {
	Auto = "auto",
	Manual = "manual",
}

export const defaultSeriesMapping = (): SeriesMapping => (SeriesMapping.Auto);

export enum ScatterShow {
	Points = "points",
	Lines = "lines",
	PointsAndLines = "points+lines",
}

export const defaultScatterShow = (): ScatterShow => (ScatterShow.Points);

// Configuration for the Table/Auto mode
export interface XYDimensionConfig {
	frame: number;
	x?: string;
	exclude?: string[];
}

export const defaultXYDimensionConfig = (): XYDimensionConfig => ({
	frame: 0,
});

export interface FieldConfig {
	show?: ScatterShow;
	pointSize?: common.ScaleDimensionConfig;
	pointColor?: common.ColorDimensionConfig;
	lineColor?: common.ColorDimensionConfig;
	lineWidth?: number;
	lineStyle?: common.LineStyle;
	label?: common.VisibilityMode;
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
	labelValue?: common.TextDimensionConfig;
	axisBorderShow?: boolean;
}

export const defaultFieldConfig = (): FieldConfig => ({
	show: ScatterShow.Points,
	label: common.VisibilityMode.Auto,
});

export interface ScatterSeriesConfig {
	x?: string;
	y?: string;
	name?: string;
	show?: ScatterShow;
	pointSize?: common.ScaleDimensionConfig;
	pointColor?: common.ColorDimensionConfig;
	lineColor?: common.ColorDimensionConfig;
	lineWidth?: number;
	lineStyle?: common.LineStyle;
	label?: common.VisibilityMode;
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
	frame?: number;
	labelValue?: common.TextDimensionConfig;
	axisBorderShow?: boolean;
}

export const defaultScatterSeriesConfig = (): ScatterSeriesConfig => ({
	show: ScatterShow.Points,
	label: common.VisibilityMode.Auto,
});

export interface Options {
	seriesMapping?: SeriesMapping;
	// Table Mode (auto)
	dims: XYDimensionConfig;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	// Manual Mode
	series: ScatterSeriesConfig[];
}

export const defaultOptions = (): Options => ({
	dims: defaultXYDimensionConfig(),
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
	series: [],
});

