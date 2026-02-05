"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.LibraryPanelRefBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
class LibraryPanelRefBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultLibraryPanelRef();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Library panel name
    name(name) {
        this.internal.name = name;
        return this;
    }
    // Library panel uid
    uid(uid) {
        this.internal.uid = uid;
        return this;
    }
}
exports.LibraryPanelRefBuilder = LibraryPanelRefBuilder;
//# sourceMappingURL=libraryPanelRefBuilder.gen.js.map