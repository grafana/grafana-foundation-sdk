"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.KindBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// --- Common types ---
class KindBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultKind();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    kind(kind) {
        this.internal.kind = kind;
        return this;
    }
    spec(spec) {
        this.internal.spec = spec;
        return this;
    }
    metadata(metadata) {
        this.internal.metadata = metadata;
        return this;
    }
}
exports.KindBuilder = KindBuilder;
//# sourceMappingURL=kindBuilder.gen.js.map