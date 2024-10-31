<?php

namespace Grafana\Foundation\Team;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Team\Team>
 */
class TeamBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Team\Team $internal;

    public function __construct(string $name)
    {
    	$this->internal = new \Grafana\Foundation\Team\Team();
    $this->internal->name = $name;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Team\Team
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Name of the team.
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * Email of the team.
     */
    public function email(string $email): static
    {
        $this->internal->email = $email;
    
        return $this;
    }

}
