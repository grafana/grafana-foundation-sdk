// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorOperatorExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorOperatorExpression> {
    protected readonly internal: cloudwatch.QueryEditorOperatorExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorOperatorExpression();
        this.internal.type = "operator";
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorOperatorExpression {
        return this.internal;
    }

    property(property: cog.Builder<cloudwatch.QueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }

    // TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
    operator(operator: cog.Builder<cloudwatch.QueryEditorOperator>): this {
        const operatorResource = operator.build();
        this.internal.operator = operatorResource;
        return this;
    }
}
