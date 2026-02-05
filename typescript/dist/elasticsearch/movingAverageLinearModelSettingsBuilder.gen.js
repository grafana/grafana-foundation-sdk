"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MovingAverageLinearModelSettingsBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class MovingAverageLinearModelSettingsBuilder {
    constructor() {
        this.internal = elasticsearch.defaultMovingAverageLinearModelSettings();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    window(window) {
        this.internal.window = window;
        return this;
    }
    predict(predict) {
        this.internal.predict = predict;
        return this;
    }
}
exports.MovingAverageLinearModelSettingsBuilder = MovingAverageLinearModelSettingsBuilder;
//# sourceMappingURL=movingAverageLinearModelSettingsBuilder.gen.js.map