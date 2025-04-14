// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bigquery from '../bigquery';

export class QueryEditorPropertyBuilder implements cog.Builder<bigquery.QueryEditorProperty> {
    protected readonly internal: bigquery.QueryEditorProperty;

    constructor() {
        this.internal = bigquery.defaultQueryEditorProperty();
    }

    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorProperty {
        return this.internal;
    }

    type(type: bigquery.QueryEditorPropertyType): this {
        this.internal.type = type;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}

