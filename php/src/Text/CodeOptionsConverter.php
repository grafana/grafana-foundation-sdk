<?php

namespace Grafana\Foundation\Text;

final class CodeOptionsConverter
{
    public static function convert(\Grafana\Foundation\Text\CodeOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Text\CodeOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'language(';
        $arg0 ='\Grafana\Foundation\Text\CodeLanguage::fromValue("'.$input->language.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->showLineNumbers !== false) {
    
        
    $buffer = 'showLineNumbers(';
        $arg0 =\var_export($input->showLineNumbers, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showMiniMap !== false) {
    
        
    $buffer = 'showMiniMap(';
        $arg0 =\var_export($input->showMiniMap, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

