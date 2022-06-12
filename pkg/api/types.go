package api

import "fmt"

type User struct {
	ID             int         `json:"id"`
	Username       string      `json:"username"`
	Email          string      `json:"email"`
	Color          string      `json:"color"`
	ProfilePicture string      `json:"profilePicture"`
	Initials       string      `json:"initials"`
	Role           int         `json:"role"`
	CustomRole     interface{} `json:"custom_role"`
	LastActive     string      `json:"last_active"`
	DateJoined     string      `json:"date_joined"`
	DateInvited    string      `json:"date_invited"`
}

func (u *User) GoString() string {
	return fmt.Sprintf(`{
		ID: %d,
		Username: %s
		Email: %s
		Color: %s
		ProfilePicture: %s
		Initials: %s
		Role: %d
		CustomRole: %s
		LastActive: %s
		DateJoined: %s
		DateInvited: %s
	}`,
		u.ID,
		u.Username,
		u.Email,
		u.Color,
		u.ProfilePicture,
		u.Initials,
		u.Role,
		u.CustomRole,
		u.LastActive,
		u.DateJoined,
		u.DateInvited,
	)
}

type InvitedBy struct {
	ID             int         `json:"id"`
	Username       string      `json:"username"`
	Color          string      `json:"color"`
	Email          string      `json:"email"`
	Initials       string      `json:"initials"`
	ProfilePicture interface{} `json:"profilePicture"`
}

func (i *InvitedBy) GoString() string {
	return fmt.Sprintf(`{
		ID: %d
		Username: %s
		Color   : %s
		Email   : %s
		Initials: %s
		ProfilePicture: %s
	}`,
		i.ID,
		i.Username,
		i.Color,
		i.Email,
		i.Initials,
		i.ProfilePicture,
	)
}

type Team struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Color   string      `json:"color"`
	Avatar  interface{} `json:"avatar"`
	Members []struct {
		User      User      `json:"user"`
		InvitedBy InvitedBy `json:"invited_by,omitempty"`
	} `json:"members"`
}

type GetTeamResponse struct {
	Teams []Team `json:"teams"`
}

type SpaceStatus struct {
	Status     string `json:"status"`
	Type       string `json:"type"`
	Orderindex int    `json:"orderindex"`
	Color      string `json:"color"`
}

type SpaceFeature struct {
	DueDates struct {
		Enabled            bool `json:"enabled"`
		StartDate          bool `json:"start_date"`
		RemapDueDates      bool `json:"remap_due_dates"`
		RemapClosedDueDate bool `json:"remap_closed_due_date"`
	} `json:"due_dates"`
	TimeTracking struct {
		Enabled bool `json:"enabled"`
	} `json:"time_tracking"`
	Tags struct {
		Enabled bool `json:"enabled"`
	} `json:"tags"`
	TimeEstimates struct {
		Enabled bool `json:"enabled"`
	} `json:"time_estimates"`
	Checklists struct {
		Enabled bool `json:"enabled"`
	} `json:"checklists"`
	CustomFields struct {
		Enabled bool `json:"enabled"`
	} `json:"custom_fields"`
	RemapDependencies struct {
		Enabled bool `json:"enabled"`
	} `json:"remap_dependencies"`
	DependencyWarning struct {
		Enabled bool `json:"enabled"`
	} `json:"dependency_warning"`
	Portfolios struct {
		Enabled bool `json:"enabled"`
	} `json:"portfolios"`
}

type Space struct {
	ID                string        `json:"id"`
	Name              string        `json:"name"`
	Private           bool          `json:"private"`
	Statuses          []SpaceStatus `json:"statuses"`
	MultipleAssignees bool          `json:"multiple_assignees"`
	Features          SpaceFeature  `json:"features,omitempty"`
}

type GetSpacesResponse struct {
	Spaces []Space `json:"spaces"`
}

type Folder struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Orderindex       int    `json:"orderindex"`
	OverrideStatuses bool   `json:"override_statuses"`
	Hidden           bool   `json:"hidden"`
	Space            struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"space"`
	TaskCount string        `json:"task_count"`
	Lists     []interface{} `json:"lists"`
}

type GetFoldersResponse = []Folder

type List struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Orderindex int         `json:"orderindex"`
	Status     interface{} `json:"status"`
	Priority   interface{} `json:"priority"`
	Assignee   interface{} `json:"assignee"`
	TaskCount  int         `json:"task_count"`
	DueDate    interface{} `json:"due_date"`
	StartDate  interface{} `json:"start_date"`
	Folder     struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"space"`
	Archived         bool   `json:"archived"`
	OverrideStatuses bool   `json:"override_statuses"`
	PermissionLevel  string `json:"permission_level"`
}

type GetListsResponse struct {
	Lists []List `json:"lists"`
}

type Task struct {
	ID          string      `json:"id"`
	CustomID    interface{} `json:"custom_id,omitempty"`
	Name        string      `json:"name"`
	TextContent string      `json:"text_content"`
	Description string      `json:"description"`
	Status      struct {
		Status     string `json:"status"`
		Color      string `json:"color"`
		Orderindex int    `json:"orderindex"`
		Type       string `json:"type"`
	} `json:"status"`
	Orderindex  string      `json:"orderindex"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated"`
	DateClosed  interface{} `json:"date_closed"`
	Creator     struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"creator"`
	Assignees    []interface{} `json:"assignees"`
	Checklists   []interface{} `json:"checklists"`
	Tags         []interface{} `json:"tags"`
	Parent       interface{}   `json:"parent"`
	Priority     interface{}   `json:"priority"`
	DueDate      interface{}   `json:"due_date"`
	StartDate    interface{}   `json:"start_date"`
	TimeEstimate interface{}   `json:"time_estimate"`
	TimeSpent    interface{}   `json:"time_spent"`
	CustomFields []struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Type           string `json:"type"`
		DateCreated    string `json:"date_created"`
		HideFromGuests bool   `json:"hide_from_guests"`
		Value          string `json:"value"`
		Required       bool   `json:"required"`
		TypeConfig     struct {
			SingleUser         bool `json:"single_user"`
			IncludeGroups      bool `json:"include_groups"`
			IncludeGuests      bool `json:"include_guests"`
			IncludeTeamMembers bool `json:"include_team_members"`
		} `json:"type_config,omitempty"`
	} `json:"custom_fields"`
	List struct {
		ID string `json:"id"`
	} `json:"list"`
	Folder struct {
		ID string `json:"id"`
	} `json:"folder"`
	Space struct {
		ID string `json:"id"`
	} `json:"space"`
	URL string `json:"url"`
}

type GetTasksResponse struct {
	Tasks []Task `json:"tasks"`
}
