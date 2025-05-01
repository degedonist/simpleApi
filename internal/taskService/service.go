package taskService

type TaskService interface {
	CreateTask(req RequestBody) (Task error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, req RequestBody) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func (s *taskService) CreateTask(req RequestBody) (Task error) {

}

func (s *taskService) GetAllTasks() ([]Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) UpdateTask(id string, req RequestBody) (Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s *taskService) DeleteTask(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{repo: r}
}
