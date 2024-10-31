<?php

namespace Grafana\Foundation\Alerting;

final class TimeIntervalTimeRangeConverter
{
    public static function convert(\Grafana\Foundation\Alerting\TimeIntervalTimeRange $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\TimeIntervalTimeRangeBuilder())',
        ];
            if ($input->endTime !== null && $input->endTime !== "") {
    
        
    $buffer = 'endTime(';
        $arg0 =\var_export($input->endTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->startTime !== null && $input->startTime !== "") {
    
        
    $buffer = 'startTime(';
        $arg0 =\var_export($input->startTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

