"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.isBuilder = isBuilder;
function isBuilder(input) {
    if (input === null) {
        return false;
    }
    if (!input?.build) {
        return false;
    }
    return typeof input.build === "function";
}
//# sourceMappingURL=builder_gen.js.map