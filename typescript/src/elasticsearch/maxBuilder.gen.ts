// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MaxBuilder implements cog.Builder<elasticsearch.Max> {
    protected readonly internal: elasticsearch.Max;

    constructor() {
        this.internal = elasticsearch.defaultMax();
        this.internal.type = "max";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Max {
        return this.internal;
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
	script?: elasticsearch.InlineScript;
	missing?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
