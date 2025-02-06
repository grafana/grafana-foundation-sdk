<?php

namespace Grafana\Foundation\Dashboard;

final class TimeOptionConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\TimeOption $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\TimeOptionBuilder())',
        ];
            if ($input->display !== "") {
    
        
    $buffer = 'display(';
        $arg0 =\var_export($input->display, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
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

