"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TabsLayoutTabBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class TabsLayoutTabBuilder {
    constructor(title) {
        this.internal = dashboardv2beta1.defaultTabsLayoutTabKind();
        this.internal.kind = "TabsLayoutTab";
        this.internal.spec.title = title;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    title(title) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        this.internal.spec.title = title;
        return this;
    }
    gridLayout(gridLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const gridLayoutKindResource = gridLayoutKind.build();
        this.internal.spec.layout = gridLayoutKindResource;
        return this;
    }
    rowsLayout(rowsLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const rowsLayoutKindResource = rowsLayoutKind.build();
        this.internal.spec.layout = rowsLayoutKindResource;
        return this;
    }
    autoGridLayout(autoGridLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const autoGridLayoutKindResource = autoGridLayoutKind.build();
        this.internal.spec.layout = autoGridLayoutKindResource;
        return this;
    }
    tabsLayout(tabsLayoutKind) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const tabsLayoutKindResource = tabsLayoutKind.build();
        this.internal.spec.layout = tabsLayoutKindResource;
        return this;
    }
    conditionalRendering(conditionalRendering) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    repeat(repeat) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
}
exports.TabsLayoutTabBuilder = TabsLayoutTabBuilder;
//# sourceMappingURL=tabsLayoutTabBuilder.gen.js.map