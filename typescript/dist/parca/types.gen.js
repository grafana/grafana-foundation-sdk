"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultDataquery = exports.defaultParcaQueryType = exports.ParcaQueryType = void 0;
var ParcaQueryType;
(function (ParcaQueryType) {
    ParcaQueryType["Metrics"] = "metrics";
    ParcaQueryType["Profile"] = "profile";
    ParcaQueryType["Both"] = "both";
})(ParcaQueryType || (exports.ParcaQueryType = ParcaQueryType = {}));
const defaultParcaQueryType = () => (ParcaQueryType.Both);
exports.defaultParcaQueryType = defaultParcaQueryType;
const defaultDataquery = () => ({
    labelSelector: "{}",
    profileTypeId: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultDataquery = defaultDataquery;
//# sourceMappingURL=types.gen.js.map