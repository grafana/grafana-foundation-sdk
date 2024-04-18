"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultFieldConfig = exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    orientation: common.VizOrientation.Auto,
    barRadius: 0,
    xTickLabelRotation: 0,
    xTickLabelMaxLength: 0,
    xTickLabelSpacing: 0,
    stacking: common.StackingMode.None,
    showValue: common.VisibilityMode.Auto,
    barWidth: 0.97,
    groupWidth: 0.7,
    legend: common.defaultVizLegendOptions(),
    tooltip: common.defaultVizTooltipOptions(),
    fullHighlight: false,
});
exports.defaultOptions = defaultOptions;
const defaultFieldConfig = () => ({
    lineWidth: 1,
    fillOpacity: 80,
    gradientMode: common.GraphGradientMode.None,
});
exports.defaultFieldConfig = defaultFieldConfig;
//# sourceMappingURL=types.gen.js.map