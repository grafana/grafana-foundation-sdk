// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class PercentilesBuilder implements cog.Builder<elasticsearch.Percentiles> {
    protected readonly internal: elasticsearch.Percentiles;

    constructor() {
        this.internal = elasticsearch.defaultPercentiles();
        this.internal.type = "percentiles";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Percentiles {
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
	percents?: string[];
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
