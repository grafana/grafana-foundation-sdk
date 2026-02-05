"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TabsLayoutBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class TabsLayoutBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultTabsLayoutKind();
        this.internal.kind = "TabsLayout";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    tabs(tabs) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutSpec();
        }
        const tabsResources = tabs.map(builder1 => builder1.build());
        this.internal.spec.tabs = tabsResources;
        return this;
    }
    tab(tab) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutSpec();
        }
        if (!this.internal.spec.tabs) {
            this.internal.spec.tabs = [];
        }
        const tabResource = tab.build();
        this.internal.spec.tabs.push(tabResource);
        return this;
    }
}
exports.TabsLayoutBuilder = TabsLayoutBuilder;
//# sourceMappingURL=tabsLayoutBuilder.gen.js.map