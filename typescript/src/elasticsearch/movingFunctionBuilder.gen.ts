// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MovingFunctionBuilder implements cog.Builder<elasticsearch.MovingFunction> {
    protected readonly internal: elasticsearch.MovingFunction;

    constructor() {
        this.internal = elasticsearch.defaultMovingFunction();
        this.internal.type = "moving_fn";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingFunction {
        return this.internal;
    }

    pipelineAgg(pipelineAgg: string): this {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }

    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	window?: string;
	script?: elasticsearch.InlineScript;
	shift?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
