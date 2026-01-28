<?php

namespace Grafana\Foundation\Canvas;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Canvas\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\OptionsBuilder())',
        ];
            if ($input->inlineEditing !== true) {
    
        
    $buffer = 'inlineEditing(';
        $arg0 =\var_export($input->inlineEditing, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showAdvancedTypes !== true) {
    
        
    $buffer = 'showAdvancedTypes(';
        $arg0 =\var_export($input->showAdvancedTypes, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->panZoom !== true) {
    
        
    $buffer = 'panZoom(';
        $arg0 =\var_export($input->panZoom, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'root(';
        $arg0 = \Grafana\Foundation\Canvas\CanvasOptionsRootConverter::convert($input->root);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

