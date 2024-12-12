<?php

namespace Grafana\Foundation\Dashboard;

final class ValueMapConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\ValueMap $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\ValueMapBuilder())',
        ];
            
    
        {
    $buffer = 'options(';
        $arg0 = "[";
        foreach ($input->options as $key => $arg1) {
            $tmpoptionsarg1 ='(new \Grafana\Foundation\Dashboard\ValueMappingResult('.(($arg1->text !== null) ? 'text: '.\var_export($arg1->text, true).', ' : '').''.(($arg1->color !== null) ? 'color: '.\var_export($arg1->color, true).', ' : '').''.(($arg1->icon !== null) ? 'icon: '.\var_export($arg1->icon, true).', ' : '').''.(($arg1->index !== null) ? 'index: '.\var_export($arg1->index, true).', ' : '').'))';
            $arg0 .= "\t".var_export($key, true)." => $tmpoptionsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

