"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MovingAverageEWMAModelSettingsBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class MovingAverageEWMAModelSettingsBuilder {
    constructor() {
        this.internal = elasticsearch.defaultMovingAverageEWMAModelSettings();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    settings(settings) {
        this.internal.settings = settings;
        return this;
    }
    window(window) {
        this.internal.window = window;
        return this;
    }
    minimize(minimize) {
        this.internal.minimize = minimize;
        return this;
    }
    predict(predict) {
        this.internal.predict = predict;
        return this;
    }
}
exports.MovingAverageEWMAModelSettingsBuilder = MovingAverageEWMAModelSettingsBuilder;
//# sourceMappingURL=movingAverageEWMAModelSettingsBuilder.gen.js.map