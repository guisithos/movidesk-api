package models

import (
	"time"
)

type Ticket struct {
	ID                             string             `json:"id"`
	Protocol                       interface{}        `json:"protocol"`
	Type                           int                `json:"type"`
	Subject                        string             `json:"subject"`
	Category                       string             `json:"category"`
	Urgency                        string             `json:"urgency"`
	Status                         string             `json:"status"`
	BaseStatus                     string             `json:"baseStatus"`
	Justification                  interface{}        `json:"justification"`
	Origin                         int                `json:"origin"`
	CreatedDate                    time.Time          `json:"createdDate"`
	IsDeleted                      bool               `json:"isDeleted"`
	OriginEmailAccount             interface{}        `json:"originEmailAccount"`
	Owner                          Owner              `json:"owner"`
	OwnerTeam                      string             `json:"ownerTeam"`
	CreatedBy                      Owner              `json:"createdBy"`
	ServiceFull                    []string           `json:"serviceFull"`
	ServiceFirstLevelId            int                `json:"serviceFirstLevelId"`
	ServiceFirstLevel              string             `json:"serviceFirstLevel"`
	ServiceSecondLevel             string             `json:"serviceSecondLevel"`
	ServiceThirdLevel              string             `json:"serviceThirdLevel"`
	ContactForm                    interface{}        `json:"contactForm"`
	Tags                           []interface{}      `json:"tags"`
	CC                             interface{}        `json:"cc"`
	ResolvedIn                     interface{}        `json:"resolvedIn"`
	ClosedIn                       interface{}        `json:"closedIn"`
	CanceledIn                     interface{}        `json:"canceledIn"`
	ActionCount                    int                `json:"actionCount"`
	LifeTimeWorkingTime            interface{}        `json:"lifeTimeWorkingTime"`
	StoppedTime                    interface{}        `json:"stoppedTime"`
	StoppedTimeWorkingTime         interface{}        `json:"stoppedTimeWorkingTime"`
	ResolvedInFirstCall            bool               `json:"resolvedInFirstCall"`
	ChatWidget                     interface{}        `json:"chatWidget"`
	ChatGroup                      interface{}        `json:"chatGroup"`
	ChatTalkTime                   interface{}        `json:"chatTalkTime"`
	ChatWaitingTime                interface{}        `json:"chatWaitingTime"`
	Sequence                       interface{}        `json:"sequence"`
	SlaAgreement                   string             `json:"slaAgreement"`
	SlaAgreementRule               string             `json:"slaAgreementRule"`
	SlaSolutionTime                int                `json:"slaSolutionTime"`
	SlaResponseTime                int                `json:"slaResponseTime"`
	SlaSolutionChangedByUser       bool               `json:"slaSolutionChangedByUser"`
	SlaSolutionChangedBy           Owner              `json:"slaSolutionChangedBy"`
	SlaSolutionDate                time.Time          `json:"slaSolutionDate"`
	SlaSolutionDateIsPaused        bool               `json:"slaSolutionDateIsPaused"`
	JiraIssueKey                   interface{}        `json:"jiraIssueKey"`
	RedmineIssueId                 interface{}        `json:"redmineIssueId"`
	MovideskTicketNumber           interface{}        `json:"movideskTicketNumber"`
	LinkedToIntegratedTicketNumber interface{}        `json:"linkedToIntegratedTicketNumber"`
	ReopenedIn                     interface{}        `json:"reopenedIn"`
	LastActionDate                 time.Time          `json:"lastActionDate"`
	LastUpdate                     time.Time          `json:"lastUpdate"`
	SlaResponseDate                interface{}        `json:"slaResponseDate"`
	SlaRealResponseDate            interface{}        `json:"slaRealResponseDate"`
	Clients                        []Client           `json:"clients"`
	Actions                        []Action           `json:"actions"`
	ParentTickets                  []interface{}      `json:"parentTickets"`
	ChildrenTickets                []interface{}      `json:"childrenTickets"`
	OwnerHistories                 []OwnerHistory     `json:"ownerHistories"`
	StatusHistories                []StatusHistory    `json:"statusHistories"`
	SatisfactionSurveyResponses    []interface{}      `json:"satisfactionSurveyResponses"`
	CustomFieldValues              []CustomFieldValue `json:"customFieldValues"`
	Assets                         []interface{}      `json:"assets"`
	WebhookEvents                  interface{}        `json:"webhookEvents"`
}

type Owner struct {
	ID           string       `json:"id"`
	PersonType   int          `json:"personType"`
	ProfileType  int          `json:"profileType"`
	BusinessName string       `json:"businessName"`
	Email        string       `json:"email"`
	Phone        string       `json:"phone"`
	IsDeleted    bool         `json:"isDeleted"`
	Organization Organization `json:"organization"`
	Address      interface{}  `json:"address"`
	Complement   interface{}  `json:"complement"`
	CEP          interface{}  `json:"cep"`
	City         interface{}  `json:"city"`
	Bairro       interface{}  `json:"bairro"`
	Number       interface{}  `json:"number"`
	Reference    interface{}  `json:"reference"`
}

type Organization struct {
	ID           string      `json:"id"`
	PersonType   int         `json:"personType"`
	ProfileType  int         `json:"profileType"`
	BusinessName string      `json:"businessName"`
	Email        interface{} `json:"email"`
	Phone        interface{} `json:"phone"`
}

type Client struct {
	ID           string       `json:"id"`
	PersonType   int          `json:"personType"`
	ProfileType  int          `json:"profileType"`
	BusinessName string       `json:"businessName"`
	Email        string       `json:"email"`
	Phone        string       `json:"phone"`
	IsDeleted    bool         `json:"isDeleted"`
	Organization Organization `json:"organization"`
	Address      interface{}  `json:"address"`
	Complement   interface{}  `json:"complement"`
	CEP          interface{}  `json:"cep"`
	City         interface{}  `json:"city"`
	Bairro       interface{}  `json:"bairro"`
	Number       interface{}  `json:"number"`
	Reference    interface{}  `json:"reference"`
}

type Action struct {
	ID               int           `json:"id"`
	Type             int           `json:"type"`
	Origin           int           `json:"origin"`
	Description      string        `json:"description"`
	HtmlDescription  string        `json:"htmlDescription"`
	Status           string        `json:"status"`
	Justification    interface{}   `json:"justification"`
	CreatedDate      time.Time     `json:"createdDate"`
	CreatedBy        Owner         `json:"createdBy"`
	IsDeleted        bool          `json:"isDeleted"`
	TimeAppointments []interface{} `json:"timeAppointments"`
	Attachments      []interface{} `json:"attachments"`
	Expenses         []interface{} `json:"expenses"`
	Tags             []interface{} `json:"tags"`
}

type OwnerHistory struct {
	OwnerTeam                 string      `json:"ownerTeam"`
	Owner                     Owner       `json:"owner"`
	ChangedBy                 Owner       `json:"changedBy"`
	ChangedDate               time.Time   `json:"changedDate"`
	PermanencyTimeFullTime    interface{} `json:"permanencyTimeFullTime"`
	PermanencyTimeWorkingTime interface{} `json:"permanencyTimeWorkingTime"`
}

type StatusHistory struct {
	Status                    string      `json:"status"`
	Justification             interface{} `json:"justification"`
	ChangedBy                 Owner       `json:"changedBy"`
	ChangedDate               time.Time   `json:"changedDate"`
	PermanencyTimeFullTime    interface{} `json:"permanencyTimeFullTime"`
	PermanencyTimeWorkingTime interface{} `json:"permanencyTimeWorkingTime"`
}

type CustomFieldValue struct {
	CustomFieldId     int           `json:"customFieldId"`
	CustomFieldRuleId int           `json:"customFieldRuleId"`
	Line              int           `json:"line"`
	Value             interface{}   `json:"value"`
	Items             []interface{} `json:"items"`
}
