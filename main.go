package main

import "fmt"

type Task struct {
	Name   string
	isDone bool
	count  int
}

func filterUndoneTasks(tasks []Task) []Task {
	var undoneTasks []Task
	for _, task := range tasks {
		if !task.isDone {
			undoneTasks = append(undoneTasks, task)
		}
	}
	return undoneTasks
}

func printToDoList(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("Задача: %s, Выполнена: %t, №: %d", task.Name, task.isDone, task.count)
		fmt.Println()
	}
}

func isFinished(task *Task) Task {
	task.isDone = true
	return *task
}

func CheckingCompletedTask(task *Task) {
	if task.isDone {
		fmt.Println("Данная задача выполнена")
	} else {
		fmt.Println("Эта задача еще не выполнена")
	}
}
func CountTasks(tasks []Task) int {
	var count int
	for _, task := range tasks {
		if !task.isDone {
			count++
		}
	}
	return count
}

func CountOfCompletedTasks(tasks []Task) int {
	var count int
	for _, task := range tasks {
		if task.isDone {
			count++
		}
	}
	return count
}
func AddTask(tasks []Task, name string, isDone bool, count int) []Task {
	task := Task{name, isDone, count}
	fmt.Println("Задача добавлена")
	return append(tasks, task)
}

func DeleteTask(tasks []Task, name string) []Task {
	var updatedTasks []Task
	for _, task := range tasks {
		if task.Name == name {
			updatedTasks = append(updatedTasks, task)
		}
	}
	fmt.Println("Задача удалена!")
	return updatedTasks
}

func main() {

	toDoList := []Task{
		{Name: "Сходить в магазин", isDone: false, count: 1},
		{Name: "Прибраться", isDone: false, count: 2},
		{Name: "Сходить на тренировку", isDone: false, count: 3},
		{Name: "Кодить", isDone: false, count: 4},
		{Name: "Почитать", isDone: false, count: 5},
		{Name: "Кинуть друзей на деньги", isDone: false, count: 6},
		{Name: "Выпить пиво", isDone: false, count: 7},
	}

	taskk := &toDoList[0]
	task := &toDoList[1]
	fmt.Println(isFinished(taskk))
	fmt.Println(isFinished(task))
	fmt.Println()
	printToDoList(toDoList)
	fmt.Println()
	fmt.Println()
	fmt.Println(filterUndoneTasks(toDoList))
	fmt.Println()
	fmt.Println()
	CheckingCompletedTask(task)
	CheckingCompletedTask(taskk)
	CheckingCompletedTask(&toDoList[4])
	fmt.Println(CountOfCompletedTasks(toDoList))
	fmt.Println()
	fmt.Println(CountTasks(toDoList))
	DeleteTask(toDoList, "Выпить пиво")
	AddTask(toDoList, "Помыть машину", false, 8)
	fmt.Println(CountTasks(toDoList))
	DeleteTask(toDoList, "Почитать")
	fmt.Println(CountTasks(toDoList))
}
