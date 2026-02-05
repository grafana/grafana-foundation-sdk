"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = exports.defaultXYSeriesConfig = exports.defaultFieldConfig = exports.defaultMatcherConfig = exports.defaultXYShowMode = exports.XYShowMode = exports.defaultSeriesMapping = exports.SeriesMapping = exports.defaultPointShape = exports.PointShape = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
var PointShape;
(function (PointShape) {
    PointShape["Circle"] = "circle";
    PointShape["Square"] = "square";
})(PointShape || (exports.PointShape = PointShape = {}));
const defaultPointShape = () => (PointShape.Circle);
exports.defaultPointShape = defaultPointShape;
var SeriesMapping;
(function (SeriesMapping) {
    SeriesMapping["Auto"] = "auto";
    SeriesMapping["Manual"] = "manual";
})(SeriesMapping || (exports.SeriesMapping = SeriesMapping = {}));
const defaultSeriesMapping = () => (SeriesMapping.Auto);
exports.defaultSeriesMapping = defaultSeriesMapping;
var XYShowMode;
(function (XYShowMode) {
    XYShowMode["Points"] = "points";
    XYShowMode["Lines"] = "lines";
    XYShowMode["PointsAndLines"] = "points+lines";
})(XYShowMode || (exports.XYShowMode = XYShowMode = {}));
const defaultXYShowMode = () => (XYShowMode.Points);
exports.defaultXYShowMode = defaultXYShowMode;
const defaultMatcherConfig = () => ({
    id: "",
});
exports.defaultMatcherConfig = defaultMatcherConfig;
const defaultFieldConfig = () => ({
    show: XYShowMode.Points,
    fillOpacity: 50,
});
exports.defaultFieldConfig = defaultFieldConfig;
const defaultXYSeriesConfig = () => ({});
exports.defaultXYSeriesConfig = defaultXYSeriesConfig;
const defaultOptions = () => ({
    mapping: SeriesMapping.Auto,
    legend: common.defaultVizLegendOptions(),
    tooltip: common.defaultVizTooltipOptions(),
    series: [],
});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map