"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    showLabels: false,
    showCommonLabels: false,
    showTime: false,
    showLogContextToggle: false,
    wrapLogMessage: false,
    prettifyLogMessage: false,
    enableLogDetails: false,
    sortOrder: common.LogsSortOrder.Descending,
    dedupStrategy: common.LogsDedupStrategy.None,
});
exports.defaultOptions = defaultOptions;
//# sourceMappingURL=types.gen.js.map