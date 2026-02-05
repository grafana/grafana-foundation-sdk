"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.PipelineMetricAggregationWithMultipleBucketPathsBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class PipelineMetricAggregationWithMultipleBucketPathsBuilder {
    constructor() {
        this.internal = elasticsearch.defaultPipelineMetricAggregationWithMultipleBucketPaths();
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
    type(type) {
        this.internal.type = type;
        return this;
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
exports.PipelineMetricAggregationWithMultipleBucketPathsBuilder = PipelineMetricAggregationWithMultipleBucketPathsBuilder;
//# sourceMappingURL=pipelineMetricAggregationWithMultipleBucketPathsBuilder.gen.js.map