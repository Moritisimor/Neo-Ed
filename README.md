# Neo-Ed
The UNIX Editor Ed, but modernized.

## What is this project about?
This project aims to create an editor which is fully open source, easy to use and Ed-Like. As such, it is released under the MIT License and command- and line-based.

In its core, it works a bit like a shell, where it takes commands and arguments and does something based on the input, except it doesn't use those commands to directly interact with the OS, but to read and manipulate files.

## Curent Version
The current version is ```v1.4.0```

## Planned features
- Add a simple syntax-highlighting system

## How to compile
To compile Neo-Ed from source you will need a Go-Compiler, preferably the latest. You can download one [here](https://go.dev/ "Official Go Website")

### With ```go install```
You can install Neo-Ed with the ```go install``` command like so:
```bash
go install github.com/Moritisimor/Neo-Ed/cmd/ned@latest
```
You can also use another version if you'd like.

After that, it will be put into the ```bin``` folder of your ```GOROOT```

### Manually
First, you'll need to clone this Repository:
```bash
git clone https://github.com/Moritisimor/Neo-Ed
```

Then, cd into where the Main Program's code lies:
```bash
cd Neo-Ed/cmd/ned
```

And finally compile it:
```bash
go build -ldflags="-s -w" .
```
(The ```-ldflags``` are for making the compiled binary smaller. You can keep those out if you want to.)

## How to use
To use it you simply run ```ned``` followed by the name of the file:
```console
ned File
```

If the file doesn't exist, Ned will create it for you.

After that, you should be greeted by a prompt which looks like this:

```bash
[File] Ned >>
```
The text in the square brackets shows you what file you're editing.

## Commands
### Help
```bash
[File] Ned >> h
```
This command will print helpful text to the terminal.

### Append
```bash
[File] Ned >> a
```
This command will append text that the user inputs to the end of the file. It uses a ```.``` as its termination symbol. So appending text to a file would look like this:

```bash
[File] Ned >> a
APPEND >> This is a simple demonstration
APPEND >> For the append command!
APPEND >> You use it to append text
APPEND >> To the end of a file.
APPEND >> .
[File] Ned >> 
```

### Read
```bash
[File] Ned >> r
```
This command will read a file. It will either read a specified line, read a specified range or the entire file. To read an entire file:

```bash
[File] Ned >> r
1     | This is a simple demonstration
2     | For the append command!
3     | You use it to append text
4     | To the end of a file.
[File] Ned >>
```

You can also only read a specified line like this:
```bash
[File] Ned >> r 2
2     | For the append command!
[File] Ned >>
```

Or you can read a range like this:
```bash
[File] Ned >> r 1 3
1     | This is a simple demonstration
2     | For the append command!
3     | You use it to append text
[File] Ned >>
```

### Delete
```bash
[File] Ned >> d 2
[File] Ned >> r
1     | This is a simple demonstration
2     | You use it to append text
3     | To the end of a file.
[File] Ned >>
```
This command will delete either a specified line, or a range of lines.
```bash
[File] Ned >> d 1 3 
[File] Ned >> r
1     | To the end of a file.
[File] Ned >> 
```

### Insert
```bash
[File] Ned >> i 2
INSERT 2 >> This is inserted text.
INSERT 2 >> .
[File] Ned >> r
1     | This is a simple demonstration
2     | This is inserted text.
3     | For the append command!
4     | You use it to append text
5     | To the end of a file.
[File] Ned >>
```
This command can be used for inserting text between lines. Just like the append command, it uses a . as its termination symbol.

### Find
```bash
[File] Ned >> f This
Match in line 1
1     | This is a simple demonstration
Match in line 2
2     | This is inserted text.
2 Matches.
[File] Ned >> 
```

### Replace
```bash
[File] Ned >> p This That
[File] Ned >> r
1     | That is a simple demonstration
2     | That is inserted text.
3     | For the append command!
4     | You use it to append text
5     | To the end of a file.
[File] Ned >>
```
Replaces Text with other Text

### Write
```bash
[File] Ned >> w
Successfully wrote buffer to 'File'
[File] Ned >>
```

This command writes the internal buffer, which the editor stores, to the file, essentially saving what you wrote.

### Edit
```bash
[File] Ned >> e 2
EDIT 2 >> This is what editing a line looks like!
[File] Ned >> 
```
Unlike the ```Append``` command, ```Edit``` does not use a termination symbol and will simply save your input to that line after pressing Enter.

Edit also inserts what was already in the line into the input, making small edits a lot more comfortable.

### Execute
```bash
[File] Ned >> x echo hello world!
hello world!
[File] Ned >>
```

This command is used for executing commands directly from the operating system. It's mainly meant for streamlining development, so that you don't need to exit the editor to, for example, compile or run the program. But of course, it can be used for other things.

### Clear
```bash
[File] Ned >> clear
```
This command just clears the screen.

### Quit
```bash
[File] Ned >> q
Bye!
```
This command closes Ned.
