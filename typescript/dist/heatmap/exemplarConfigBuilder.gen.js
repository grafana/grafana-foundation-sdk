"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExemplarConfigBuilder = void 0;
const tslib_1 = require("tslib");
const heatmap = tslib_1.__importStar(require("../heatmap"));
// Controls exemplar options
class ExemplarConfigBuilder {
    constructor() {
        this.internal = heatmap.defaultExemplarConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Sets the color of the exemplar markers
    color(color) {
        this.internal.color = color;
        return this;
    }
}
exports.ExemplarConfigBuilder = ExemplarConfigBuilder;
//# sourceMappingURL=exemplarConfigBuilder.gen.js.map