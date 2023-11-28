// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class SQLExpressionBuilder implements cog.OptionsBuilder<cloudwatch.SQLExpression> {
    private readonly internal: cloudwatch.SQLExpression;

    constructor() {
        this.internal = cloudwatch.defaultSQLExpression();
    }

    build(): cloudwatch.SQLExpression {
        return this.internal;
    }

    // SELECT part of the SQL expression
    select(select: cog.OptionsBuilder<cloudwatch.QueryEditorFunctionExpression>): this {
        const selectResource = select.build();
        this.internal.select = selectResource;
        return this;
    }

    // FROM part of the SQL expression
    from(from: cog.OptionsBuilder<cloudwatch.QueryEditorPropertyExpression> | cog.OptionsBuilder<cloudwatch.QueryEditorFunctionExpression>): this {
        const fromResource = from.build();
        this.internal.from = fromResource;
        return this;
    }

    // WHERE part of the SQL expression
    where(where: cog.OptionsBuilder<cloudwatch.QueryEditorArrayExpression>): this {
        const whereResource = where.build();
        this.internal.where = whereResource;
        return this;
    }

    // GROUP BY part of the SQL expression
    groupBy(groupBy: cog.OptionsBuilder<cloudwatch.QueryEditorArrayExpression>): this {
        const groupByResource = groupBy.build();
        this.internal.groupBy = groupByResource;
        return this;
    }

    // ORDER BY part of the SQL expression
    orderBy(orderBy: cog.OptionsBuilder<cloudwatch.QueryEditorFunctionExpression>): this {
        const orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
        return this;
    }

    // The sort order of the SQL expression, `ASC` or `DESC`
    orderByDirection(orderByDirection: string): this {
        this.internal.orderByDirection = orderByDirection;
        return this;
    }

    // LIMIT part of the SQL expression
    limit(limit: number): this {
        this.internal.limit = limit;
        return this;
    }
}
