"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.NodeOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const nodegraph = tslib_1.__importStar(require("../nodegraph"));
class NodeOptionsBuilder {
    constructor() {
        this.internal = nodegraph.defaultNodeOptions();
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
    // Define which fields are shown as part of the node arc (colored circle around the node).
    arcs(arcs) {
        const arcsResources = arcs.map(builder1 => builder1.build());
        this.internal.arcs = arcsResources;
        return this;
    }
}
exports.NodeOptionsBuilder = NodeOptionsBuilder;
//# sourceMappingURL=nodeOptionsBuilder.gen.js.map