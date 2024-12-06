// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class SerialDiffBuilder implements cog.Builder<elasticsearch.SerialDiff> {
    protected readonly internal: elasticsearch.SerialDiff;

    constructor() {
        this.internal = elasticsearch.defaultSerialDiff();
        this.internal.type = "serial_diff";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.SerialDiff {
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
	lag?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
