"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.OptionsWithTimezonesBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class OptionsWithTimezonesBuilder {
    constructor() {
        this.internal = common.defaultOptionsWithTimezones();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    timezone(timezone) {
        this.internal.timezone = timezone;
        return this;
    }
}
exports.OptionsWithTimezonesBuilder = OptionsWithTimezonesBuilder;
//# sourceMappingURL=optionsWithTimezonesBuilder.gen.js.map