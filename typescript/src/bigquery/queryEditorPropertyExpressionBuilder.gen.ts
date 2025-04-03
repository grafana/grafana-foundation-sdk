// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bigquery from '../bigquery';

export class QueryEditorPropertyExpressionBuilder implements cog.Builder<bigquery.QueryEditorPropertyExpression> {
    protected readonly internal: bigquery.QueryEditorPropertyExpression;

    constructor() {
        this.internal = bigquery.defaultQueryEditorPropertyExpression();
    }

    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorPropertyExpression {
        return this.internal;
    }

    property(property: cog.Builder<bigquery.QueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
}

