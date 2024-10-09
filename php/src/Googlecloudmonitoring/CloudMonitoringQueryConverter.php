<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class CloudMonitoringQueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery);
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQueryBuilder())',
        ];
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->aliasBy !== null && $input->aliasBy !== "") {
    
        
    $buffer = 'aliasBy(';
        $arg0 =\var_export($input->aliasBy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeSeriesList !== null) {
    
        
    $buffer = 'timeSeriesList(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesListConverter::convert($input->timeSeriesList);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeSeriesQuery !== null) {
    
        
    $buffer = 'timeSeriesQuery(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQueryConverter::convert($input->timeSeriesQuery);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sloQuery !== null) {
    
        
    $buffer = 'sloQuery(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\SLOQueryConverter::convert($input->sloQuery);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->promQLQuery !== null) {
    
        
    $buffer = 'promQLQuery(';
        $arg0 = \Grafana\Foundation\Googlecloudmonitoring\PromQLQueryConverter::convert($input->promQLQuery);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->intervalMs !== null) {
    
        
    $buffer = 'intervalMs(';
        $arg0 =\var_export($input->intervalMs, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

