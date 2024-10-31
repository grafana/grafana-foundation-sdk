// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as team from '../team';

export class TeamBuilder implements cog.Builder<team.Team> {
    protected readonly internal: team.Team;

    constructor(name: string) {
        this.internal = team.defaultTeam();
        this.internal.name = name;
    }

    /**
     * Builds the object.
     */
    build(): team.Team {
        return this.internal;
    }

    // Name of the team.
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Email of the team.
    email(email: string): this {
        this.internal.email = email;
        return this;
    }
}
