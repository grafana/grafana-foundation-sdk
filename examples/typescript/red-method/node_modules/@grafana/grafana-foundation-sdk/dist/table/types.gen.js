"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultFieldConfig = exports.defaultOptions = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultOptions = () => ({
    frameIndex: 0,
    showHeader: true,
    showTypeIcons: false,
    footer: { show: false, reducer: [], countRows: false, },
    cellHeight: common.TableCellHeight.Sm,
});
exports.defaultOptions = defaultOptions;
const defaultFieldConfig = () => (common.defaultTableFieldOptions());
exports.defaultFieldConfig = defaultFieldConfig;
//# sourceMappingURL=types.gen.js.map