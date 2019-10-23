# Application setup

To execute the application, we need to run we
bin/parking_lot script from the root of out project directory

We can provide the input by two ways

1. From the command prompt ( we don't need to pass any parameter in the command prompt)

example-

/go/src/parking_lot\$ bin/parking_lot - it will open the command prompt to take the inputs

For each invalid input, it will skip the execution & exit the program on giving exit on the command prompt

2. From the command file ( we need to pass file path in the command prompt as an argument)

example-

/go/src/parking_lot\$ bin/parking_lot /home/username/go/src/parking_lot/bin/file_inputs.txt

- it will read the input from the file path that is provide as an argument

We must provide the full path of the test file, else it will not be able to find the path

Due to the go directory convention, we should always provide the full path of the test directory( easier for test automation)
