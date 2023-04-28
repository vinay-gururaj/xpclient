package xpclient

type XPProject struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Id             string `json:"id,omitempty"`
	OrganizationId string `json:"organizationId,omitempty"`
}

type SQSConnection struct {
	AwsAccessKeyId     string `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey,omitempty"`
	ReadQueue          string `json:"readQueue,omitempty"`
	WriteQueue         string `json:"writeQueue,omitempty"`
	AwsRegion          string `json:"awsRegion,omitempty"`
	ProjectId          string `json:"projectId,omitempty"`
	EnvironmentId      string `json:"environmentId,omitempty"`
	ConnectionId       string `json:"connectionId,omitempty"`
	ConnectorId        string `json:"connectorId,omitempty"`
	QueueType          string `json:"queueType,omitempty"`
}

type S3Connection struct {
	AwsAccessKeyId     string `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey string `json:"awsSecretAccessKey,omitempty"`
	AwsBucket          string `json:"awsBucket,omitempty"`
	ProjectId          string `json:"projectId,omitempty"`
	EnvironmentId      string `json:"environmentId,omitempty"`
	ConnectionId       string `json:"connectionId,omitempty"`
	ConnectorId        string `json:"connectorId,omitempty"`
	FileType           string `json:"fileType,omitempty"`
}

type Environment struct {
	Name              string   `json:"name,omitempty"`
	Id                string   `json:"id,omitempty"`
	ProjectId         string   `json:"projectId,omitempty"`
	EngineGroupId     string   `json:"engineGroupId,omitempty"`
	EngineStackId     string   `json:"engineStackId,omitempty"`
	Clone             string   `json:"clone,omitempty"`
	ClonedEnvironment string   `json:"clonedEnvironment,omitempty"`
	Tags              []string `json:"tags,omitempty"`
}

type Connections struct {
	Name           string `json:"name,omitempty"`
	ConnectionType string `json:"connectionType,omitempty"`
	ConnectionId   string `json:"connectionId,omitempty"`
	ConnectorId    string `json:"connectorId,omitempty"`
	ProjectId      string `json:"projectId,omitempty"`
}

type DBConnection struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	Port           string `json:"port,omitempty"`
	Database       string `json:"database,omitempty"`
	DatabaseType   string `json:"databaseType,omitempty"`
	IsolationLevel string `json:"isolationLevel,omitempty"`
	ProjectId      string `json:"projectId,omitempty"`
	EnvironmentId  string `json:"environmentId,omitempty"`
	ConnectionId   string `json:"connectionId,omitempty"`
	ConnectorId    string `json:"connectorId,omitempty"`
}
