
ASCII Art Generator :

This Go program generates ASCII art based on a given sentence and an asset file. It reads the content of the asset file, processes it, and prints the corresponding ASCII art for the input sentence.


USAGE :

To run the program, you need to provide exactly two arguments: the sentence you want to convert to ASCII art and the name of the asset file.


EXAMPLE :

>go run main.go "Hello, World!" standard.txt
>
 _    _          _   _                __          __                 _       _   _  
| |  | |        | | | |               \ \        / /                | |     | | | | 
| |__| |   ___  | | | |   ___          \ \  /\  / /    ___    _ __  | |   __| | | | 
|  __  |  / _ \ | | | |  / _ \          \ \/  \/ /    / _ \  | '__| | |  / _` | | | 
| |  | | |  __/ | | | | | (_) |          \  /\  /    | (_) | | |    | | | (_| | |_| 
|_|  |_|  \___| |_| |_|  \___/            \/  \/      \___/  |_|    |_|  \__,_| (_) 
                                                                                    
    

ERROR HANDLING :

The program checks if the correct number of arguments is provided. If not, it prints the following error message and exits:
Error: too many or missing arguments


FUNCTIONS :

lib.GetFileContent(filePath string) (string, string) - Reads the content of the asset file.
lib.CreateTable(asset string, thinkertoy string) map[string]string - Processes the asset file content and creates a table.
lib.PrintAscii(table map[string]string, sentence string) - Prints the ASCII art of the given sentence using the table.

LICENSE :

This project is licensed under the MIT License - see the LICENSE file for details.


AUTHORS :

cnuttens - amorlaya - nsannier
