"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.LibraryPanelKindSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class LibraryPanelKindSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultLibraryPanelKindSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Panel ID for the library panel in the dashboard
    id(id) {
        this.internal.id = id;
        return this;
    }
    // Title for the library panel in the dashboard
    title(title) {
        this.internal.title = title;
        return this;
    }
    libraryPanel(libraryPanel) {
        const libraryPanelResource = libraryPanel.build();
        this.internal.libraryPanel = libraryPanelResource;
        return this;
    }
}
exports.LibraryPanelKindSpecBuilder = LibraryPanelKindSpecBuilder;
//# sourceMappingURL=libraryPanelKindSpecBuilder.gen.js.map