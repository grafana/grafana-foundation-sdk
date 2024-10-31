// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class NestedBuilder implements cog.Builder<elasticsearch.Nested> {
    protected readonly internal: elasticsearch.Nested;

    constructor() {
        this.internal = elasticsearch.defaultNested();
        this.internal.type = "nested";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.Nested {
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

    settings(settings: any): this {
        this.internal.settings = settings;
        return this;
    }
}
