"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.StreamingQueryBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class StreamingQueryBuilder {
    constructor() {
        this.internal = testdata.defaultStreamingQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    bands(bands) {
        this.internal.bands = bands;
        return this;
    }
    noise(noise) {
        this.internal.noise = noise;
        return this;
    }
    speed(speed) {
        this.internal.speed = speed;
        return this;
    }
    spread(spread) {
        this.internal.spread = spread;
        return this;
    }
    // Possible enum values:
    //  - `"fetch"` 
    //  - `"logs"` 
    //  - `"signal"` 
    //  - `"traces"` 
    type(type) {
        this.internal.type = type;
        return this;
    }
    url(url) {
        this.internal.url = url;
        return this;
    }
}
exports.StreamingQueryBuilder = StreamingQueryBuilder;
//# sourceMappingURL=streamingQueryBuilder.gen.js.map