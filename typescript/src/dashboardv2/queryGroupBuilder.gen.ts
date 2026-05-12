// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class QueryGroupBuilder implements cog.Builder<dashboardv2.QueryGroupKind> {
    protected readonly internal: dashboardv2.QueryGroupKind;
    private nextQueryId: number = 0;

    constructor() {
        this.internal = dashboardv2.defaultQueryGroupKind();
        this.internal.kind = "QueryGroup";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.QueryGroupKind {
        return this.internal;
    }

    targets(queries: cog.Builder<dashboardv2.PanelQueryKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryGroupSpec();
        }
        const queriesResources = queries.map(builder1 => builder1.build());
	queriesResources.forEach((query) => {
		if (query.spec.refId === "") {
			query.spec.refId = `query-${this.nextQueryId}`;
			this.nextQueryId += 1;
		}
	});
        this.internal.spec.queries = queriesResources;
        return this;
    }

    target(querie: cog.Builder<dashboardv2.PanelQueryKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryGroupSpec();
        }
        if (!this.internal.spec.queries) {
            this.internal.spec.queries = [];
        }
        const querieResource = querie.build();
	if (querieResource.spec.refId === "") {
		querieResource.spec.refId = `query-${this.nextQueryId}`;
		this.nextQueryId += 1;
	}
        this.internal.spec.queries.push(querieResource);
        return this;
    }

    transformations(transformations: cog.Builder<dashboardv2.TransformationKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryGroupSpec();
        }
        const transformationsResources = transformations.map(builder1 => builder1.build());
        this.internal.spec.transformations = transformationsResources;
        return this;
    }

    transformation(transformation: cog.Builder<dashboardv2.TransformationKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryGroupSpec();
        }
        if (!this.internal.spec.transformations) {
            this.internal.spec.transformations = [];
        }
        const transformationResource = transformation.build();
        this.internal.spec.transformations.push(transformationResource);
        return this;
    }

    queryOptions(queryOptions: cog.Builder<dashboardv2.QueryOptionsSpec>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultQueryGroupSpec();
        }
        const queryOptionsResource = queryOptions.build();
        this.internal.spec.queryOptions = queryOptionsResource;
        return this;
    }
}

