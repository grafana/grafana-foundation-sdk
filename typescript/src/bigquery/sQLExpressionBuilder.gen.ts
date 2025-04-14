// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bigquery from '../bigquery';

export class SQLExpressionBuilder implements cog.Builder<bigquery.SQLExpression> {
    protected readonly internal: bigquery.SQLExpression;

    constructor() {
        this.internal = bigquery.defaultSQLExpression();
    }

    /**
     * Builds the object.
     */
    build(): bigquery.SQLExpression {
        return this.internal;
    }

    columns(columns: cog.Builder<bigquery.QueryEditorFunctionExpression>[]): this {
        const columnsResources = columns.map(builder1 => builder1.build());
        this.internal.columns = columnsResources;
        return this;
    }

    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    // whereJsonTree?: _
    whereString(whereString: string): this {
        this.internal.whereString = whereString;
        return this;
    }

    groupBy(groupBy: cog.Builder<bigquery.QueryEditorGroupByExpression>[]): this {
        const groupByResources = groupBy.map(builder1 => builder1.build());
        this.internal.groupBy = groupByResources;
        return this;
    }

    orderBy(orderBy: cog.Builder<bigquery.QueryEditorPropertyExpression>): this {
        const orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
        return this;
    }

    orderByDirection(orderByDirection: bigquery.OrderByDirection): this {
        this.internal.orderByDirection = orderByDirection;
        return this;
    }

    limit(limit: number): this {
        this.internal.limit = limit;
        return this;
    }

    offset(offset: number): this {
        this.internal.offset = offset;
        return this;
    }
}

