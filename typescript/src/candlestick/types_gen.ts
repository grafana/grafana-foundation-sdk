// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export enum VizDisplayMode {
	CandlesVolume = "candles+volume",
	Candles = "candles",
	Volume = "volume",
}

export const defaultVizDisplayMode = (): VizDisplayMode => (VizDisplayMode.CandlesVolume);

export enum CandleStyle {
	Candles = "candles",
	OHLCBars = "ohlcbars",
}

export const defaultCandleStyle = (): CandleStyle => (CandleStyle.Candles);

export enum ColorStrategy {
	OpenClose = "open-close",
	CloseClose = "close-close",
}

export const defaultColorStrategy = (): ColorStrategy => (ColorStrategy.OpenClose);

export interface CandlestickFieldMap {
	// Corresponds to the starting value of the given period
	open?: string;
	// Corresponds to the highest value of the given period
	high?: string;
	// Corresponds to the lowest value of the given period
	low?: string;
	// Corresponds to the final (end) value of the given period
	close?: string;
	// Corresponds to the sample count in the given period. (e.g. number of trades)
	volume?: string;
}

export const defaultCandlestickFieldMap = (): CandlestickFieldMap => ({
});

export interface CandlestickColors {
	up: string;
	down: string;
	flat: string;
}

export const defaultCandlestickColors = (): CandlestickColors => ({
	up: "green",
	down: "red",
	flat: "gray",
});

export interface Options {
	// Sets which dimensions are used for the visualization
	mode: VizDisplayMode;
	// Sets the style of the candlesticks
	candleStyle: CandleStyle;
	// Sets the color strategy for the candlesticks
	colorStrategy: ColorStrategy;
	// Map fields to appropriate dimension
	fields: CandlestickFieldMap;
	// Set which colors are used when the price movement is up or down
	colors: {
		down: "red";
		up: "green";
		flat: "gray";
	};
	legend: common.VizLegendOptions;
	// When enabled, all fields will be sent to the graph
	includeAllFields?: boolean;
}

export const defaultOptions = (): Options => ({
	mode: VizDisplayMode.CandlesVolume,
	candleStyle: CandleStyle.Candles,
	colorStrategy: ColorStrategy.OpenClose,
	fields: defaultCandlestickFieldMap(),
	colors: { down: "red", up: "green", flat: "gray", },
	legend: common.defaultVizLegendOptions(),
	includeAllFields: false,
});

export type FieldConfig = common.GraphFieldConfig;

export const defaultFieldConfig = (): FieldConfig => (common.defaultGraphFieldConfig());

