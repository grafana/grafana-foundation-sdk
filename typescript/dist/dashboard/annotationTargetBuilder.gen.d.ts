import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class AnnotationTargetBuilder implements cog.Builder<dashboard.AnnotationTarget> {
    protected readonly internal: dashboard.AnnotationTarget;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationTarget;
    limit(limit: number): this;
    matchAny(matchAny: boolean): this;
    tags(tags: string[]): this;
    type(type: string): this;
}
