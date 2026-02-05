"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TermsSettingsBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class TermsSettingsBuilder {
    constructor() {
        this.internal = elasticsearch.defaultTermsSettings();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    order(order) {
        this.internal.order = order;
        return this;
    }
    size(size) {
        this.internal.size = size;
        return this;
    }
    minDocCount(minDocCount) {
        this.internal.min_doc_count = minDocCount;
        return this;
    }
    orderBy(orderBy) {
        this.internal.orderBy = orderBy;
        return this;
    }
    missing(missing) {
        this.internal.missing = missing;
        return this;
    }
}
exports.TermsSettingsBuilder = TermsSettingsBuilder;
//# sourceMappingURL=termsSettingsBuilder.gen.js.map