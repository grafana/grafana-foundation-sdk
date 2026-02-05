"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MatcherBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class MatcherBuilder {
    constructor() {
        this.internal = alerting.defaultMatcher();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.Name = name;
        return this;
    }
    type(type) {
        this.internal.Type = type;
        return this;
    }
    value(value) {
        this.internal.Value = value;
        return this;
    }
}
exports.MatcherBuilder = MatcherBuilder;
//# sourceMappingURL=matcherBuilder.gen.js.map