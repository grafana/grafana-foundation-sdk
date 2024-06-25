<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeResampleDatasource>
 */
class ExprTypeResampleDatasourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeResampleDatasource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeResampleDatasource();
    $this->internal->type = "__expr__";
    }

    /**
     * @return \Grafana\Foundation\Expr\ExprTypeResampleDatasource
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The apiserver version
     */
    public function apiVersion(string $apiVersion): static
    {
        $this->internal->apiVersion = $apiVersion;
    
        return $this;
    }
    /**
     * Datasource UID (NOTE: name in k8s)
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }

}
