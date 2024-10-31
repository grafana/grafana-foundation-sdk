// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
export class AnnotationQueryBuilder implements cog.Builder<dashboard.AnnotationQuery> {
    protected readonly internal: dashboard.AnnotationQuery;

    constructor() {
        this.internal = dashboard.defaultAnnotationQuery();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationQuery {
        return this.internal;
    }

    // Name of annotation.
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Datasource where the annotations data is
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // When enabled the annotation query is issued with every dashboard refresh
    enable(enable: boolean): this {
        this.internal.enable = enable;
        return this;
    }

    // Annotation queries can be toggled on or off at the top of the dashboard.
    // When hide is true, the toggle is not shown in the dashboard.
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Color to use for the annotation event markers
    iconColor(iconColor: string): this {
        this.internal.iconColor = iconColor;
        return this;
    }

    // Filters to apply when fetching annotations
    filter(filter: cog.Builder<dashboard.AnnotationPanelFilter>): this {
        const filterResource = filter.build();
        this.internal.filter = filterResource;
        return this;
    }

    // TODO.. this should just be a normal query target
    target(target: cog.Builder<dashboard.AnnotationTarget>): this {
        const targetResource = target.build();
        this.internal.target = targetResource;
        return this;
    }

    // TODO -- this should not exist here, it is based on the --grafana-- datasource
    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    expr(expr: string): this {
        this.internal.expr = expr;
        return this;
    }
}
