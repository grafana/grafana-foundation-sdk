import * as cog from '../cog';
import * as team from '../team';
export declare class TeamBuilder implements cog.Builder<team.Team> {
    protected readonly internal: team.Team;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): team.Team;
    email(email: string): this;
    name(name: string): this;
}
