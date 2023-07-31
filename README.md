# Site Monitoring

This is a website monitoring program developed in Go. It is able to check the availability of a list of websites provided in a text file and record logs of its activities.

## Requirements

Before running the program, make sure you have the Go environment installed on your machine.

## Execution

1. Download and install Go from [golang.org](https://go.dev/).
2. Download the source code for this project.
3. Create a text file called "sites.txt" containing the list of URLs you want to monitor. Each URL must be on a separate line.
4. Open a terminal and navigate to the directory where the source code and "sites.txt" file are located.
5. Run the following command to compile and run the program:

```
go run main.go

```

## Functionalities

The program has the following features:

1. Start Monitoring: This option allows you to start monitoring the sites listed in the "sites.txt" file. The program will perform site availability checks at regular intervals and record the results in a log file called "log.txt".

2. View Logs: This option allows you to view the contents of the log file "log.txt", which contains information about the availability status of the sites and the times they were checked.

3. Exit Program: This option closes the execution of the program.

## Files

- `main.go`: Contains the main program source code.
- `sites.txt`: Text file where the sites to be monitored should be listed, each on a separate line.
- `log.txt`: Log file where information about the availability status of sites is recorded.

## Important notes

- The program uses the standard "net/http" library to check the availability of websites.
- At each monitoring cycle, the program checks all sites listed in "sites.txt" and records the status of each of them in the log file "log.txt" with a date and time stamp.
- The program monitors every 5 seconds (value defined by the delay constant) and repeats the process 3 times (value defined by the monitoring constant).

## Final considerations

This is a simple project for educational purposes and can be extended and improved in many ways. Feel free to modify and adapt it as per your needs. If you have any questions or suggestions, feel free to contact us.

Have fun monitoring your websites! 😊
