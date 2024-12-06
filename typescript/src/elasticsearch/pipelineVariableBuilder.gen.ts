// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class PipelineVariableBuilder implements cog.Builder<elasticsearch.PipelineVariable> {
    protected readonly internal: elasticsearch.PipelineVariable;

    constructor() {
        this.internal = elasticsearch.defaultPipelineVariable();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.PipelineVariable {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    pipelineAgg(pipelineAgg: string): this {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }
}
