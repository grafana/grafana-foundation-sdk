"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.NodesQueryBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class NodesQueryBuilder {
    constructor() {
        this.internal = testdata.defaultNodesQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    count(count) {
        this.internal.count = count;
        return this;
    }
    seed(seed) {
        this.internal.seed = seed;
        return this;
    }
    // Possible enum values:
    //  - `"random"` 
    //  - `"random edges"` 
    //  - `"response_medium"` 
    //  - `"response_small"` 
    //  - `"feature_showcase"` 
    type(type) {
        this.internal.type = type;
        return this;
    }
}
exports.NodesQueryBuilder = NodesQueryBuilder;
//# sourceMappingURL=nodesQueryBuilder.gen.js.map