import * as cog from '../cog';
import * as dashboard from '../dashboard';
import * as common from '../common';
export declare class AnnotationQueryBuilder implements cog.Builder<dashboard.AnnotationQuery> {
    protected readonly internal: dashboard.AnnotationQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationQuery;
    name(name: string): this;
    datasource(datasource: common.DataSourceRef): this;
    enable(enable: boolean): this;
    hide(hide: boolean): this;
    iconColor(iconColor: string): this;
    filter(filter: cog.Builder<dashboard.AnnotationPanelFilter>): this;
    target(target: cog.Builder<dashboard.AnnotationTarget>): this;
    type(type: string): this;
    builtIn(builtIn: number): this;
    expr(expr: string): this;
}
