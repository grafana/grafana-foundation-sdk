<?php

namespace Grafana\Foundation\Starsv1alpha1;

final class StarsConverter
{
    public static function convert(\Grafana\Foundation\Starsv1alpha1\Stars $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Starsv1alpha1\StarsBuilder())',
        ];
            if (count($input->resource) >= 1) {
    
        
    $buffer = 'resources(';
        $tmparg0 = [];
        foreach ($input->resource as $arg1) {
        $tmpresourcearg1 = \Grafana\Foundation\Starsv1alpha1\ResourceConverter::convert($arg1);
        $tmparg0[] = $tmpresourcearg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

