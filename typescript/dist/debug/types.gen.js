"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = exports.defaultDebugMode = exports.DebugMode = exports.defaultUpdateConfig = void 0;
const defaultUpdateConfig = () => ({
    render: false,
    dataChanged: false,
    schemaChanged: false,
});
exports.defaultUpdateConfig = defaultUpdateConfig;
var DebugMode;
(function (DebugMode) {
    DebugMode["Render"] = "render";
    DebugMode["Events"] = "events";
    DebugMode["Cursor"] = "cursor";
    DebugMode["State"] = "State";
    DebugMode["ThrowError"] = "ThrowError";
})(DebugMode || (exports.DebugMode = DebugMode = {}));
const defaultDebugMode = () => (DebugMode.Render);
exports.defaultDebugMode = defaultDebugMode;
const defaultOptions = () => ({
    mode: DebugMode.Render,
});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map