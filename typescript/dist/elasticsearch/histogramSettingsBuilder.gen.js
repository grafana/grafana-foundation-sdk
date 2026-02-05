"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HistogramSettingsBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class HistogramSettingsBuilder {
    constructor() {
        this.internal = elasticsearch.defaultHistogramSettings();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    interval(interval) {
        this.internal.interval = interval;
        return this;
    }
    minDocCount(minDocCount) {
        this.internal.min_doc_count = minDocCount;
        return this;
    }
}
exports.HistogramSettingsBuilder = HistogramSettingsBuilder;
//# sourceMappingURL=histogramSettingsBuilder.gen.js.map