"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RepeatOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class RepeatOptionsBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultRepeatOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    value(value) {
        this.internal.value = value;
        return this;
    }
    direction(direction) {
        this.internal.direction = direction;
        return this;
    }
    maxPerRow(maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
}
exports.RepeatOptionsBuilder = RepeatOptionsBuilder;
//# sourceMappingURL=repeatOptionsBuilder.gen.js.map