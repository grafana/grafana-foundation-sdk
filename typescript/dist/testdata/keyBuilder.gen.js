"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.KeyBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class KeyBuilder {
    constructor() {
        this.internal = testdata.defaultKey();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    tick(tick) {
        this.internal.tick = tick;
        return this;
    }
    type(type) {
        this.internal.type = type;
        return this;
    }
    uid(uid) {
        this.internal.uid = uid;
        return this;
    }
}
exports.KeyBuilder = KeyBuilder;
//# sourceMappingURL=keyBuilder.gen.js.map