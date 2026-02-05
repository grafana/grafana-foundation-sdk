"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultKey = exports.defaultConnectionArgs = exports.defaultFormatOptions = exports.FormatOptions = exports.defaultDataquery = void 0;
const defaultDataquery = () => ({
    format: FormatOptions.TimeSeries,
    connectionArgs: (0, exports.defaultConnectionArgs)(),
    rawSQL: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultDataquery = defaultDataquery;
var FormatOptions;
(function (FormatOptions) {
    FormatOptions[FormatOptions["TimeSeries"] = 0] = "TimeSeries";
    FormatOptions[FormatOptions["Table"] = 1] = "Table";
    FormatOptions[FormatOptions["Logs"] = 2] = "Logs";
})(FormatOptions || (exports.FormatOptions = FormatOptions = {}));
const defaultFormatOptions = () => (FormatOptions.TimeSeries);
exports.defaultFormatOptions = defaultFormatOptions;
const defaultConnectionArgs = () => ({
    region: "__default",
    catalog: "__default",
    database: "__default",
    resultReuseEnabled: false,
    resultReuseMaxAgeInMinutes: 60,
});
exports.defaultConnectionArgs = defaultConnectionArgs;
exports.defaultKey = "__default";
//# sourceMappingURL=types.gen.js.map