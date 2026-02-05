"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.GridLayoutItemBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class GridLayoutItemBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultGridLayoutItemKind();
        this.internal.kind = "GridLayoutItem";
        this.internal.spec.element.kind = "ElementReference";
        this.internal.spec.element.name = name;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    x(x) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.x = x;
        return this;
    }
    y(y) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.y = y;
        return this;
    }
    width(width) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.width = width;
        return this;
    }
    height(height) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        this.internal.spec.height = height;
        return this;
    }
    repeat(repeat) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    element(name) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGridLayoutItemSpec();
        }
        if (!this.internal.spec.element) {
            this.internal.spec.element = dashboardv2beta1.defaultElementReference();
        }
        this.internal.spec.element.name = name;
        return this;
    }
}
exports.GridLayoutItemBuilder = GridLayoutItemBuilder;
//# sourceMappingURL=gridLayoutItemBuilder.gen.js.map