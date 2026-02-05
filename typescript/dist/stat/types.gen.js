"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    graphMode: common.BigValueGraphMode.Area,
    colorMode: common.BigValueColorMode.Value,
    justifyMode: common.BigValueJustifyMode.Auto,
    textMode: common.BigValueTextMode.Auto,
    wideLayout: true,
    showPercentChange: false,
    reduceOptions: common.defaultReduceDataOptions(),
    percentChangeColorMode: common.PercentChangeColorMode.Standard,
    orientation: common.VizOrientation.Auto,
});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map