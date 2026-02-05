"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.BaseMovingAverageModelSettingsBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class BaseMovingAverageModelSettingsBuilder {
    constructor() {
        this.internal = elasticsearch.defaultBaseMovingAverageModelSettings();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    model(model) {
        this.internal.model = model;
        return this;
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
exports.BaseMovingAverageModelSettingsBuilder = BaseMovingAverageModelSettingsBuilder;
//# sourceMappingURL=baseMovingAverageModelSettingsBuilder.gen.js.map