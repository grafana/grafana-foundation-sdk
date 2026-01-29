<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TabsLayoutSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpecBuilder())',
        ];
            if (count($input->tabs) >= 1) {
    
        
    $buffer = 'tabs(';
        $tmparg0 = [];
        foreach ($input->tabs as $arg1) {
        $tmptabsarg1 = \Grafana\Foundation\Dashboardv2beta1\TabsLayoutTabConverter::convert($arg1);
        $tmparg0[] = $tmptabsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

