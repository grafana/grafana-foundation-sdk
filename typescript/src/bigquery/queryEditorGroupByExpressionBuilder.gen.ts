// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bigquery from '../bigquery';

export class QueryEditorGroupByExpressionBuilder implements cog.Builder<bigquery.QueryEditorGroupByExpression> {
    protected readonly internal: bigquery.QueryEditorGroupByExpression;

    constructor() {
        this.internal = bigquery.defaultQueryEditorGroupByExpression();
    }

    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorGroupByExpression {
        return this.internal;
    }

    property(property: cog.Builder<bigquery.QueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
}

