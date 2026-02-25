// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class RowsBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutKind> {
    protected readonly internal: dashboardv2beta1.RowsLayoutKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultRowsLayoutKind();
        this.internal.kind = "RowsLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutKind {
        return this.internal;
    }

    rows(rows: cog.Builder<dashboardv2beta1.RowsLayoutRowKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutSpec();
        }
        const rowsResources = rows.map(builder1 => builder1.build());
        this.internal.spec.rows = rowsResources;
        return this;
    }

    row(row: cog.Builder<dashboardv2beta1.RowsLayoutRowKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutSpec();
        }
        if (!this.internal.spec.rows) {
            this.internal.spec.rows = [];
        }
        const rowResource = row.build();
        this.internal.spec.rows.push(rowResource);
        return this;
    }
}

export function rows(): RowsBuilder {
	const builder = new RowsBuilder();

	return builder;
}

