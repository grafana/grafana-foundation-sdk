// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class RateBuilder implements cog.Builder<elasticsearch.Rate> {
    protected readonly internal: elasticsearch.Rate;

    constructor() {
        this.internal = elasticsearch.defaultRate();
        this.internal.type = "rate";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Rate {
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
	unit?: string;
	mode?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
