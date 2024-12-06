// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class FilterBuilder implements cog.Builder<elasticsearch.Filter> {
    protected readonly internal: elasticsearch.Filter;

    constructor() {
        this.internal = elasticsearch.defaultFilter();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Filter {
        return this.internal;
    }

    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    label(label: string): this {
        this.internal.label = label;
        return this;
    }
}
