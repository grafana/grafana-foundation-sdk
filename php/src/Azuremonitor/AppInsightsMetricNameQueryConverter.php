<?php

namespace Grafana\Foundation\Azuremonitor;

final class AppInsightsMetricNameQueryConverter
{
    public static function convert(\Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Azuremonitor\AppInsightsMetricNameQueryBuilder())',
        ];
            if ($input->rawQuery !== null && $input->rawQuery !== "") {
    
        
    $buffer = 'rawQuery(';
        $arg0 =\var_export($input->rawQuery, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

