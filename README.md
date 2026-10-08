# homework-sdk-go

Go SDK для API домашки.

## Установка

```bash
go get github.com/твой-username/homework-sdk-go
```

## Использование

```go
client := homework.New(homework.WithBaseURL("https://hw.bulatik.website"))

// публичное
tasks, _ := client.Tasks(ctx, homework.TasksFilter{Date: "2026-10-09"})

// админское — после Login
client.Login(ctx, "username", "password")
adminTasks, _ := client.AdminTasks(ctx, "2026-10-09")
```

## Методы

- Ping(ctx)
- Subjects(ctx)
- Tasks(ctx, filter)
- TasksByDate(ctx, date)
- TaskBySubject(ctx, subject)
- Register(ctx, req)
- Login(ctx, username, password)
- Logout(ctx)
- Me(ctx)
- AdminTasks(ctx, date)
- CreateTask(ctx, req)
- UpdateTask(ctx, id, req)
- DeleteTask(ctx, id)
- AdminSubjects(ctx)
- ScheduleDay(ctx, date)
- ScheduleNext(ctx, subject)