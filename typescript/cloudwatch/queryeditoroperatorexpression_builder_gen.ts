// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorOperatorExpressionBuilder implements cog.OptionsBuilder<cloudwatch.QueryEditorOperatorExpression> {
    private readonly internal: cloudwatch.QueryEditorOperatorExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorOperatorExpression();
        this.internal.type = "operator";
    }

    build(): cloudwatch.QueryEditorOperatorExpression {
        return this.internal;
    }

    property(property: cog.OptionsBuilder<cloudwatch.QueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }

    // TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
    operator(operator: cog.OptionsBuilder<cloudwatch.QueryEditorOperator>): this {
        const operatorResource = operator.build();
        this.internal.operator = operatorResource;
        return this;
    }
}
