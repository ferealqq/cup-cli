package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

var Client *ApiClient

type ApiClient struct {
	token string
	base  string
}

func NewClient(token string) {
	Client = &ApiClient{
		token: token,
		base:  "https://api.clickup.com/api/v2/",
	}
}

// Should I implement maybe in go?
// func (a *ApiClient) GetLists(folderId int) (interface{}, error) {
// 	return a.get("folder/"+strconv.Itoa(folderId)+"/list?archived=false")
// }

type TaskQuery struct {
	List int
}

func (a *ApiClient) GetTasks(query TaskQuery) (*GetTasksResponse, error) {
	return getStructure[GetTasksResponse]("list/"+strconv.Itoa(query.List)+"/task", a.base, a.token)
}

func (a *ApiClient) GetFolderlessLists(spaceId int) (*GetListsResponse, error) {
	return getStructure[GetListsResponse]("space/"+strconv.Itoa(spaceId)+"/list", a.base, a.token)
}

func (a *ApiClient) GetFolders(spaceId int) (*GetFoldersResponse, error) {
	return getStructure[GetFoldersResponse]("space/"+strconv.Itoa(spaceId)+"/folder", a.base, a.token)
}

func (a *ApiClient) GetWorkspaces(teamId int) (*GetSpacesResponse, error) {
	return getStructure[GetSpacesResponse]("team/"+strconv.Itoa(teamId)+"/space", a.base, a.token)
}

func (a *ApiClient) GetTeams() (*GetTeamResponse, error) {
	return getStructure[GetTeamResponse]("team", a.base, a.token)
}

func getStructure[M interface{}](path string, base string, token string) (*M, error) {
	if d, err := getJson(path, base, token); err != nil {
		return nil, err
	} else {
		var m M
		if err := mapstructure.Decode(d, &m); err != nil {
			return nil, err
		}
		return &m, nil
	}
}

func getJson(path string, base string, token string) (map[string]interface{}, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", base+path, nil)

	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
