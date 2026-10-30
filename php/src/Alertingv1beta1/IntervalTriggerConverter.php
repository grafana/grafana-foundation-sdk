<?php

namespace Grafana\Foundation\Alertingv1beta1;

final class IntervalTriggerConverter
{
    public static function convert(\Grafana\Foundation\Alertingv1beta1\IntervalTrigger $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv1beta1\IntervalTriggerBuilder())',
        ];
            if ($input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

