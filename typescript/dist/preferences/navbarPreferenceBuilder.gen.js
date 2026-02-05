"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.NavbarPreferenceBuilder = void 0;
const tslib_1 = require("tslib");
const preferences = tslib_1.__importStar(require("../preferences"));
class NavbarPreferenceBuilder {
    constructor() {
        this.internal = preferences.defaultNavbarPreference();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    bookmarkUrls(bookmarkUrls) {
        this.internal.bookmarkUrls = bookmarkUrls;
        return this;
    }
}
exports.NavbarPreferenceBuilder = NavbarPreferenceBuilder;
//# sourceMappingURL=navbarPreferenceBuilder.gen.js.map