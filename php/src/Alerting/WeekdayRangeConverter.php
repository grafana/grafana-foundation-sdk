<?php

namespace Grafana\Foundation\Alerting;

final class WeekdayRangeConverter
{
    public static function convert(\Grafana\Foundation\Alerting\WeekdayRange $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\WeekdayRangeBuilder())',
        ];
            if ($input->begin !== null) {
    
        
    $buffer = 'begin(';
        $arg0 =\var_export($input->begin, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->end !== null) {
    
        
    $buffer = 'end(';
        $arg0 =\var_export($input->end, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

