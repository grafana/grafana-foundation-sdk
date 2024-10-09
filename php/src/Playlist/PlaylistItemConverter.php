<?php

namespace Grafana\Foundation\Playlist;

final class PlaylistItemConverter
{
    public static function convert(\Grafana\Foundation\Playlist\PlaylistItem $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Playlist\PlaylistItemBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Playlist\PlaylistItemType::fromValue("'.$input->type.'")';
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
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

