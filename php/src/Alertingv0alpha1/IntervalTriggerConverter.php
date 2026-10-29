<?php

namespace Grafana\Foundation\Alertingv0alpha1;

final class IntervalTriggerConverter
{
    public static function convert(\Grafana\Foundation\Alertingv0alpha1\IntervalTrigger $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alertingv0alpha1\IntervalTriggerBuilder())',
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

