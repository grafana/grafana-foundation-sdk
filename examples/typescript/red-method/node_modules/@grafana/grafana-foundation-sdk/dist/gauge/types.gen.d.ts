import * as common from '../common';
export interface Options {
    showThresholdLabels: boolean;
    showThresholdMarkers: boolean;
    sizing: common.BarGaugeSizing;
    minVizWidth: number;
    reduceOptions: common.ReduceDataOptions;
    text?: common.VizTextDisplayOptions;
    minVizHeight: number;
    orientation: common.VizOrientation;
}
export declare const defaultOptions: () => Options;
