// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class RowsBuilder implements cog.Builder<dashboardv2.RowsLayoutKind> {
    protected readonly internal: dashboardv2.RowsLayoutKind;

    constructor() {
        this.internal = dashboardv2.defaultRowsLayoutKind();
        this.internal.kind = "RowsLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.RowsLayoutKind {
        return this.internal;
    }

    rows(rows: cog.Builder<dashboardv2.RowsLayoutRowKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutSpec();
        }
        const rowsResources = rows.map(builder1 => builder1.build());
        this.internal.spec.rows = rowsResources;
        return this;
    }

    row(row: cog.Builder<dashboardv2.RowsLayoutRowKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutSpec();
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

