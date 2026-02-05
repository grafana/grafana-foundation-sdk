"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryHistoryPreferenceBuilder = void 0;
const tslib_1 = require("tslib");
const preferences = tslib_1.__importStar(require("../preferences"));
class QueryHistoryPreferenceBuilder {
    constructor() {
        this.internal = preferences.defaultQueryHistoryPreference();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // one of: '' | 'query' | 'starred';
    homeTab(homeTab) {
        this.internal.homeTab = homeTab;
        return this;
    }
}
exports.QueryHistoryPreferenceBuilder = QueryHistoryPreferenceBuilder;
//# sourceMappingURL=queryHistoryPreferenceBuilder.gen.js.map