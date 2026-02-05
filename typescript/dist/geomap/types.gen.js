"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultMapCenterID = exports.MapCenterID = exports.defaultTooltipMode = exports.TooltipMode = exports.defaultTooltipOptions = exports.defaultControlsOptions = exports.defaultMapViewConfig = exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    view: (0, exports.defaultMapViewConfig)(),
    controls: (0, exports.defaultControlsOptions)(),
    basemap: common.defaultMapLayerOptions(),
    layers: [],
    tooltip: (0, exports.defaultTooltipOptions)(),
});
exports.defaultOptions = defaultOptions;
const defaultMapViewConfig = () => ({
    id: "zero",
    lat: 0,
    lon: 0,
    zoom: 1,
    allLayers: true,
});
exports.defaultMapViewConfig = defaultMapViewConfig;
const defaultControlsOptions = () => ({});
exports.defaultControlsOptions = defaultControlsOptions;
const defaultTooltipOptions = () => ({
    mode: TooltipMode.None,
});
exports.defaultTooltipOptions = defaultTooltipOptions;
var TooltipMode;
(function (TooltipMode) {
    TooltipMode["None"] = "none";
    TooltipMode["Details"] = "details";
})(TooltipMode || (exports.TooltipMode = TooltipMode = {}));
const defaultTooltipMode = () => (TooltipMode.None);
exports.defaultTooltipMode = defaultTooltipMode;
var MapCenterID;
(function (MapCenterID) {
    MapCenterID["Zero"] = "zero";
    MapCenterID["Coords"] = "coords";
    MapCenterID["Fit"] = "fit";
})(MapCenterID || (exports.MapCenterID = MapCenterID = {}));
const defaultMapCenterID = () => (MapCenterID.Zero);
exports.defaultMapCenterID = defaultMapCenterID;
//# sourceMappingURL=types.gen.js.map