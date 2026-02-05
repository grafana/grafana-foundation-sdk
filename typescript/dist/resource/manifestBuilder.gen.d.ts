import * as cog from '../cog';
import * as resource from '../resource';
export declare class ManifestBuilder implements cog.Builder<resource.Manifest> {
    protected readonly internal: resource.Manifest;
    constructor();
    /**
     * Builds the object.
     */
    build(): resource.Manifest;
    apiVersion(apiVersion: string): this;
    kind(kind: string): this;
    metadata(metadata: cog.Builder<resource.Metadata>): this;
    spec(spec: any): this;
}
