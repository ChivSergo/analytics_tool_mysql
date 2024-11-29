Пример использования программы

Impors:
    go run main.go import positions positions.csv
        2024/11/29 16:49:59 8 Raw inserted
    go run main.go import employees employees.csv
        2024/11/29 16:50:15 99 Raw inserted
    go run main.go import timesheet timesheet.csv
        2024/11/29 16:50:35 730 Raw inserted in tasks
        2024/11/29 16:51:01 1765 Raw inserted in Timesheets


List employees:

    go run main.go list employee:
        ...
        Valentina
        Zoe
        Austin
        Ava
        Evelyn
        Sophia
        Leah
        Hudson
        Silas
        Milo
        Braxton
        Ellie
        Ivy
        Amir
        99 Records loaded successfully

Get timesheet for an employee:

    go run main.go get [employeeName]

    На примере Valentina

    go run main.go get Valentina:
    Task Name: SEC-154, Start Time: 2021-01-09 12:00:00 +0700 +07, End Time: 2021-01-09 16:00:00 +0700 +07
    Task Name: ANALYTICS-312, Start Time: 2021-01-22 07:00:00 +0700 +07, End Time: 2021-01-22 09:00:00 +0700 +07
    Task Name: BILLING-998, Start Time: 2021-01-31 11:00:00 +0700 +07, End Time: 2021-01-31 14:00:00 +0700 +07
    Task Name: DATA-139, Start Time: 2021-02-04 16:00:00 +0700 +07, End Time: 2021-02-04 19:00:00 +0700 +07
    Task Name: BILLING-989, Start Time: 2021-02-06 21:00:00 +0700 +07, End Time: 2021-02-07 01:00:00 +0700 +07
    Task Name: DATA-145, Start Time: 2021-02-07 19:00:00 +0700 +07, End Time: 2021-02-07 20:00:00 +0700 +07
    Task Name: ANALYTICS-392, Start Time: 2021-02-12 08:00:00 +0700 +07, End Time: 2021-02-12 10:00:00 +0700 +07
    Task Name: SEC-78, Start Time: 2021-02-21 02:00:00 +0700 +07, End Time: 2021-02-21 05:00:00 +0700 +07
    Task Name: SEC-43, Start Time: 2021-02-25 08:00:00 +0700 +07, End Time: 2021-02-25 12:00:00 +0700 +07
    Task Name: SEC-62, Start Time: 2021-03-08 05:00:00 +0700 +07, End Time: 2021-03-08 08:00:00 +0700 +07
    Task Name: ANALYTICS-374, Start Time: 2021-03-08 11:00:00 +0700 +07, End Time: 2021-03-08 13:00:00 +0700 +07
    Task Name: SEC-151, Start Time: 2021-03-13 07:00:00 +0700 +07, End Time: 2021-03-13 10:00:00 +0700 +07
    Task Name: WEB-174, Start Time: 2021-03-31 13:00:00 +0700 +07, End Time: 2021-03-31 17:00:00 +0700 +07
    13 Records loaded successfully

Remove an employee:

    go run main.go remove [employeeName]

    На примере Valentina
    go run main.go remove Valentina:
    Timesheets for employee Valentina removed successfully
    Проверим
    go run main.go get Valentina:
    0 Records loaded successfully

Top 5 longest tasks:
    go run main.go report top5longTasks
    Task name: SEC-96, Duration: 4h0m0s
    Task name: BILLING-970, Duration: 4h0m0s
    Task name: DATA-226, Duration: 4h0m0s
    Task name: BILLING-969, Duration: 4h0m0s
    Task name: BILLING-970, Duration: 4h0m0s

Top 5 costliest tasks:

    go run main.go report top5costTasks

    Task Name: BILLING-932, Cost: 400.000000
    Task Name: SEC-74, Cost: 400.000000
    Task Name: SEC-103, Cost: 400.000000
    Task Name: BILLING-1026, Cost: 300.000000
    Task Name: SEC-22, Cost: 300.000000

Top 5 employees by total time worked:

    go run main.go report top5employees

    Employee: Savannah, Total Time: 67h0m0s
    Employee: Robert, Total Time: 67h0m0s
    Employee: Leo, Total Time: 65h0m0s
    Employee: Anthony, Total Time: 65h0m0s
    Employee: Dominic, Total Time: 64h0m0s