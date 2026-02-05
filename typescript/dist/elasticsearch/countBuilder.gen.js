"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CountBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class CountBuilder {
    constructor() {
        this.internal = elasticsearch.defaultCount();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
}
exports.CountBuilder = CountBuilder;
//# sourceMappingURL=countBuilder.gen.js.map