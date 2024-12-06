// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class PipelineMetricAggregationWithMultipleBucketPathsBuilder implements cog.Builder<elasticsearch.PipelineMetricAggregationWithMultipleBucketPaths> {
    protected readonly internal: elasticsearch.PipelineMetricAggregationWithMultipleBucketPaths;

    constructor() {
        this.internal = elasticsearch.defaultPipelineMetricAggregationWithMultipleBucketPaths();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.PipelineMetricAggregationWithMultipleBucketPaths {
        return this.internal;
    }

    pipelineVariables(pipelineVariables: cog.Builder<elasticsearch.PipelineVariable>[]): this {
        const pipelineVariablesResources = pipelineVariables.map(builder1 => builder1.build());
        this.internal.pipelineVariables = pipelineVariablesResources;
        return this;
    }

    type(type: elasticsearch.MetricAggregationType): this {
        this.internal.type = type;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
