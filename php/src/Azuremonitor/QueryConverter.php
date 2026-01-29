<?php

namespace Grafana\Foundation\Azuremonitor;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->subscription !== null && $input->spec->subscription !== "") {
    
        
    $buffer = 'subscription(';
        $arg0 =\var_export($input->spec->subscription, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->subscriptions !== null && count($input->spec->subscriptions) >= 1) {
    
        
    $buffer = 'subscriptions(';
        $tmparg0 = [];
        foreach ($input->spec->subscriptions as $arg1) {
        $tmpsubscriptionsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpsubscriptionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->azureMonitor !== null) {
    
        
    $buffer = 'azureMonitor(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureMetricQueryConverter::convert($input->spec->azureMonitor);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->azureLogAnalytics !== null) {
    
        
    $buffer = 'azureLogAnalytics(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureLogsQueryConverter::convert($input->spec->azureLogAnalytics);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->azureResourceGraph !== null) {
    
        
    $buffer = 'azureResourceGraph(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureResourceGraphQueryConverter::convert($input->spec->azureResourceGraph);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->azureTraces !== null) {
    
        
    $buffer = 'azureTraces(';
        $arg0 = \Grafana\Foundation\Azuremonitor\AzureTracesQueryConverter::convert($input->spec->azureTraces);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery) {
    
        
    $buffer = 'grafanaTemplateVariableFn(';
        switch (true) {
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\SubscriptionsQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\SubscriptionsQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\ResourceGroupsQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\ResourceGroupsQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\ResourceNamesQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\ResourceNamesQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\MetricNamespaceQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\MetricNamespaceQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\MetricDefinitionsQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\MetricDefinitionsQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\MetricNamesQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\MetricNamesQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\WorkspacesQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\WorkspacesQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            case $input->spec->grafanaTemplateVariableFn instanceof \Grafana\Foundation\Azuremonitor\UnknownQuery:
                $disjunctiongrafanaTemplateVariableFn = \Grafana\Foundation\Azuremonitor\UnknownQueryConverter::convert($input->spec->grafanaTemplateVariableFn);
                $arg0 = $disjunctiongrafanaTemplateVariableFn;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->resourceGroup !== null && $input->spec->resourceGroup !== "") {
    
        
    $buffer = 'resourceGroup(';
        $arg0 =\var_export($input->spec->resourceGroup, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->namespace !== null && $input->spec->namespace !== "") {
    
        
    $buffer = 'namespace(';
        $arg0 =\var_export($input->spec->namespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->resource !== null && $input->spec->resource !== "") {
    
        
    $buffer = 'resource(';
        $arg0 =\var_export($input->spec->resource, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->region !== null && $input->spec->region !== "") {
    
        
    $buffer = 'region(';
        $arg0 =\var_export($input->spec->region, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->customNamespace !== null && $input->spec->customNamespace !== "") {
    
        
    $buffer = 'customNamespace(';
        $arg0 =\var_export($input->spec->customNamespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->datasource !== null) {
    
        
    $buffer = 'oldDatasource(';
        $arg0 ='(new \Grafana\Foundation\Common\DataSourceRef('.(($input->spec->datasource->type !== null) ? 'type: '.\var_export($input->spec->datasource->type, true).', ' : '').''.(($input->spec->datasource->uid !== null) ? 'uid: '.\var_export($input->spec->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Azuremonitor\AzureMonitorQuery && $input->spec->query !== null && $input->spec->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->spec->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

