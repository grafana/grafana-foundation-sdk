<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class TimeSeriesQueryConverter
{
    public static function convert(\Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\TimeSeriesQueryBuilder())',
        ];
            if ($input->projectName !== "") {
    
        
    $buffer = 'projectName(';
        $arg0 =\var_export($input->projectName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->graphPeriod !== null && $input->graphPeriod !== "" && $input->graphPeriod !== "disabled") {
    
        
    $buffer = 'graphPeriod(';
        $arg0 =\var_export($input->graphPeriod, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

