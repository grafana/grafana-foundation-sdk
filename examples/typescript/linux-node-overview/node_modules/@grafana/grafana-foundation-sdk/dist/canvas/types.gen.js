"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = exports.defaultCanvasElementOptions = exports.defaultCanvasConnection = exports.defaultConnectionPath = exports.ConnectionPath = exports.defaultConnectionCoordinates = exports.defaultHttpRequestMethod = exports.HttpRequestMethod = exports.defaultLineConfig = exports.defaultBackgroundConfig = exports.defaultBackgroundImageSize = exports.BackgroundImageSize = exports.defaultPlacement = exports.defaultConstraint = exports.defaultVerticalConstraint = exports.VerticalConstraint = exports.defaultHorizontalConstraint = exports.HorizontalConstraint = void 0;
var HorizontalConstraint;
(function (HorizontalConstraint) {
    HorizontalConstraint["Left"] = "left";
    HorizontalConstraint["Right"] = "right";
    HorizontalConstraint["LeftRight"] = "leftright";
    HorizontalConstraint["Center"] = "center";
    HorizontalConstraint["Scale"] = "scale";
})(HorizontalConstraint || (exports.HorizontalConstraint = HorizontalConstraint = {}));
const defaultHorizontalConstraint = () => (HorizontalConstraint.Left);
exports.defaultHorizontalConstraint = defaultHorizontalConstraint;
var VerticalConstraint;
(function (VerticalConstraint) {
    VerticalConstraint["Top"] = "top";
    VerticalConstraint["Bottom"] = "bottom";
    VerticalConstraint["TopBottom"] = "topbottom";
    VerticalConstraint["Center"] = "center";
    VerticalConstraint["Scale"] = "scale";
})(VerticalConstraint || (exports.VerticalConstraint = VerticalConstraint = {}));
const defaultVerticalConstraint = () => (VerticalConstraint.Top);
exports.defaultVerticalConstraint = defaultVerticalConstraint;
const defaultConstraint = () => ({});
exports.defaultConstraint = defaultConstraint;
const defaultPlacement = () => ({});
exports.defaultPlacement = defaultPlacement;
var BackgroundImageSize;
(function (BackgroundImageSize) {
    BackgroundImageSize["Original"] = "original";
    BackgroundImageSize["Contain"] = "contain";
    BackgroundImageSize["Cover"] = "cover";
    BackgroundImageSize["Fill"] = "fill";
    BackgroundImageSize["Tile"] = "tile";
})(BackgroundImageSize || (exports.BackgroundImageSize = BackgroundImageSize = {}));
const defaultBackgroundImageSize = () => (BackgroundImageSize.Original);
exports.defaultBackgroundImageSize = defaultBackgroundImageSize;
const defaultBackgroundConfig = () => ({});
exports.defaultBackgroundConfig = defaultBackgroundConfig;
const defaultLineConfig = () => ({});
exports.defaultLineConfig = defaultLineConfig;
var HttpRequestMethod;
(function (HttpRequestMethod) {
    HttpRequestMethod["GET"] = "GET";
    HttpRequestMethod["POST"] = "POST";
    HttpRequestMethod["PUT"] = "PUT";
})(HttpRequestMethod || (exports.HttpRequestMethod = HttpRequestMethod = {}));
const defaultHttpRequestMethod = () => (HttpRequestMethod.GET);
exports.defaultHttpRequestMethod = defaultHttpRequestMethod;
const defaultConnectionCoordinates = () => ({
    x: 0,
    y: 0,
});
exports.defaultConnectionCoordinates = defaultConnectionCoordinates;
var ConnectionPath;
(function (ConnectionPath) {
    ConnectionPath["Straight"] = "straight";
})(ConnectionPath || (exports.ConnectionPath = ConnectionPath = {}));
const defaultConnectionPath = () => (ConnectionPath.Straight);
exports.defaultConnectionPath = defaultConnectionPath;
const defaultCanvasConnection = () => ({
    source: (0, exports.defaultConnectionCoordinates)(),
    target: (0, exports.defaultConnectionCoordinates)(),
    path: ConnectionPath.Straight,
});
exports.defaultCanvasConnection = defaultCanvasConnection;
const defaultCanvasElementOptions = () => ({
    name: "",
    type: "",
});
exports.defaultCanvasElementOptions = defaultCanvasElementOptions;
const defaultOptions = () => ({
    inlineEditing: true,
    showAdvancedTypes: true,
    panZoom: true,
    infinitePan: true,
    root: {
        name: "",
        type: "frame",
        elements: [],
    },
});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map