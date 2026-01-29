<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Dashboardv2beta1ActionStyleConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyle $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyleBuilder())',
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

