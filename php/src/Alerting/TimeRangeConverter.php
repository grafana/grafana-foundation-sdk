<?php

namespace Grafana\Foundation\Alerting;

final class TimeRangeConverter
{
    public static function convert(\Grafana\Foundation\Alerting\TimeRange $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\TimeRangeBuilder())',
        ];
            if ($input->from !== null) {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->to !== null) {
    
        
    $buffer = 'to(';
        $arg0 =\var_export($input->to, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

