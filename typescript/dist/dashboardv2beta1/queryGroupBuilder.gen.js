"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryGroupBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class QueryGroupBuilder {
    constructor() {
        this.nextQueryId = 0;
        this.internal = dashboardv2beta1.defaultQueryGroupKind();
        this.internal.kind = "QueryGroup";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    targets(queries) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryGroupSpec();
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
    target(querie) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryGroupSpec();
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
    transformations(transformations) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryGroupSpec();
        }
        const transformationsResources = transformations.map(builder1 => builder1.build());
        this.internal.spec.transformations = transformationsResources;
        return this;
    }
    transformation(transformation) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryGroupSpec();
        }
        if (!this.internal.spec.transformations) {
            this.internal.spec.transformations = [];
        }
        const transformationResource = transformation.build();
        this.internal.spec.transformations.push(transformationResource);
        return this;
    }
    queryOptions(queryOptions) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultQueryGroupSpec();
        }
        const queryOptionsResource = queryOptions.build();
        this.internal.spec.queryOptions = queryOptionsResource;
        return this;
    }
}
exports.QueryGroupBuilder = QueryGroupBuilder;
//# sourceMappingURL=queryGroupBuilder.gen.js.map