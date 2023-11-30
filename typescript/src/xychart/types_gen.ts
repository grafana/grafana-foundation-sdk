// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


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
	name?: string;
	labelValue?: common.TextDimensionConfig;
	axisBorderShow?: boolean;
}

export const defaultScatterSeriesConfig = (): ScatterSeriesConfig => ({
	show: ScatterShow.Points,
	label: common.VisibilityMode.Auto,
});

export interface Options {
	seriesMapping?: SeriesMapping;
	dims: XYDimensionConfig;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	series: ScatterSeriesConfig[];
}

export const defaultOptions = (): Options => ({
	dims: defaultXYDimensionConfig(),
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
	series: [],
});

