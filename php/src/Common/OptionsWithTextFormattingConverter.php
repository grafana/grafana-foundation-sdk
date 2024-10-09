<?php

namespace Grafana\Foundation\Common;

final class OptionsWithTextFormattingConverter
{
    public static function convert(\Grafana\Foundation\Common\OptionsWithTextFormatting $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\OptionsWithTextFormattingBuilder())',
        ];
            if ($input->text !== null) {
    
        
    $buffer = 'text(';
        $arg0 = \Grafana\Foundation\Common\VizTextDisplayOptionsConverter::convert($input->text);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

