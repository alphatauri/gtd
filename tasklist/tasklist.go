package tasklist

type TaskList struct {
	Tasks []*Task
}

func (list *TaskList) Size() int {
	return len(list.Tasks)
}

func (list *TaskList) AddTask(task *Task) {
	list.Tasks = append(list.Tasks, task)
}

func (list *TaskList) GetTask(i int) *Task {
	return list.Tasks[i]
}

func InitTaskList() *TaskList {
	inbox := new(TaskList)
	inbox.Tasks = make([]*Task, 0, 10)
	return inbox
}

func (list *TaskList) ToStorableString() string {
	return "{}"
}

func (list *TaskList) LoadFromStorableString(data string) {

}

func (list *TaskList) FileName() string {
	return "inbox.json"
}