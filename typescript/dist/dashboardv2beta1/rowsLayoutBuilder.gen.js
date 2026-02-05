"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RowsLayoutBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class RowsLayoutBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultRowsLayoutKind();
        this.internal.kind = "RowsLayout";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    rows(rows) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutSpec();
        }
        const rowsResources = rows.map(builder1 => builder1.build());
        this.internal.spec.rows = rowsResources;
        return this;
    }
    row(row) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutSpec();
        }
        if (!this.internal.spec.rows) {
            this.internal.spec.rows = [];
        }
        const rowResource = row.build();
        this.internal.spec.rows.push(rowResource);
        return this;
    }
}
exports.RowsLayoutBuilder = RowsLayoutBuilder;
//# sourceMappingURL=rowsLayoutBuilder.gen.js.map