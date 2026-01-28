// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class RowsLayoutSpecBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutSpec> {
    protected readonly internal: dashboardv2beta1.RowsLayoutSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultRowsLayoutSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutSpec {
        return this.internal;
    }

    rows(rows: cog.Builder<dashboardv2beta1.RowsLayoutRowKind>[]): this {
        const rowsResources = rows.map(builder1 => builder1.build());
        this.internal.rows = rowsResources;
        return this;
    }
}

