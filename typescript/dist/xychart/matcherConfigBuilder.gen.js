"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MatcherConfigBuilder = void 0;
const tslib_1 = require("tslib");
const xychart = tslib_1.__importStar(require("../xychart"));
// NOTE: (copied from dashboard_kind.cue, since not exported)
// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
class MatcherConfigBuilder {
    constructor() {
        this.internal = xychart.defaultMatcherConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // The matcher id. This is used to find the matcher implementation from registry.
    id(id) {
        this.internal.id = id;
        return this;
    }
    // The matcher options. This is specific to the matcher implementation.
    options(options) {
        this.internal.options = options;
        return this;
    }
}
exports.MatcherConfigBuilder = MatcherConfigBuilder;
//# sourceMappingURL=matcherConfigBuilder.gen.js.map