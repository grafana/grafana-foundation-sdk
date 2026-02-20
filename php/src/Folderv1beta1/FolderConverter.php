<?php

namespace Grafana\Foundation\Folderv1beta1;

final class FolderConverter
{
    public static function convert(\Grafana\Foundation\Folderv1beta1\Folder $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Folderv1beta1\FolderBuilder('.\var_export($input->title, true).'))',
        ];
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

