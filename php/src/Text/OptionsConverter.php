<?php

namespace Grafana\Foundation\Text;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Text\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Text\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Text\TextMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->code !== null) {
    
        
    $buffer = 'code(';
        $arg0 = \Grafana\Foundation\Text\CodeOptionsConverter::convert($input->code);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->content !== "" && $input->content !== "# Title\n\nFor markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)") {
    
        
    $buffer = 'content(';
        $arg0 =\var_export($input->content, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

