import * as common from '../common';
export declare enum VizDisplayMode {
    CandlesVolume = "candles+volume",
    Candles = "candles",
    Volume = "volume"
}
export declare const defaultVizDisplayMode: () => VizDisplayMode;
export declare enum CandleStyle {
    Candles = "candles",
    OHLCBars = "ohlcbars"
}
export declare const defaultCandleStyle: () => CandleStyle;
export declare enum ColorStrategy {
    OpenClose = "open-close",
    CloseClose = "close-close"
}
export declare const defaultColorStrategy: () => ColorStrategy;
export interface CandlestickFieldMap {
    open?: string;
    high?: string;
    low?: string;
    close?: string;
    volume?: string;
}
export declare const defaultCandlestickFieldMap: () => CandlestickFieldMap;
export interface CandlestickColors {
    up: string;
    down: string;
    flat: string;
}
export declare const defaultCandlestickColors: () => CandlestickColors;
export interface Options {
    mode: VizDisplayMode;
    candleStyle: CandleStyle;
    colorStrategy: ColorStrategy;
    fields: CandlestickFieldMap;
    colors: CandlestickColors;
    legend: common.VizLegendOptions;
    tooltip: common.VizTooltipOptions;
    includeAllFields?: boolean;
}
export declare const defaultOptions: () => Options;
export type FieldConfig = common.GraphFieldConfig;
export declare const defaultFieldConfig: () => FieldConfig;
