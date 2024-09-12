<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeMathDatasource>
 */
class ExprTypeMathDatasourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeMathDatasource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeMathDatasource();
    $this->internal->type = "__expr__";
    }

    /**
     * @return \Grafana\Foundation\Expr\ExprTypeMathDatasource
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
