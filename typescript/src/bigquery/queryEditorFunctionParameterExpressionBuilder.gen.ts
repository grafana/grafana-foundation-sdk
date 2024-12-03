// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bigquery from '../bigquery';

export class QueryEditorFunctionParameterExpressionBuilder implements cog.Builder<bigquery.QueryEditorFunctionParameterExpression> {
    protected readonly internal: bigquery.QueryEditorFunctionParameterExpression;

    constructor() {
        this.internal = bigquery.defaultQueryEditorFunctionParameterExpression();
        this.internal.type = "functionParameter";
    }

    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorFunctionParameterExpression {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}
