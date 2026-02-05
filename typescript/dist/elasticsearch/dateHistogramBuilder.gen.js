"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DateHistogramBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class DateHistogramBuilder {
    constructor() {
        this.internal = elasticsearch.defaultDateHistogram();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    field(field) {
        this.internal.field = field;
        return this;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    settings(settings) {
        this.internal.settings = settings;
        return this;
    }
}
exports.DateHistogramBuilder = DateHistogramBuilder;
//# sourceMappingURL=dateHistogramBuilder.gen.js.map