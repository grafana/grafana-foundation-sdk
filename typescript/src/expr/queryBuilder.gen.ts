// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as expr from '../expr';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "__expr__";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    typeMath(typeMath: cog.Builder<expr.TypeMath>): this {
        const typeMathResource = typeMath.build();
        this.internal.spec = typeMathResource;
        return this;
    }

    typeReduce(typeReduce: cog.Builder<expr.TypeReduce>): this {
        const typeReduceResource = typeReduce.build();
        this.internal.spec = typeReduceResource;
        return this;
    }

    typeResample(typeResample: cog.Builder<expr.TypeResample>): this {
        const typeResampleResource = typeResample.build();
        this.internal.spec = typeResampleResource;
        return this;
    }

    typeClassicConditions(typeClassicConditions: cog.Builder<expr.TypeClassicConditions>): this {
        const typeClassicConditionsResource = typeClassicConditions.build();
        this.internal.spec = typeClassicConditionsResource;
        return this;
    }

    typeThreshold(typeThreshold: cog.Builder<expr.TypeThreshold>): this {
        const typeThresholdResource = typeThreshold.build();
        this.internal.spec = typeThresholdResource;
        return this;
    }

    typeSql(typeSql: cog.Builder<expr.TypeSql>): this {
        const typeSqlResource = typeSql.build();
        this.internal.spec = typeSqlResource;
        return this;
    }
}

