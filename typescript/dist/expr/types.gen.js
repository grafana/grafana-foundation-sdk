"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultTypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql = exports.defaultTypeSql = exports.defaultTypeThreshold = exports.defaultTypeClassicConditions = exports.defaultTypeResample = exports.defaultTypeReduce = exports.defaultTypeMath = exports.defaultExpr = void 0;
const defaultExpr = () => ((0, exports.defaultTypeMath)());
exports.defaultExpr = defaultExpr;
const defaultTypeMath = () => ({
    expression: "",
    type: "math",
    _implementsDataqueryVariant: () => { },
});
exports.defaultTypeMath = defaultTypeMath;
const defaultTypeReduce = () => ({
    expression: "",
    reducer: "sum",
    type: "reduce",
    _implementsDataqueryVariant: () => { },
});
exports.defaultTypeReduce = defaultTypeReduce;
const defaultTypeResample = () => ({
    downsampler: "sum",
    expression: "",
    type: "resample",
    upsampler: "pad",
    window: "",
    _implementsDataqueryVariant: () => { },
});
exports.defaultTypeResample = defaultTypeResample;
const defaultTypeClassicConditions = () => ({
    conditions: [],
    type: "classic_conditions",
    _implementsDataqueryVariant: () => { },
});
exports.defaultTypeClassicConditions = defaultTypeClassicConditions;
const defaultTypeThreshold = () => ({
    conditions: [],
    expression: "",
    type: "threshold",
    _implementsDataqueryVariant: () => { },
});
exports.defaultTypeThreshold = defaultTypeThreshold;
const defaultTypeSql = () => ({
    expression: "",
    format: "",
    type: "sql",
    _implementsDataqueryVariant: () => { },
});
exports.defaultTypeSql = defaultTypeSql;
const defaultTypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql = () => ({});
exports.defaultTypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql = defaultTypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql;
//# sourceMappingURL=types.gen.js.map