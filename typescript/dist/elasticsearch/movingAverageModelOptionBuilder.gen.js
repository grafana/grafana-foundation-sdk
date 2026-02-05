"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MovingAverageModelOptionBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class MovingAverageModelOptionBuilder {
    constructor() {
        this.internal = elasticsearch.defaultMovingAverageModelOption();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    label(label) {
        this.internal.label = label;
        return this;
    }
    value(value) {
        this.internal.value = value;
        return this;
    }
}
exports.MovingAverageModelOptionBuilder = MovingAverageModelOptionBuilder;
//# sourceMappingURL=movingAverageModelOptionBuilder.gen.js.map