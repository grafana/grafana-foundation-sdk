"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.GridLayoutBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class GridLayoutBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultGridLayoutKind();
        this.internal.kind = "GridLayout";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    items(items) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }
    item(item) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        const itemResource = item.build();
        this.internal.spec.items.push(itemResource);
        return this;
    }
}
exports.GridLayoutBuilder = GridLayoutBuilder;
//# sourceMappingURL=gridLayoutBuilder.gen.js.map