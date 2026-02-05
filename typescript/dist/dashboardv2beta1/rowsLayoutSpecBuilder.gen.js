"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RowsLayoutSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class RowsLayoutSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultRowsLayoutSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    rows(rows) {
        const rowsResources = rows.map(builder1 => builder1.build());
        this.internal.rows = rowsResources;
        return this;
    }
}
exports.RowsLayoutSpecBuilder = RowsLayoutSpecBuilder;
//# sourceMappingURL=rowsLayoutSpecBuilder.gen.js.map