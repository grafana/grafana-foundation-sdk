"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DerivativeBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class DerivativeBuilder {
    constructor() {
        this.internal = elasticsearch.defaultDerivative();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    pipelineAgg(pipelineAgg) {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
    field(field) {
        this.internal.field = field;
        return this;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    settings(settings) {
        this.internal.settings = settings;
        return this;
    }
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
}
exports.DerivativeBuilder = DerivativeBuilder;
//# sourceMappingURL=derivativeBuilder.gen.js.map