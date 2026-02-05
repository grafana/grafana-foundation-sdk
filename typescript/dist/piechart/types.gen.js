"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultFieldConfig = exports.defaultOptions = exports.defaultPieChartLegendOptions = exports.defaultPieChartLegendValues = exports.PieChartLegendValues = exports.defaultPieChartLabels = exports.PieChartLabels = exports.defaultPieChartType = exports.PieChartType = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// Select the pie chart display style.
var PieChartType;
(function (PieChartType) {
    PieChartType["Pie"] = "pie";
    PieChartType["Donut"] = "donut";
})(PieChartType || (exports.PieChartType = PieChartType = {}));
const defaultPieChartType = () => (PieChartType.Pie);
exports.defaultPieChartType = defaultPieChartType;
// Select labels to display on the pie chart.
//  - Name - The series or field name.
//  - Percent - The percentage of the whole.
//  - Value - The raw numerical value.
var PieChartLabels;
(function (PieChartLabels) {
    PieChartLabels["Name"] = "name";
    PieChartLabels["Value"] = "value";
    PieChartLabels["Percent"] = "percent";
})(PieChartLabels || (exports.PieChartLabels = PieChartLabels = {}));
const defaultPieChartLabels = () => (PieChartLabels.Name);
exports.defaultPieChartLabels = defaultPieChartLabels;
// Select values to display in the legend.
//  - Percent: The percentage of the whole.
//  - Value: The raw numerical value.
var PieChartLegendValues;
(function (PieChartLegendValues) {
    PieChartLegendValues["Value"] = "value";
    PieChartLegendValues["Percent"] = "percent";
})(PieChartLegendValues || (exports.PieChartLegendValues = PieChartLegendValues = {}));
const defaultPieChartLegendValues = () => (PieChartLegendValues.Value);
exports.defaultPieChartLegendValues = defaultPieChartLegendValues;
const defaultPieChartLegendOptions = () => ({
    values: [],
    displayMode: common.LegendDisplayMode.List,
    placement: common.LegendPlacement.Bottom,
    showLegend: false,
    calcs: [],
});
exports.defaultPieChartLegendOptions = defaultPieChartLegendOptions;
const defaultOptions = () => ({
    pieType: PieChartType.Pie,
    tooltip: common.defaultVizTooltipOptions(),
    reduceOptions: common.defaultReduceDataOptions(),
    legend: (0, exports.defaultPieChartLegendOptions)(),
    orientation: common.VizOrientation.Auto,
});
exports.defaultOptions = defaultOptions;
const defaultFieldConfig = () => (common.defaultHideableFieldConfig());
exports.defaultFieldConfig = defaultFieldConfig;
//# sourceMappingURL=types.gen.js.map