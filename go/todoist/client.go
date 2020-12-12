package todoist

type TodoistClient struct {
}

func Connect() TodoistClient {
  return TodoistClient{}
}

func (c *TodoistClient) CreateEntry(text string) {
}

