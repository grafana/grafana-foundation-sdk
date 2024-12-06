// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorPropertyExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorPropertyExpression> {
    protected readonly internal: cloudwatch.QueryEditorPropertyExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorPropertyExpression();
        this.internal.type = "property";
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorPropertyExpression {
        return this.internal;
    }

    property(property: cog.Builder<cloudwatch.QueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
}
