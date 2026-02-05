"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TabsLayoutSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class TabsLayoutSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultTabsLayoutSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    tabs(tabs) {
        const tabsResources = tabs.map(builder1 => builder1.build());
        this.internal.tabs = tabsResources;
        return this;
    }
}
exports.TabsLayoutSpecBuilder = TabsLayoutSpecBuilder;
//# sourceMappingURL=tabsLayoutSpecBuilder.gen.js.map