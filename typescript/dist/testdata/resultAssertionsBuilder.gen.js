"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ResultAssertionsBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class ResultAssertionsBuilder {
    constructor() {
        this.internal = testdata.defaultResultAssertions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Maximum frame count
    maxFrames(maxFrames) {
        this.internal.maxFrames = maxFrames;
        return this;
    }
    // Type asserts that the frame matches a known type structure.
    // Possible enum values:
    //  - `""` 
    //  - `"timeseries-wide"` 
    //  - `"timeseries-long"` 
    //  - `"timeseries-many"` 
    //  - `"timeseries-multi"` 
    //  - `"directory-listing"` 
    //  - `"table"` 
    //  - `"numeric-wide"` 
    //  - `"numeric-multi"` 
    //  - `"numeric-long"` 
    //  - `"log-lines"` 
    type(type) {
        this.internal.type = type;
        return this;
    }
    // TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    // contract documentation https://grafana.github.io/dataplane/contract/.
    typeVersion(typeVersion) {
        this.internal.typeVersion = typeVersion;
        return this;
    }
}
exports.ResultAssertionsBuilder = ResultAssertionsBuilder;
//# sourceMappingURL=resultAssertionsBuilder.gen.js.map