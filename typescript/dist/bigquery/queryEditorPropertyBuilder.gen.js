"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryEditorPropertyBuilder = void 0;
const tslib_1 = require("tslib");
const bigquery = tslib_1.__importStar(require("../bigquery"));
class QueryEditorPropertyBuilder {
    constructor() {
        this.internal = bigquery.defaultQueryEditorProperty();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
}
exports.QueryEditorPropertyBuilder = QueryEditorPropertyBuilder;
//# sourceMappingURL=queryEditorPropertyBuilder.gen.js.map