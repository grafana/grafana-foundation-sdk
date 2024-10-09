<?php

namespace Grafana\Foundation\Azuremonitor;

final class AppInsightsGroupByQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AppInsightsGroupByQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AppInsightsGroupByQueryBuilder())',
        ];
            if ($input->rawQuery !== null && $input->rawQuery !== "") {
    
        
    $buffer = 'rawQuery(';
        $arg0 =\var_export($input->rawQuery, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metricName !== "") {
    
        
    $buffer = 'metricName(';
        $arg0 =\var_export($input->metricName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

