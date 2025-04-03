// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export enum PointShape {
	Circle = "circle",
	Square = "square",
}

export const defaultPointShape = (): PointShape => (PointShape.Circle);

export enum SeriesMapping {
	Auto = "auto",
	Manual = "manual",
}

export const defaultSeriesMapping = (): SeriesMapping => (SeriesMapping.Auto);

export enum XYShowMode {
	Points = "points",
	Lines = "lines",
	PointsAndLines = "points+lines",
}

export const defaultXYShowMode = (): XYShowMode => (XYShowMode.Points);

// NOTE: (copied from dashboard_kind.cue, since not exported)
// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
export interface MatcherConfig {
	// The matcher id. This is used to find the matcher implementation from registry.
	id: string;
	// The matcher options. This is specific to the matcher implementation.
	options?: any;
}

export const defaultMatcherConfig = (): MatcherConfig => ({
	id: "",
});

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

export const defaultFieldConfig = (): FieldConfig => ({
	show: XYShowMode.Points,
	fillOpacity: 50,
});

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

export const defaultXYSeriesConfig = (): XYSeriesConfig => ({
});

export interface Options {
	mapping: SeriesMapping;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	series: XYSeriesConfig[];
}

export const defaultOptions = (): Options => ({
	mapping: SeriesMapping.Auto,
	legend: common.defaultVizLegendOptions(),
	tooltip: common.defaultVizTooltipOptions(),
	series: [],
});

