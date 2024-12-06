// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class DerivativeBuilder implements cog.Builder<elasticsearch.Derivative> {
    protected readonly internal: elasticsearch.Derivative;

    constructor() {
        this.internal = elasticsearch.defaultDerivative();
        this.internal.type = "derivative";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Derivative {
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
	unit?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
