"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultFieldConfig = exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    showValue: common.VisibilityMode.Auto,
    rowHeight: 0.9,
    mergeValues: true,
    alignValue: common.TimelineValueAlignment.Left,
    legend: common.defaultVizLegendOptions(),
    tooltip: common.defaultVizTooltipOptions(),
    perPage: 20,
});
exports.defaultOptions = defaultOptions;
const defaultFieldConfig = () => ({
    lineWidth: 0,
    fillOpacity: 70,
});
exports.defaultFieldConfig = defaultFieldConfig;
//# sourceMappingURL=types.gen.js.map