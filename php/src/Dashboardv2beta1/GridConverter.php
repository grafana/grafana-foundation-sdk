<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class GridConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\GridBuilder())',
        ];
            if (count($input->spec->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->spec->items as $arg1) {
        $tmpitemsarg1 = \Grafana\Foundation\Dashboardv2beta1\GridItemConverter::convert($arg1);
        $tmparg0[] = $tmpitemsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

