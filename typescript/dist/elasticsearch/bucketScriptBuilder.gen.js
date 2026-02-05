"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.BucketScriptBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class BucketScriptBuilder {
    constructor() {
        this.internal = elasticsearch.defaultBucketScript();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    pipelineVariables(pipelineVariables) {
        const pipelineVariablesResources = pipelineVariables.map(builder1 => builder1.build());
        this.internal.pipelineVariables = pipelineVariablesResources;
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
exports.BucketScriptBuilder = BucketScriptBuilder;
//# sourceMappingURL=bucketScriptBuilder.gen.js.map