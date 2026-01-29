<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class GridLayoutItemConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemBuilder('.\var_export($input->spec->element->name, true).'))',
        ];
            
    
        {
    $buffer = 'x(';
        $arg0 =\var_export($input->spec->x, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'y(';
        $arg0 =\var_export($input->spec->y, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'width(';
        $arg0 =\var_export($input->spec->width, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'height(';
        $arg0 =\var_export($input->spec->height, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->repeat !== null) {
    
        
    $buffer = 'repeat(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\RepeatOptionsConverter::convert($input->spec->repeat);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->element->name !== "") {
    
        
    $buffer = 'element(';
        $arg0 =\var_export($input->spec->element->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

