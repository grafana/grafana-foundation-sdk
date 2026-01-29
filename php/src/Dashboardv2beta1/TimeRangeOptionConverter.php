<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TimeRangeOptionConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TimeRangeOption $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TimeRangeOptionBuilder())',
        ];
            if ($input->display !== "" && $input->display !== "Last 6 hours") {
    
        
    $buffer = 'display(';
        $arg0 =\var_export($input->display, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->from !== "" && $input->from !== "now-6h") {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->to !== "" && $input->to !== "now") {
    
        
    $buffer = 'to(';
        $arg0 =\var_export($input->to, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

