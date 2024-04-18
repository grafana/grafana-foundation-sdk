"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultFieldConfig = exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    bucketCount: 30,
    bucketOffset: 0,
    legend: common.defaultVizLegendOptions(),
    tooltip: common.defaultVizTooltipOptions(),
});
exports.defaultOptions = defaultOptions;
const defaultFieldConfig = () => ({
    lineWidth: 1,
    fillOpacity: 80,
    gradientMode: common.GraphGradientMode.None,
});
exports.defaultFieldConfig = defaultFieldConfig;
//# sourceMappingURL=types.gen.js.map