<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\QueryBuilder())',
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->aliasBy !== null && $input->spec->aliasBy !== "") {
    
        
    $buffer = 'aliasBy(';
        $arg0 =\var_export($input->spec->aliasBy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->timeSeriesList !== null) {
    
        
    $buffer = 'timeSeriesList(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesListConverter::convert($input->spec->timeSeriesList);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->timeSeriesQuery !== null) {
    
        
    $buffer = 'timeSeriesQuery(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQueryConverter::convert($input->spec->timeSeriesQuery);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->sloQuery !== null) {
    
        
    $buffer = 'sloQuery(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\SLOQueryConverter::convert($input->spec->sloQuery);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->promQLQuery !== null) {
    
        
    $buffer = 'promQLQuery(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\PromQLQueryConverter::convert($input->spec->promQLQuery);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->datasource !== null) {
    
        
    $buffer = 'oldDatasource(';
        $arg0 ='(new \Grafana\Foundation\Common\DataSourceRef('.(($input->spec->datasource->type !== null) ? 'type: '.\var_export($input->spec->datasource->type, true).', ' : '').''.(($input->spec->datasource->uid !== null) ? 'uid: '.\var_export($input->spec->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery && $input->spec->intervalMs !== null) {
    
        
    $buffer = 'intervalMs(';
        $arg0 =\var_export($input->spec->intervalMs, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

