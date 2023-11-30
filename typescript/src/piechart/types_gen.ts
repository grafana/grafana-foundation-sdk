// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


// Select the pie chart display style.
export enum PieChartType {
	Pie = "pie",
	Donut = "donut",
}

export const defaultPieChartType = (): PieChartType => (PieChartType.Pie);

// Select labels to display on the pie chart.
//  - Name - The series or field name.
//  - Percent - The percentage of the whole.
//  - Value - The raw numerical value.
export enum PieChartLabels {
	Name = "name",
	Value = "value",
	Percent = "percent",
}

export const defaultPieChartLabels = (): PieChartLabels => (PieChartLabels.Name);

// Select values to display in the legend.
//  - Percent: The percentage of the whole.
//  - Value: The raw numerical value.
export enum PieChartLegendValues {
	Value = "value",
	Percent = "percent",
}

export const defaultPieChartLegendValues = (): PieChartLegendValues => (PieChartLegendValues.Value);

export interface PieChartLegendOptions {
	values: PieChartLegendValues[];
	displayMode: common.LegendDisplayMode;
	placement: common.LegendPlacement;
	showLegend: boolean;
	asTable?: boolean;
	isVisible?: boolean;
	sortBy?: string;
	sortDesc?: boolean;
	width?: number;
	calcs: string[];
}

export const defaultPieChartLegendOptions = (): PieChartLegendOptions => ({
	values: [],
	displayMode: common.LegendDisplayMode.List,
	placement: common.LegendPlacement.Bottom,
	showLegend: false,
	calcs: [],
});

export interface Options {
	pieType: PieChartType;
	displayLabels: PieChartLabels[];
	tooltip: common.VizTooltipOptions;
	reduceOptions: common.ReduceDataOptions;
	text?: common.VizTextDisplayOptions;
	legend: PieChartLegendOptions;
	orientation: common.VizOrientation;
}

export const defaultOptions = (): Options => ({
	pieType: PieChartType.Pie,
	displayLabels: [],
	tooltip: common.defaultVizTooltipOptions(),
	reduceOptions: common.defaultReduceDataOptions(),
	legend: defaultPieChartLegendOptions(),
	orientation: common.VizOrientation.Auto,
});

export type FieldConfig = common.HideableFieldConfig;

export const defaultFieldConfig = (): FieldConfig => (common.defaultHideableFieldConfig());

