// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorPropertyExpressionBuilder implements cog.OptionsBuilder<cloudwatch.QueryEditorPropertyExpression> {
    private readonly internal: cloudwatch.QueryEditorPropertyExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorPropertyExpression();
        this.internal.type = "property";
    }

    build(): cloudwatch.QueryEditorPropertyExpression {
        return this.internal;
    }

    property(property: cog.OptionsBuilder<cloudwatch.QueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
}
