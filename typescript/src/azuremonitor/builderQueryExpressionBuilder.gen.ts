// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryExpression> {
    protected readonly internal: azuremonitor.BuilderQueryExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryExpression {
        return this.internal;
    }

    from(from: cog.Builder<azuremonitor.BuilderQueryEditorPropertyExpression>): this {
        const fromResource = from.build();
        this.internal.from = fromResource;
        return this;
    }

    columns(columns: cog.Builder<azuremonitor.BuilderQueryEditorColumnsExpression>): this {
        const columnsResource = columns.build();
        this.internal.columns = columnsResource;
        return this;
    }

    where(where: cog.Builder<azuremonitor.BuilderQueryEditorWhereExpressionArray>): this {
        const whereResource = where.build();
        this.internal.where = whereResource;
        return this;
    }

    reduce(reduce: cog.Builder<azuremonitor.BuilderQueryEditorReduceExpressionArray>): this {
        const reduceResource = reduce.build();
        this.internal.reduce = reduceResource;
        return this;
    }

    groupBy(groupBy: cog.Builder<azuremonitor.BuilderQueryEditorGroupByExpressionArray>): this {
        const groupByResource = groupBy.build();
        this.internal.groupBy = groupByResource;
        return this;
    }

    limit(limit: number): this {
        this.internal.limit = limit;
        return this;
    }

    orderBy(orderBy: cog.Builder<azuremonitor.BuilderQueryEditorOrderByExpressionArray>): this {
        const orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
        return this;
    }

    fuzzySearch(fuzzySearch: cog.Builder<azuremonitor.BuilderQueryEditorWhereExpressionArray>): this {
        const fuzzySearchResource = fuzzySearch.build();
        this.internal.fuzzySearch = fuzzySearchResource;
        return this;
    }

    timeFilter(timeFilter: cog.Builder<azuremonitor.BuilderQueryEditorWhereExpressionArray>): this {
        const timeFilterResource = timeFilter.build();
        this.internal.timeFilter = timeFilterResource;
        return this;
    }
}

