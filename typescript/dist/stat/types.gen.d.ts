import * as common from '../common';
export interface Options {
    graphMode: common.BigValueGraphMode;
    colorMode: common.BigValueColorMode;
    justifyMode: common.BigValueJustifyMode;
    textMode: common.BigValueTextMode;
    wideLayout: boolean;
    showPercentChange: boolean;
    reduceOptions: common.ReduceDataOptions;
    text?: common.VizTextDisplayOptions;
    percentChangeColorMode: common.PercentChangeColorMode;
    orientation: common.VizOrientation;
}
export declare const defaultOptions: () => Options;
