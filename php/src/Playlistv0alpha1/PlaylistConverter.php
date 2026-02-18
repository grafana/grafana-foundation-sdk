<?php

namespace Grafana\Foundation\Playlistv0alpha1;

final class PlaylistConverter
{
    public static function convert(\Grafana\Foundation\Playlistv0alpha1\Playlist $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Playlistv0alpha1\PlaylistBuilder('.\var_export($input->title, true).'))',
        ];
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->interval !== "") {
    
        
    $buffer = 'interval(';
        $arg0 =\var_export($input->interval, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->items as $arg1) {
        $tmpitemsarg1 = \Grafana\Foundation\Playlistv0alpha1\ItemConverter::convert($arg1);
        $tmparg0[] = $tmpitemsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

