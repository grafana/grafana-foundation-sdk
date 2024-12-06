<?php

namespace Grafana\Foundation\Librarypanel;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser>
 */
class LibraryElementDTOMetaUserBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser
     */
    public function build()
    {
        return $this->internal;
    }

    public function id(int $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    public function avatarUrl(string $avatarUrl): static
    {
        $this->internal->avatarUrl = $avatarUrl;
    
        return $this;
    }

}
