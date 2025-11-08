# Neo-Ed
The UNIX Editor Ed but modernized.

## What is this project about?
This project aims to create an editor which is fully open source, easy to use and Ed-Like. As such, it is released under the MIT License and command- and line-based.

In its core, it works a bit like a shell, where it takes commands and arguments and does something based on the input, except it doesn't use those commands to directly interact with the OS, but to read and manipulate files.

## Planned features
- Add an insert command for inserting text between lines
- Add a find command for finding text
- Add a replace command for replacing text with other text

## Commands
### Append
```console
[File] Ned >> a
```
This command will append text that the user inputs to the end of the file. It uses a ```.``` as its termination symbol. So appending text to a file would look like this:

```console
[File] Ned >> a
APPEND >> This is a simple demonstration
APPEND >> For the append command!
APPEND >> You use it to append text
APPEND >> To the end of a file.
APPEND >> .
[File] Ned >> 
```

### Read
```console
[File] Ned >> r
```
This command will read a file. It will either read a specified line, read a specified range or the entire file. To read an entire file:

```console
[File] Ned >> r
1     | This is a simple demonstration
2     | For the append command!
3     | You use it to append text
4     | To the end of a file.
[File] Ned >>
```

You can also only read a specified line like this:
```console
[File] Ned >> r 2
2     | For the append command!
[File] Ned >>
```

Or you can read a range like this:
```console
[File] Ned >> r 1 3
1     | This is a simple demonstration
2     | For the append command!
3     | You use it to append text
[File] Ned >>
```

### Write
```console
[File] Ned >> w
Successfully wrote buffer to 'File'
[File] Ned >>
```

This command writes the internal buffer, which the editor stores, to the file, essentially saving what you wrote.

### Edit
```console
[File] Ned >> e 2
EDIT 2 >> This is what editing a line looks like!
[File] Ned >> 
```
Unlike the ```Append``` command, ```Edit``` does not use a termination symbol and will simply save your input to that line after pressing Enter.

Edit also inserts what was already in the line into the input, making small edits a lot more comfortable.

### Execute
```console
[File] Ned >> x echo hello world!
hello world!
[File] Ned >>
```

This command is used for executing commands directly from the operating system. It's mainly meant for streamlining development, so that you don't need to exit the editor to, for example, compile or run the program. But of course, it can be used for other things.

### Clear
```console
[File] Ned >> clear
```
This command just clears the screen.

### Quit
```console
[File] Ned >> q
Bye!
```
This command closes Ned.