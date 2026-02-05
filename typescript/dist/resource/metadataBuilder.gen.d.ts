import * as cog from '../cog';
import * as resource from '../resource';
export declare class MetadataBuilder implements cog.Builder<resource.Metadata> {
    protected readonly internal: resource.Metadata;
    constructor();
    /**
     * Builds the object.
     */
    build(): resource.Metadata;
    name(name: string): this;
    namespace(namespace: string): this;
    labels(labels: Record<string, string>): this;
    annotations(annotations: Record<string, string>): this;
    uid(uid: string): this;
    resourceVersion(resourceVersion: string): this;
    generation(generation: number): this;
    creationTimestamp(creationTimestamp: string): this;
    updateTimestamp(updateTimestamp: string): this;
    deletionTimestamp(deletionTimestamp: string): this;
}
