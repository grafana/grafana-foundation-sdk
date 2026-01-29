<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\FetchOptions>
 */
class FetchOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\FetchOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\FetchOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\FetchOptions
     */
    public function build()
    {
        return $this->internal;
    }

    public function method(\Grafana\Foundation\Dashboardv2beta1\HttpRequestMethod $method): static
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

}
