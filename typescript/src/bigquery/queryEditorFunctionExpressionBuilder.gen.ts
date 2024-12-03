// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bigquery from '../bigquery';

export class QueryEditorFunctionExpressionBuilder implements cog.Builder<bigquery.QueryEditorFunctionExpression> {
    protected readonly internal: bigquery.QueryEditorFunctionExpression;

    constructor() {
        this.internal = bigquery.defaultQueryEditorFunctionExpression();
        this.internal.type = "function";
    }

    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorFunctionExpression {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    parameters(parameters: cog.Builder<bigquery.QueryEditorFunctionParameterExpression>[]): this {
        const parametersResources = parameters.map(builder1 => builder1.build());
        this.internal.parameters = parametersResources;
        return this;
    }
}
