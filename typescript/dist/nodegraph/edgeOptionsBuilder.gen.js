"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.EdgeOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const nodegraph = tslib_1.__importStar(require("../nodegraph"));
class EdgeOptionsBuilder {
    constructor() {
        this.internal = nodegraph.defaultEdgeOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Unit for the main stat to override what ever is set in the data frame.
    mainStatUnit(mainStatUnit) {
        this.internal.mainStatUnit = mainStatUnit;
        return this;
    }
    // Unit for the secondary stat to override what ever is set in the data frame.
    secondaryStatUnit(secondaryStatUnit) {
        this.internal.secondaryStatUnit = secondaryStatUnit;
        return this;
    }
}
exports.EdgeOptionsBuilder = EdgeOptionsBuilder;
//# sourceMappingURL=edgeOptionsBuilder.gen.js.map