<?php

namespace Grafana\Foundation\Canvas;

final class PlacementConverter
{
    public static function convert(\Grafana\Foundation\Canvas\Placement $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\PlacementBuilder())',
        ];
            if ($input->top !== null) {
    
        
    $buffer = 'top(';
        $arg0 =\var_export($input->top, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->left !== null) {
    
        
    $buffer = 'left(';
        $arg0 =\var_export($input->left, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->right !== null) {
    
        
    $buffer = 'right(';
        $arg0 =\var_export($input->right, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->bottom !== null) {
    
        
    $buffer = 'bottom(';
        $arg0 =\var_export($input->bottom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->width !== null) {
    
        
    $buffer = 'width(';
        $arg0 =\var_export($input->width, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->height !== null) {
    
        
    $buffer = 'height(';
        $arg0 =\var_export($input->height, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->rotation !== null) {
    
        
    $buffer = 'rotation(';
        $arg0 =\var_export($input->rotation, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

