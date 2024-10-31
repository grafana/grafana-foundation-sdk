// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class BaseBucketAggregationBuilder implements cog.Builder<elasticsearch.BaseBucketAggregation> {
    protected readonly internal: elasticsearch.BaseBucketAggregation;

    constructor() {
        this.internal = elasticsearch.defaultBaseBucketAggregation();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.BaseBucketAggregation {
        return this.internal;
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
