// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class CountBuilder implements cog.Builder<elasticsearch.Count> {
    protected readonly internal: elasticsearch.Count;

    constructor() {
        this.internal = elasticsearch.defaultCount();
        this.internal.type = "count";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Count {
        return this.internal;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
