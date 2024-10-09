<?php

namespace Grafana\Foundation\Canvas;

final class CanvasOptionsRootConverter
{
    public static function convert(\Grafana\Foundation\Canvas\CanvasOptionsRoot $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\CanvasOptionsRootBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->elements) >= 1) {
    
        
    $buffer = 'elements(';
        $tmparg0 = [];
        foreach ($input->elements as $arg1) {
        $tmpelementsarg1 = \Grafana\Foundation\Canvas\CanvasElementOptionsConverter::convert($arg1);
        $tmparg0[] = $tmpelementsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

