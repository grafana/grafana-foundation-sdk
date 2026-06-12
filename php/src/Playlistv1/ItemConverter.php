<?php

namespace Grafana\Foundation\Playlistv1;

final class ItemConverter
{
    public static function convert(\Grafana\Foundation\Playlistv1\Item $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Playlistv1\ItemBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Playlistv1\PlaylistItemType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

