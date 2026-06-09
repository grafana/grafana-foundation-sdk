<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryV2Converter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryV2Builder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->labels !== null) {
    
        
    $buffer = 'labels(';
        $arg0 = "[";
        foreach ($input->labels as $key => $arg1) {
            $tmplabelsarg1 =\var_export($arg1, true);
            $arg0 .= "\t".var_export($key, true)." => $tmplabelsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2\Dashboardv2DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Cloudwatch\MetricsQuery) {
    
        
    $buffer = 'metricsQuery(';
        $arg0 = \Grafana\Foundation\Cloudwatch\MetricsQueryConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Cloudwatch\LogsQuery) {
    
        
    $buffer = 'logsQuery(';
        $arg0 = \Grafana\Foundation\Cloudwatch\LogsQueryConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Cloudwatch\AnnotationQuery) {
    
        
    $buffer = 'annotationQuery(';
        $arg0 = \Grafana\Foundation\Cloudwatch\AnnotationQueryConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

