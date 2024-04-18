"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultFieldConfig = exports.defaultOptions = exports.defaultCandlestickColors = exports.defaultCandlestickFieldMap = exports.defaultColorStrategy = exports.ColorStrategy = exports.defaultCandleStyle = exports.CandleStyle = exports.defaultVizDisplayMode = exports.VizDisplayMode = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
var VizDisplayMode;
(function (VizDisplayMode) {
    VizDisplayMode["CandlesVolume"] = "candles+volume";
    VizDisplayMode["Candles"] = "candles";
    VizDisplayMode["Volume"] = "volume";
})(VizDisplayMode || (exports.VizDisplayMode = VizDisplayMode = {}));
const defaultVizDisplayMode = () => (VizDisplayMode.CandlesVolume);
exports.defaultVizDisplayMode = defaultVizDisplayMode;
var CandleStyle;
(function (CandleStyle) {
    CandleStyle["Candles"] = "candles";
    CandleStyle["OHLCBars"] = "ohlcbars";
})(CandleStyle || (exports.CandleStyle = CandleStyle = {}));
const defaultCandleStyle = () => (CandleStyle.Candles);
exports.defaultCandleStyle = defaultCandleStyle;
var ColorStrategy;
(function (ColorStrategy) {
    ColorStrategy["OpenClose"] = "open-close";
    ColorStrategy["CloseClose"] = "close-close";
})(ColorStrategy || (exports.ColorStrategy = ColorStrategy = {}));
const defaultColorStrategy = () => (ColorStrategy.OpenClose);
exports.defaultColorStrategy = defaultColorStrategy;
const defaultCandlestickFieldMap = () => ({});
exports.defaultCandlestickFieldMap = defaultCandlestickFieldMap;
const defaultCandlestickColors = () => ({
    up: "green",
    down: "red",
    flat: "gray",
});
exports.defaultCandlestickColors = defaultCandlestickColors;
const defaultOptions = () => ({
    mode: VizDisplayMode.CandlesVolume,
    candleStyle: CandleStyle.Candles,
    colorStrategy: ColorStrategy.OpenClose,
    fields: (0, exports.defaultCandlestickFieldMap)(),
    colors: { up: "green", down: "red", flat: "gray", },
    legend: common.defaultVizLegendOptions(),
    tooltip: common.defaultVizTooltipOptions(),
    includeAllFields: false,
});
exports.defaultOptions = defaultOptions;
const defaultFieldConfig = () => (common.defaultGraphFieldConfig());
exports.defaultFieldConfig = defaultFieldConfig;
//# sourceMappingURL=types.gen.js.map