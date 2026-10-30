<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class RelativeTimeRangeConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\RelativeTimeRange $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\RelativeTimeRangeBuilder())',
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

