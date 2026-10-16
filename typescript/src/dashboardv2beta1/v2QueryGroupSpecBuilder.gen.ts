// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class V2QueryGroupSpecBuilder implements cog.Builder<dashboardv2beta1.V2QueryGroupSpec> {
    protected readonly internal: dashboardv2beta1.V2QueryGroupSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultV2QueryGroupSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.V2QueryGroupSpec {
        return this.internal;
    }

    queries(queries: cog.Builder<dashboardv2beta1.PanelQueryKind>[]): this {
        const queriesResources = queries.map(builder1 => builder1.build());
        this.internal.queries = queriesResources;
        return this;
    }

    transformations(transformations: cog.Builder<dashboardv2beta1.V2TransformationKind>[]): this {
        const transformationsResources = transformations.map(builder1 => builder1.build());
        this.internal.transformations = transformationsResources;
        return this;
    }

    queryOptions(queryOptions: cog.Builder<dashboardv2beta1.QueryOptionsSpec>): this {
        const queryOptionsResource = queryOptions.build();
        this.internal.queryOptions = queryOptionsResource;
        return this;
    }
}

