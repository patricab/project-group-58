## Module development in Go
Source: https://go.dev/doc/tutorial/getting-started

TL;DR: Run these commands.

### 1. Do you want to make a new module?
Make a new file _<name>.go_ in a folder called _name_
In the top part of the document write `package name`

Run the command `go mod init name`. This creates the file _name.mod_ in your folder. 

### 2. Do you want to include your brand spanking new module?
If your module is in your local folder, add the line `import "path to your folder"` to your Go file.

You can also include your module by adding this line to your Go file:  `import "name"`  
In this case, you have to run the command  `go mod edit -replace <module name>=<path to your module>` (don't include the triangle brackets "<>")  
You can also add this line to your module file (_name.mod_):  `replace <name> => <path to your module>` (again, don't include the triangle brackets "<>"!!!)  

