<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class RelativeTimeRangeConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\RelativeTimeRangeBuilder())',
        ];
            if ($input->from !== "") {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->to !== "") {
    
        
    $buffer = 'to(';
        $arg0 =\var_export($input->to, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

