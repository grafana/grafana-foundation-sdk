"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    showThresholdLabels: false,
    showThresholdMarkers: true,
    sizing: common.BarGaugeSizing.Auto,
    minVizWidth: 75,
    reduceOptions: common.defaultReduceDataOptions(),
    minVizHeight: 75,
    orientation: common.VizOrientation.Auto,
});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map