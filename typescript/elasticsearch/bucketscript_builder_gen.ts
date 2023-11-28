// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class BucketScriptBuilder implements cog.OptionsBuilder<elasticsearch.BucketScript> {
    private readonly internal: elasticsearch.BucketScript;

    constructor() {
        this.internal = elasticsearch.defaultBucketScript();
        this.internal.type = "bucket_script";
    }

    build(): elasticsearch.BucketScript {
        return this.internal;
    }

    pipelineVariables(pipelineVariables: cog.OptionsBuilder<elasticsearch.PipelineVariable>[]): this {
        const pipelineVariablesResources = pipelineVariables.map(builder => builder.build());
        this.internal.pipelineVariables = pipelineVariablesResources;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	script?: elasticsearch.InlineScript;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
