"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = exports.defaultZoomMode = exports.ZoomMode = exports.defaultEdgeOptions = exports.defaultNodeOptions = exports.defaultArcOption = void 0;
const defaultArcOption = () => ({});
exports.defaultArcOption = defaultArcOption;
const defaultNodeOptions = () => ({});
exports.defaultNodeOptions = defaultNodeOptions;
const defaultEdgeOptions = () => ({});
exports.defaultEdgeOptions = defaultEdgeOptions;
var ZoomMode;
(function (ZoomMode) {
    ZoomMode["Cooperative"] = "cooperative";
    ZoomMode["Greedy"] = "greedy";
})(ZoomMode || (exports.ZoomMode = ZoomMode = {}));
const defaultZoomMode = () => (ZoomMode.Cooperative);
exports.defaultZoomMode = defaultZoomMode;
const defaultOptions = () => ({});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map