import * as common from '../common';
export interface Options {
    displayMode: common.BarGaugeDisplayMode;
    valueMode: common.BarGaugeValueMode;
    namePlacement: common.BarGaugeNamePlacement;
    showUnfilled: boolean;
    sizing: common.BarGaugeSizing;
    minVizWidth: number;
    minVizHeight: number;
    legend: common.VizLegendOptions;
    reduceOptions: common.ReduceDataOptions;
    text?: common.VizTextDisplayOptions;
    maxVizHeight: number;
    orientation: common.VizOrientation;
}
export declare const defaultOptions: () => Options;
