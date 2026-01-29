<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryBuilder())',
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Cloudwatch\CloudWatchMetricsQuery) {
    
        
    $buffer = 'cloudWatchMetricsQuery(';
        $arg0 = \Grafana\Foundation\Cloudwatch\CloudWatchMetricsQueryConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Cloudwatch\CloudWatchLogsQuery) {
    
        
    $buffer = 'cloudWatchLogsQuery(';
        $arg0 = \Grafana\Foundation\Cloudwatch\CloudWatchLogsQueryConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Cloudwatch\CloudWatchAnnotationQuery) {
    
        
    $buffer = 'cloudWatchAnnotationQuery(';
        $arg0 = \Grafana\Foundation\Cloudwatch\CloudWatchAnnotationQueryConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

