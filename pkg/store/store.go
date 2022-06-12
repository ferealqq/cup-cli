package store

import "github.com/ferealqq/cup-util/pkg/api"

// int represents list id
var tasks = make(map[int][]*api.Task)

type Store struct{}

func FetchTasks(query api.TaskQuery, refetch bool) ([]*api.Task, error) {
	if list, ok := tasks[query.List]; refetch || !ok {
		if resp, err := api.Client.GetTasks(query); err != nil {
			return nil, err
		} else {
			var p []*api.Task
			for i := range resp.Tasks {
				p = append(p, &resp.Tasks[i])
			}
			tasks[query.List] = p
			return p, nil
		}
	} else {
		return list, nil
	}
}
