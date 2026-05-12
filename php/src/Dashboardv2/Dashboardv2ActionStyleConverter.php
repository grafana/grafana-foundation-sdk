<?php

namespace Grafana\Foundation\Dashboardv2;

final class Dashboardv2ActionStyleConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\Dashboardv2ActionStyle $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\Dashboardv2ActionStyleBuilder())',
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

