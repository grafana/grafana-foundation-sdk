// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class BucketAggregationWithFieldBuilder implements cog.Builder<elasticsearch.BucketAggregationWithField> {
    protected readonly internal: elasticsearch.BucketAggregationWithField;

    constructor() {
        this.internal = elasticsearch.defaultBucketAggregationWithField();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.BucketAggregationWithField {
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

    type(type: elasticsearch.BucketAggregationType): this {
        this.internal.type = type;
        return this;
    }

    settings(settings: any): this {
        this.internal.settings = settings;
        return this;
    }
}
