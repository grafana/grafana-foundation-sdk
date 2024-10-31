// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorGroupByExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorGroupByExpression> {
    protected readonly internal: cloudwatch.QueryEditorGroupByExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorGroupByExpression();
        this.internal.type = "groupBy";
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorGroupByExpression {
        return this.internal;
    }

    property(property: cog.Builder<cloudwatch.QueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
}
