// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class UniqueCountBuilder implements cog.Builder<elasticsearch.UniqueCount> {
    protected readonly internal: elasticsearch.UniqueCount;

    constructor() {
        this.internal = elasticsearch.defaultUniqueCount();
        this.internal.type = "cardinality";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.UniqueCount {
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
	precision_threshold?: string;
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
