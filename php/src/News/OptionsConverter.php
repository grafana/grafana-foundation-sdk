<?php

namespace Grafana\Foundation\News;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\News\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\News\OptionsBuilder())',
        ];
            if ($input->feedUrl !== null && $input->feedUrl !== "") {
    
        
    $buffer = 'feedUrl(';
        $arg0 =\var_export($input->feedUrl, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showImage !== null && $input->showImage !== true) {
    
        
    $buffer = 'showImage(';
        $arg0 =\var_export($input->showImage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

