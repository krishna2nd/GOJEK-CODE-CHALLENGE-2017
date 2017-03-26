### Parking lot
The problem solution is written in golang. Application can accept command file 
as well as commands in interactive mode.
Even though you provide a command file it operate and provide a prompt to user to 
operate further instead of operate and quit

#### Application build requirements
    go1.8, golint are required binaries. Executable permission enabled for 
`parking_lot`. Please make sure its in executable mode in your system
    Run 
    `chmod 755 parking_lot`
    in command prompt if not.
   Please make sure your  ***GOPATH*** is set to application directory
   or set as folows.
   
   ***cd [solution directory]***
   
   ***export GOPATH=$(pwd)***
   
    Run above commands in your shell
   
Application includes following components

1. Vehicle
2. Slot
3. Parking
4. Store
5. CommandProcessing

##### Vehicle
    The basic required object in the system. Object have properties
    1. Registration number {uniq identifier}
    2. Colour

##### Slot
    One parking center is devided into number of slots, That means slots
    are the building blocks of parking lot.
    Slots can have a start index, sequence ordering
    
    Slot properties are
        1. Slot Number
        2. Vehicle Object
 
##### Parking Center
    Parking center refers to parking lot. which is the major component in this system
    Parking center have number slots and operating methods for addition,
    removal, reporting of vehicles in slots.
    Major properties of Parking center as follows
        1. Capacity
        2. Slots
    
    Bahaviour 
        1. Add vehicle
        2. Remove vehicle
        3. Reporting status
        4. Other reportings
  
##### Store
    Which act as an in memory data store which keeps an instance of parking center

##### CommandProcessing
    Another backbone of this application, which implemented as a command
    processing modules goes through server stages
    1. Check command
    2. Parse arguments
    3. Verify arguments
    4. Run command
    5. Report output

##### Shell
    Provided an interactive shell for users to operate

##### FileProcessor
    Provide a way to execute list of command from a file. Accept command file
    read the command and execute without user intervension
    
###Functional / Unit Testing Files
    1. commands/command_test.go [Unit Test]
        which test each command one by one and validating the output
        Here we are validating user interactive shell provides expected output.
        The whole system functionality also tested here.
        `execute: go test commands`
        
    2. commands/fileprocess_test.go [Functional Test]
        which test the command output and functionality of command file processing
        while running this test we can make sure that application is running fine in the 
        give requirement.
        `execute: go test commands`
        
    3. parking/parking_test.go [Unit Test]
        Verified the parking center functionaliy Add/Remove vehicles
        Also checks the Benchmarks
        `execute: go test parking`
        
    4. parking/search_test.go [Unit Test]
        Verifies the various Reporting functionalities required to give report based
        on slots, vehicles colour, number status
        `execute: go test parking`
        
    5. slot/slot_test.go [Unit Test]
        Verifies the slot behaviour. Slot must be able to allocate and free
        for a given vehicle.
        `execute: go test slot`
        
    6. vehicle/vehicle_test.go [Unit Test]
        Verifies basic methods defined on vehicles instance. get colour, number
        `execute: go test vehicle`
   