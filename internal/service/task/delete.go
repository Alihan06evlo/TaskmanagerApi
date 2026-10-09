package task

func (s *TaskService) Delete(id int) error {
 return s.repo.Delete(id)
}