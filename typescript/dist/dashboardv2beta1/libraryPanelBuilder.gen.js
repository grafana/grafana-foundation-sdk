"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.LibraryPanelBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class LibraryPanelBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultLibraryPanelKind();
        this.internal.kind = "LibraryPanel";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    spec(spec) {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    // Panel ID for the library panel in the dashboard
    id(id) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultLibraryPanelKindSpec();
        }
        this.internal.spec.id = id;
        return this;
    }
    // Title for the library panel in the dashboard
    title(title) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultLibraryPanelKindSpec();
        }
        this.internal.spec.title = title;
        return this;
    }
    libraryPanel(libraryPanel) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultLibraryPanelKindSpec();
        }
        const libraryPanelResource = libraryPanel.build();
        this.internal.spec.libraryPanel = libraryPanelResource;
        return this;
    }
}
exports.LibraryPanelBuilder = LibraryPanelBuilder;
//# sourceMappingURL=libraryPanelBuilder.gen.js.map