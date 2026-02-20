<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardActionStyleConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardActionStyle $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardActionStyleBuilder())',
        ];
            if ($input->backgroundColor !== null && $input->backgroundColor !== "") {
    
        
    $buffer = 'backgroundColor(';
        $arg0 =\var_export($input->backgroundColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

