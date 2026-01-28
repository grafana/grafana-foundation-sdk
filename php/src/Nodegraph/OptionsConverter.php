<?php

namespace Grafana\Foundation\Nodegraph;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Nodegraph\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Nodegraph\OptionsBuilder())',
        ];
            if ($input->nodes !== null) {
    
        
    $buffer = 'nodes(';
        $arg0 = \Grafana\Foundation\Nodegraph\NodeOptionsConverter::convert($input->nodes);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->edges !== null) {
    
        
    $buffer = 'edges(';
        $arg0 = \Grafana\Foundation\Nodegraph\EdgeOptionsConverter::convert($input->edges);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

