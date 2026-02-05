"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeOptionBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// Counterpart for TypeScript's TimeOption type.
class TimeOptionBuilder {
    constructor() {
        this.internal = dashboard.defaultTimeOption();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    display(display) {
        this.internal.display = display;
        return this;
    }
    from(from) {
        this.internal.from = from;
        return this;
    }
    to(to) {
        this.internal.to = to;
        return this;
    }
}
exports.TimeOptionBuilder = TimeOptionBuilder;
//# sourceMappingURL=timeOptionBuilder.gen.js.map