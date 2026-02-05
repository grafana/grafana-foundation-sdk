import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as expr from '../expr';
export declare class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
    typeMath(typeMath: cog.Builder<expr.TypeMath>): this;
    typeReduce(typeReduce: cog.Builder<expr.TypeReduce>): this;
    typeResample(typeResample: cog.Builder<expr.TypeResample>): this;
    typeClassicConditions(typeClassicConditions: cog.Builder<expr.TypeClassicConditions>): this;
    typeThreshold(typeThreshold: cog.Builder<expr.TypeThreshold>): this;
    typeSql(typeSql: cog.Builder<expr.TypeSql>): this;
}
