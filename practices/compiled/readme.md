This folder 'compiled' contains all compiled versions of the Go script or program.

```bash
1. Build the Go program:
   go build -o ${name_of_compiled_program} ${go_script_file.go}
2. Change the permissions to make it executable:
   chmod a+x ${name_of_compiled_program}
3. Check the permissions:
   ls -l ${name_of_compiled_program}
4. Run the compiled program:
   ./${name_of_compiled_program}
