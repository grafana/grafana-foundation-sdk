// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MinBuilder implements cog.Builder<elasticsearch.Min> {
    protected readonly internal: elasticsearch.Min;

    constructor() {
        this.internal = elasticsearch.defaultMin();
        this.internal.type = "min";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Min {
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
