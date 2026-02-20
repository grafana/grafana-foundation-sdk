<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Infinity options
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\InfinityOptions>
 */
class InfinityOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\InfinityOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\InfinityOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\InfinityOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function method(\Grafana\Foundation\Dashboard\HttpRequestMethod $method): static
    {
        $this->internal->method = $method;
    
        return $this;
    }

    public function url(string $url): static
    {
        $this->internal->url = $url;
    
        return $this;
    }

    public function body(string $body): static
    {
        $this->internal->body = $body;
    
        return $this;
    }

    /**
     * These are 2D arrays of strings, each representing a key-value pair
     * We are defining them this way because we can't generate a go struct that
     * that would have exactly two strings in each sub-array
     * @param array<array<string>> $queryParams
     */
    public function queryParams(array $queryParams): static
    {
        $this->internal->queryParams = $queryParams;
    
        return $this;
    }

    /**
     * @param array<array<string>> $headers
     */
    public function headers(array $headers): static
    {
        $this->internal->headers = $headers;
    
        return $this;
    }

    public function datasourceUid(string $datasourceUid): static
    {
        $this->internal->datasourceUid = $datasourceUid;
    
        return $this;
    }

}
