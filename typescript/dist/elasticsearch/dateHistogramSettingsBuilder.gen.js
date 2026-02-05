"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DateHistogramSettingsBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class DateHistogramSettingsBuilder {
    constructor() {
        this.internal = elasticsearch.defaultDateHistogramSettings();
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
    trimEdges(trimEdges) {
        this.internal.trimEdges = trimEdges;
        return this;
    }
    offset(offset) {
        this.internal.offset = offset;
        return this;
    }
    timeZone(timeZone) {
        this.internal.timeZone = timeZone;
        return this;
    }
}
exports.DateHistogramSettingsBuilder = DateHistogramSettingsBuilder;
//# sourceMappingURL=dateHistogramSettingsBuilder.gen.js.map