package xpclient

import (
	"time"
)

type XPProject struct {
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Id                  string `json:"id,omitempty"`
	OrganizationId      string `json:"organizationId,omitempty"`
	YamlSourceProjectId string `json:"YamlSourceProjectId"`
}

type SQSConnection struct {
	AwsAccessKeyId         string `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey     string `json:"awsSecretAccessKey,omitempty"`
	ReadQueue              string `json:"readQueue,omitempty"`
	WriteQueue             string `json:"writeQueue,omitempty"`
	AwsRegion              string `json:"awsRegion,omitempty"`
	ProjectId              string `json:"projectId,omitempty"`
	EnvironmentId          string `json:"environmentId,omitempty"`
	ConnectionId           string `json:"connectionId,omitempty"`
	ConnectorId            string `json:"connectorId,omitempty"`
	QueueType              string `json:"queueType,omitempty"`
	YamlSourceConnectionId string `json:"YamlSourceConnectionId"`
	YamlSourceConnectorId  string `json:"YamlSourceConnectorId"`
}

type S3Connection struct {
	AwsAccessKeyId         string `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey     string `json:"awsSecretAccessKey,omitempty"`
	AwsBucket              string `json:"awsBucket,omitempty"`
	ProjectId              string `json:"projectId,omitempty"`
	EnvironmentId          string `json:"environmentId,omitempty"`
	ConnectionId           string `json:"connectionId,omitempty"`
	ConnectorId            string `json:"connectorId,omitempty"`
	FileType               string `json:"fileType,omitempty"`
	YamlSourceConnectionId string `json:"YamlSourceConnectionId"`
	YamlSourceConnectorId  string `json:"YamlSourceConnectorId"`
}

type APIConnection struct {
	Endpoint               string `json:"endpoint,omitempty"`
	SignatureService       string `json:"signatureService,omitempty"`
	SignatureRegion        string `json:"signatureRegion,omitempty"`
	AwsAccessKeyId         string `json:"awsAccessKeyId,omitempty"`
	AwsSecretAccessKey     string `json:"awsSecretAccessKey,omitempty"`
	MutualAuth             bool   `json:"mutualAuth,omitempty"`
	AwsSignatureEnabled    bool   `json:"awsSignatureEnabled,omitempty"`
	QueryParams            string `json:"queryParams,omitempty"`
	HeaderParams           string `json:"headerParams,omitempty"`
	ProjectId              string `json:"projectId,omitempty"`
	EnvironmentId          string `json:"environmentId,omitempty"`
	ConnectionId           string `json:"connectionId,omitempty"`
	ConnectorId            string `json:"connectorId,omitempty"`
	YamlSourceConnectionId string `json:"YamlSourceConnectionId"`
	YamlSourceConnectorId  string `json:"YamlSourceConnectorId"`
}

type Environment struct {
	Name                    string   `json:"name,omitempty"`
	Id                      string   `json:"id,omitempty"`
	ProjectId               string   `json:"projectId,omitempty"`
	EngineGroupId           string   `json:"engineGroupId,omitempty"`
	EngineStackId           string   `json:"engineStackId,omitempty"`
	Clone                   string   `json:"clone,omitempty"`
	ClonedEnvironment       string   `json:"clonedEnvironment,omitempty"`
	Tags                    []string `json:"tags,omitempty"`
	YamlSourceEnvironmentId string   `json:"YamlSourceEnvironmentId"`
}

type Connections struct {
	Name                   string `json:"name,omitempty"`
	ConnectionType         string `json:"connectionType,omitempty"`
	ConnectionId           string `json:"connectionId,omitempty"`
	ConnectorId            string `json:"connectorId,omitempty"`
	ProjectId              string `json:"projectId,omitempty"`
	YamlSourceConnectionId string `json:"YamlSourceConnectionId"`
	YamlSourceConnectorId  string `json:"YamlSourceConnectorId"`
}

type DBConnection struct {
	Username               string `json:"username,omitempty"`
	Password               string `json:"password,omitempty"`
	Hostname               string `json:"hostname,omitempty"`
	Port                   int    `json:"port,omitempty"`
	Database               string `json:"database,omitempty"`
	DatabaseType           string `json:"databaseType,omitempty"`
	IsolationLevel         string `json:"isolationLevel,omitempty"`
	ProjectId              string `json:"projectId,omitempty"`
	EnvironmentId          string `json:"environmentId,omitempty"`
	ConnectionId           string `json:"connectionId,omitempty"`
	ConnectorId            string `json:"connectorId,omitempty"`
	YamlSourceConnectionId string `json:"YamlSourceConnectionId"`
	YamlSourceConnectorId  string `json:"YamlSourceConnectorId"`
}

type KWProject struct {
	Project struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Schema      struct {
			JAM struct {
				LOG struct {
					Message struct {
						Node             string `json:"Node"`
						Output           string `json:"Output"`
						TransactionState string `json:"TransactionState"`
					} `json:"Message"`
					Category    string `json:"Category"`
					Severity    string `json:"Severity"`
					RequestId   string `json:"RequestId"`
					AmbientData struct {
						KWID  string `json:"KWID"`
						Error struct {
							Code    string `json:"Code"`
							Node    string `json:"Node"`
							Time    string `json:"Time"`
							Message string `json:"Message"`
						} `json:"Error"`
						Status       string `json:"Status"`
						GraphName    string `json:"GraphName"`
						HelpGuide    string `json:"HelpGuide"`
						ElapsedTime  string `json:"ElapsedTime"`
						ProjectName  string `json:"ProjectName"`
						PublishEvent struct {
							Topic     string `json:"Topic"`
							ObjectKey string `json:"ObjectKey"`
							Publisher string `json:"Publisher"`
						} `json:"PublishEvent"`
						TransactionId       string `json:"TransactionId"`
						RecordsChanged      string `json:"RecordsChanged"`
						TransactionEnd      string `json:"TransactionEnd"`
						EnvironmentName     string `json:"EnvironmentName"`
						TransactionType     string `json:"TransactionType"`
						TransactionStart    string `json:"TransactionStart"`
						ElapsedResponseTime string `json:"ElapsedResponseTime"`
					} `json:"AmbientData"`
					MessageType    string `json:"MessageType"`
					TenantClient   string `json:"TenantClient"`
					ConversationId string `json:"ConversationId"`
				} `json:"LOG"`
				Records struct {
					S3 struct {
						Bucket struct {
							Name string `json:"name"`
						} `json:"bucket"`
						Object struct {
							Key string `json:"key"`
						} `json:"object"`
					} `json:"s3"`
				} `json:"Records"`
			} `json:"JAM"`
			Auth struct {
				Host              string `json:"Host"`
				Accept            string `json:"Accept"`
				Connection        string `json:"Connection"`
				XAmzDate          string `json:"X-Amz-Date"`
				ContentType       string `json:"Content-Type"`
				Authorization     string `json:"Authorization"`
				CacheControl      string `json:"Cache-Control"`
				ContentLength     string `json:"Content-Length"`
				AcceptEncoding    string `json:"Accept-Encoding"`
				XAmzContentSha256 string `json:"X-Amz-Content-Sha256"`
			} `json:"auth"`
			Status  string `json:"status"`
			Ingress struct {
				Event              string `json:"Event"`
				Topic              string `json:"Topic"`
				UseCase            string `json:"UseCase"`
				EventType          string `json:"EventType"`
				SourceSystem       string `json:"SourceSystem"`
				LineOfBusiness     string `json:"LineOfBusiness"`
				PatientInformation struct {
					City              string `json:"City"`
					State             string `json:"State"`
					Gender            string `json:"Gender"`
					Zipcode           string `json:"Zipcode"`
					Address1          string `json:"Address1"`
					Address2          string `json:"Address2"`
					Language          string `json:"Language"`
					LastName          string `json:"LastName"`
					CellPhone         string `json:"CellPhone"`
					FirstName         string `json:"FirstName"`
					HomePhone         string `json:"HomePhone"`
					WorkPhone         string `json:"WorkPhone"`
					MiddleName        string `json:"MiddleName"`
					SuffixName        string `json:"SuffixName"`
					DateOfBirth       string `json:"DateOfBirth"`
					PhoneNumber       string `json:"PhoneNumber"`
					EmailAddress      string `json:"EmailAddress"`
					MemberNumber      string `json:"MemberNumber"`
					PhoneNumberType   string `json:"PhoneNumberType"`
					PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
				} `json:"PatientInformation"`
				ProviderInformation struct {
					SubmitterId    string `json:"SubmitterId"`
					DoctorLastName string `json:"DoctorLastName"`
					HOFriendlyName string `json:"HOFriendlyName"`
				} `json:"ProviderInformation"`
				AdditionalParameters     string `json:"AdditionalParameters"`
				CommunicationPreferences struct {
					SMSOptIn   string `json:"SMSOptIn"`
					EmailOptIn string `json:"EmailOptIn"`
					PhoneOptIn string `json:"PhoneOptIn"`
				} `json:"CommunicationPreferences"`
			} `json:"Ingress"`
			Channel string `json:"channel"`
			Encrypt struct {
				Key             string `json:"key"`
				PlainText       string `json:"plainText"`
				EncryptedOutput struct {
					IngressPayload     string `json:"IngressPayload"`
					PatientInformation struct {
						City              string `json:"City"`
						State             string `json:"State"`
						Gender            string `json:"Gender"`
						Zipcode           string `json:"Zipcode"`
						Address1          string `json:"Address1"`
						Address2          string `json:"Address2"`
						Language          string `json:"Language"`
						LastName          string `json:"LastName"`
						CellPhone         string `json:"CellPhone"`
						FirstName         string `json:"FirstName"`
						HomePhone         string `json:"HomePhone"`
						WorkPhone         string `json:"WorkPhone"`
						MiddleName        string `json:"MiddleName"`
						SuffixName        string `json:"SuffixName"`
						DateOfBirth       string `json:"DateOfBirth"`
						PhoneNumber       string `json:"PhoneNumber"`
						EmailAddress      string `json:"EmailAddress"`
						MemberNumber      string `json:"MemberNumber"`
						PhoneNumberType   string `json:"PhoneNumberType"`
						PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
					} `json:"PatientInformation"`
					ProviderInformation struct {
						SubmitterId    string `json:"SubmitterId"`
						DoctorLastName string `json:"DoctorLastName"`
						HOFriendlyName string `json:"HOFriendlyName"`
					} `json:"ProviderInformation"`
				} `json:"EncryptedOutput"`
			} `json:"encrypt"`
			Modules struct {
				ProjectHandlerListener struct {
					Input  string `json:"input"`
					Output string `json:"output"`
				} `json:"ProjectHandlerListener"`
			} `json:"modules"`
			Outreach struct {
				Event          string `json:"Event"`
				MemberNumber   string `json:"MemberNumber"`
				OutreachStatus string `json:"OutreachStatus"`
			} `json:"outreach"`
			Credential struct {
				Key    string `json:"key"`
				Uri    string `json:"uri"`
				Host   string `json:"host"`
				Region string `json:"region"`
				Secret string `json:"secret"`
			} `json:"Credential"`
			Encryption struct {
				StoreRequest struct {
					User        string `json:"user"`
					Project     string `json:"project"`
					Environment string `json:"environment"`
				} `json:"storeRequest"`
				StoreResponse struct {
					Body struct {
						Parameter struct {
							ARN   string `json:"ARN"`
							Name  string `json:"Name"`
							Type  string `json:"Type"`
							Value struct {
								Key    string `json:"key"`
								Vector string `json:"vector"`
							} `json:"Value"`
							Version          int    `json:"Version"`
							DataType         string `json:"DataType"`
							LastModifiedDate string `json:"LastModifiedDate"`
						} `json:"Parameter"`
						ResponseMetadata string `json:"ResponseMetadata"`
					} `json:"body"`
					Status struct {
						Code   int    `json:"code"`
						Reason string `json:"reason"`
					} `json:"status"`
					Headers string `json:"headers"`
				} `json:"storeResponse"`
				EncryptedOutput struct {
					IngressPayload     string `json:"IngressPayload"`
					PatientInformation struct {
						City              string `json:"City"`
						State             string `json:"State"`
						Gender            string `json:"Gender"`
						Zipcode           string `json:"Zipcode"`
						Address1          string `json:"Address1"`
						Address2          string `json:"Address2"`
						Language          string `json:"Language"`
						LastName          string `json:"LastName"`
						CellPhone         string `json:"CellPhone"`
						FirstName         string `json:"FirstName"`
						HomePhone         string `json:"HomePhone"`
						WorkPhone         string `json:"WorkPhone"`
						MiddleName        string `json:"MiddleName"`
						SuffixName        string `json:"SuffixName"`
						DateOfBirth       string `json:"DateOfBirth"`
						PhoneNumber       string `json:"PhoneNumber"`
						EmailAddress      string `json:"EmailAddress"`
						MemberNumber      string `json:"MemberNumber"`
						PhoneNumberType   string `json:"PhoneNumberType"`
						PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
					} `json:"PatientInformation"`
					ProviderInformation struct {
						SubmitterId    string `json:"SubmitterId"`
						DoctorLastName string `json:"DoctorLastName"`
						HOFriendlyName string `json:"HOFriendlyName"`
					} `json:"ProviderInformation"`
				} `json:"EncryptedOutput"`
			} `json:"Encryption"`
			CancelEvent struct {
				Event          string `json:"Event"`
				Topic          string `json:"Topic"`
				EventType      string `json:"EventType"`
				OutreachId     string `json:"outreachId"`
				MemberNumber   string `json:"MemberNumber"`
				SourceSystem   string `json:"SourceSystem"`
				LineOfBusiness string `json:"LineOfBusiness"`
			} `json:"cancelEvent"`
			Transaction struct {
				Scheduler struct {
					RunCount           string `json:"runCount"`
					GraphName          string `json:"graphName"`
					LastRunTs          string `json:"lastRunTs"`
					NextRunTs          string `json:"nextRunTs"`
					CurrentState       string `json:"currentState"`
					LastRunEndTs       string `json:"lastRunEndTs"`
					SecondLastRunTs    string `json:"secondLastRunTs"`
					SecondLastRunEndTs string `json:"secondLastRunEndTs"`
				} `json:"scheduler"`
				SchedulerTS string `json:"schedulerTS"`
			} `json:"transaction"`
			NoEngagement struct {
				Event                   string `json:"event"`
				Channel                 string `json:"channel"`
				UseCase                 string `json:"use_case"`
				MemberId                string `json:"member_id"`
				DateStamp               string `json:"date_stamp"`
				OutreachStatus          string `json:"outreach_status"`
				CommunicationIdentifier string `json:"communication_identifier"`
			} `json:"NoEngagement"`
			InboundEvent struct {
				Event           string `json:"Event"`
				Channel         string `json:"Channel"`
				UseCase         string `json:"UseCase"`
				Timestamp       string `json:"Timestamp"`
				MemberNumber    string `json:"MemberNumber"`
				EventProperties struct {
					Misc1 string `json:"Misc1"`
				} `json:"EventProperties"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"inboundEvent"`
			SmsOutbound struct {
				Topic          string `json:"Topic"`
				Language       string `json:"Language"`
				PhoneNumber    string `json:"PhoneNumber"`
				SMSTemplate    string `json:"SMSTemplate"`
				MemberNumber   string `json:"MemberNumber"`
				SourceSystem   string `json:"SourceSystem"`
				DynamicContent struct {
					City              string `json:"City"`
					State             string `json:"State"`
					Gender            string `json:"Gender"`
					Zipcode           string `json:"Zipcode"`
					Address1          string `json:"Address1"`
					Address2          string `json:"Address2"`
					LastName          string `json:"LastName"`
					CellPhone         string `json:"CellPhone"`
					FirstName         string `json:"FirstName"`
					HomePhone         string `json:"HomePhone"`
					WorkPhone         string `json:"WorkPhone"`
					MiddleName        string `json:"MiddleName"`
					SuffixName        string `json:"SuffixName"`
					DateOfBirth       string `json:"DateOfBirth"`
					EmailAddress      string `json:"EmailAddress"`
					PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
				} `json:"DynamicContent"`
				PhoneNumberType         string `json:"PhoneNumberType"`
				AdditionalParameters    string `json:"AdditionalParameters"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"sms_outbound"`
			JourneySteps struct {
				Steps string `json:"steps"`
			} `json:"journey-steps"`
			OutreachStatus string `json:"OutreachStatus"`
			DatabaseResult string `json:"databaseResult"`
			EmailOutbound  struct {
				Topic          string `json:"Topic"`
				Language       string `json:"Language"`
				EmailAddress   string `json:"EmailAddress"`
				MemberNumber   string `json:"MemberNumber"`
				SourceSystem   string `json:"SourceSystem"`
				EmailTemplate  string `json:"EmailTemplate"`
				DynamicContent struct {
					City              string `json:"City"`
					State             string `json:"State"`
					Gender            string `json:"Gender"`
					Zipcode           string `json:"Zipcode"`
					Address1          string `json:"Address1"`
					Address2          string `json:"Address2"`
					LastName          string `json:"LastName"`
					CellPhone         string `json:"CellPhone"`
					FirstName         string `json:"FirstName"`
					HomePhone         string `json:"HomePhone"`
					WorkPhone         string `json:"WorkPhone"`
					MiddleName        string `json:"MiddleName"`
					SuffixName        string `json:"SuffixName"`
					DateOfBirth       string `json:"DateOfBirth"`
					PhoneNumber       string `json:"PhoneNumber"`
					PhoneNumberType   string `json:"PhoneNumberType"`
					PatientAgeAtOrder string `json:"PatientAgeAtOrder"`
				} `json:"DynamicContent"`
				AdditionalParameters    string `json:"AdditionalParameters"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"email_outbound"`
			EncodedIngress   string `json:"encodedIngress"`
			OutreachRecord   string `json:"outreachRecord"`
			IngressUtilities struct {
				Channel                 string `json:"Channel"`
				DateStamp               string `json:"DateStamp"`
				OutreachId              string `json:"OutreachId"`
				CommunicationIdentifier string `json:"CommunicationIdentifier"`
			} `json:"IngressUtilities"`
			JavascriptResult   string `json:"JavascriptResult"`
			CredentialLiteral  string `json:"CredentialLiteral"`
			UniversalConnector struct {
				Payload struct {
				} `json:"payload "`
				UcSettings struct {
					Url     string `json:"url"`
					Headers struct {
						Host        string `json:"Host"`
						ContentType string `json:"Content-Type"`
					} `json:"headers"`
					Timeout string `json:"timeout"`
					AuthIds struct {
						JWT struct {
							Aud string `json:"aud"`
							Exp string `json:"exp"`
							Iss string `json:"iss"`
							Kid string `json:"kid"`
							Sub string `json:"sub"`
							Typ string `json:"typ"`
						} `json:"JWT"`
						Bearer        string `json:"Bearer"`
						UrlParameters string `json:"url_parameters"`
					} `json:"auth_ids"`
					AuthType      string `json:"auth_type"`
					ServiceType   string `json:"serviceType"`
					UrlParameters struct {
						UtmMedium string `json:"utm_medium"`
						UtmSource string `json:"utm_source"`
					} `json:"url_parameters"`
					RequestPayloadType  string `json:"requestPayloadType"`
					ResponsePayloadType string `json:"responsePayloadType"`
				} `json:"uc_settings"`
			} `json:"UniversalConnector"`
			ReturnFromGCACMValue string `json:"ReturnFromGCACMValue"`
			CancelEventUtilities struct {
				Channel                 string `json:"channel"`
				Payload                 string `json:"payload"`
				UseCase                 string `json:"use_case"`
				Datestamp               string `json:"datestamp"`
				CommunicationIdentifier string `json:"communication_identifier"`
			} `json:"cancelEventUtilities"`
			InboundEventUtilities struct {
				DbResult struct {
					UseCase string `json:"use_case"`
				} `json:"DbResult"`
			} `json:"inboundEventUtilities"`
			UniversalConnectorResponse struct {
				StatusCode struct {
					Code string `json:"code"`
				} `json:"statusCode"`
				ResponseBody struct {
					TokenizedURL string `json:"tokenizedURL"`
				} `json:"responseBody"`
				ResponseHeaders struct {
					XAmznRequestId             string `json:"x-amzn-RequestId"`
					XAmznRemappedAuthorization string `json:"x-amzn-Remapped-Authorization"`
				} `json:"responseHeaders"`
			} `json:"UniversalConnectorResponse"`
			KwServiceTemp20221104        string `json:"kw_service_temp_2022_11_04"`
			CredentialFromParameterStore struct {
				Status     string `json:"status"`
				Credential struct {
					Key struct {
						Crq       string `json:"crq"`
						Latest    string `json:"latest"`
						Secret    string `json:"secret"`
						Source    string `json:"source"`
						Version   string `json:"version"`
						AccessKey string `json:"accessKey"`
						CreatedBy string `json:"createdBy"`
						CreatedOn string `json:"createdOn"`
						ExpiresOn string `json:"expiresOn"`
					} `json:"key"`
					User   string `json:"user"`
					Source string `json:"source"`
				} `json:"credential"`
			} `json:"CredentialFromParameterStore"`
		} `json:"schema"`
		ProfileSchema struct {
			Attributes struct {
				OptOuts struct {
					SMS     string `json:"SMS"`
					Email   string `json:"Email"`
					UseCase string `json:"UseCase"`
				} `json:"opt_outs"`
			} `json:"Attributes"`
			Identifiers struct {
				KWID         string `json:"KWID"`
				MemberNumber string `json:"memberNumber"`
			} `json:"Identifiers"`
			Interactions string `json:"Interactions"`
			JourneySteps string `json:"JourneySteps"`
		} `json:"profileSchema"`
		DefaultEnvironmentId    string      `json:"defaultEnvironmentId"`
		ActiveEnvironmentId     string      `json:"activeEnvironmentId"`
		ProductionEnvironmentId string      `json:"productionEnvironmentId"`
		EmergencyContactListId  string      `json:"emergencyContactListId"`
		UserId                  string      `json:"userId"`
		OrganizationId          string      `json:"organizationId"`
		HasKDMNode              bool        `json:"hasKDMNode"`
		TemplateId              interface{} `json:"templateId"`
		RemoteHub               interface{} `json:"remoteHub"`
		TopLevelId              interface{} `json:"topLevelId"`
		CreatedAt               time.Time   `json:"createdAt"`
		UpdatedAt               time.Time   `json:"updatedAt"`
		Tags                    []struct {
			Id    string `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"tags"`
		Environments []struct {
			Id        string `json:"id"`
			Name      string `json:"name"`
			ProjectId string `json:"projectId"`
		} `json:"environments"`
		Favorites []interface{} `json:"favorites"`
		Metrics   struct {
		} `json:"metrics"`
		SplitTests struct {
		} `json:"splitTests"`
		PublicVariables struct {
		} `json:"publicVariables"`
		Versions    interface{} `json:"versions"`
		Permissions struct {
		} `json:"permissions"`
		Users  []interface{} `json:"users"`
		Groups []interface{} `json:"groups"`
	} `json:"project"`
	Prototypes []struct {
		Id                 string      `json:"id"`
		Type               string      `json:"type"`
		Name               string      `json:"name"`
		Icon               string      `json:"icon"`
		UserId             string      `json:"userId"`
		ProjectId          string      `json:"projectId"`
		ChannelTypeId      *string     `json:"channelTypeId"`
		EditingBy          *string     `json:"editingBy"`
		PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
		CreatedAt          time.Time   `json:"createdAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
		Code               Script
		Graph              interface{}
		Tags               []struct {
			Name      string `json:"name"`
			Color     string `json:"color"`
			IsDefault bool   `json:"isDefault,omitempty"`
			Id        string `json:"id,omitempty"`
		} `json:"tags"`
		IconColor       string `json:"iconColor,omitempty"`
		BackgroundColor string `json:"backgroundColor,omitempty"`
		JourneyStepType string `json:"journeyStepType,omitempty"`
		Outcome         string `json:"outcome,omitempty"`
	} `json:"prototypes"`
	Types struct {
		Deleted bool   `json:"_deleted"`
		Id      string `json:"id"`
	} `json:"types"`
	Request string `json:"request"`
}

type Script struct {
	Prototype struct {
		Id                 string      `json:"id"`
		Type               string      `json:"type"`
		Name               string      `json:"name"`
		Icon               string      `json:"icon"`
		UserId             string      `json:"userId"`
		ProjectId          string      `json:"projectId"`
		ChannelTypeId      interface{} `json:"channelTypeId"`
		EditingBy          string      `json:"editingBy"`
		PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
		CreatedAt          time.Time   `json:"createdAt"`
		UpdatedAt          time.Time   `json:"updatedAt"`
	} `json:"prototype"`
	Type struct {
		Id           string      `json:"id"`
		Icon         string      `json:"icon"`
		Name         string      `json:"name"`
		Body         string      `json:"body"`
		ArgNames     []string    `json:"argNames"`
		ConnectionId interface{} `json:"connectionId"`
		ProjectId    string      `json:"projectId"`
		UserId       string      `json:"userId"`
		Type         string      `json:"type"`
	} `json:"type"`
}

type KWConnections struct {
	Id                  string  `json:"id"`
	Type                string  `json:"type"`
	Name                string  `json:"name"`
	ProjectId           string  `json:"projectId"`
	Managed             bool    `json:"managed"`
	ManagedConnectionId *string `json:"managedConnectionId"`
	Connectors          []struct {
		Id                string      `json:"id"`
		DatabaseType      string      `json:"databaseType,omitempty"`
		Username          *string     `json:"username,omitempty"`
		Password          string      `json:"password,omitempty"`
		Hostname          *string     `json:"hostname,omitempty"`
		Port              *int        `json:"port,omitempty"`
		Database          *string     `json:"database,omitempty"`
		IsolationLevel    interface{} `json:"isolationLevel"`
		ConnectionId      string      `json:"connectionId"`
		EnvironmentId     string      `json:"environmentId"`
		ConnectionType    string      `json:"connection_type,omitempty"`
		Brokers           []string    `json:"brokers,omitempty"`
		Topics            []string    `json:"topics,omitempty"`
		ConnectionDetails struct {
			Algorithm             string `json:"algorithm"`
			Username              string `json:"username,omitempty"`
			Password              string `json:"password,omitempty"`
			ConnectorRequiredName string `json:"connector_required_name,omitempty"`
		} `json:"connection_details,omitempty"`
		CreatedAt   time.Time `json:"createdAt,omitempty"`
		UpdatedAt   time.Time `json:"updatedAt,omitempty"`
		Endpoint    string    `json:"endpoint,omitempty"`
		QueryParams struct {
		} `json:"queryParams,omitempty"`
		HeaderParams struct {
		} `json:"headerParams,omitempty"`
		MutualAuth          bool        `json:"mutualAuth,omitempty"`
		AwsSignatureEnabled bool        `json:"awsSignatureEnabled,omitempty"`
		SignatureService    string      `json:"signatureService,omitempty"`
		SignatureHost       string      `json:"signatureHost,omitempty"`
		SignatureRegion     interface{} `json:"signatureRegion"`
		AwsSecretAccessKey  string      `json:"awsSecretAccessKey,omitempty"`
		QueueType           string      `json:"queueType,omitempty"`
		ReadQueue           string      `json:"readQueue,omitempty"`
		WriteQueue          string      `json:"writeQueue,omitempty"`
		AwsAccessKeyId      *string     `json:"awsAccessKeyId,omitempty"`
		AwsRegion           *string     `json:"awsRegion,omitempty"`
		FileType            string      `json:"fileType,omitempty"`
		AwsBucket           *string     `json:"awsBucket,omitempty"`
	} `json:"connectors"`
	ManagedConnectionName string `json:"managedConnectionName,omitempty"`
}

type Fuse2 struct {
	Connections    []KWConnections
	SQSConnections []SQSConnection
	S3Connections  []S3Connection
	DBConnections  []DBConnection
	Name           string
	JourneySteps   []KWPrototype
	Graphs         []KWPrototype
	JavaScripts    []KWPrototype
	ProjectId      string
}

type KWPrototype struct {
	Id                 string      `json:"id"`
	Type               string      `json:"type"`
	Name               string      `json:"name"`
	Icon               string      `json:"icon"`
	UserId             string      `json:"userId"`
	ProjectId          string      `json:"projectId"`
	ChannelTypeId      *string     `json:"channelTypeId"`
	EditingBy          *string     `json:"editingBy"`
	PrototypeDeletedAt interface{} `json:"prototypeDeletedAt"`
	CreatedAt          time.Time   `json:"createdAt"`
	UpdatedAt          time.Time   `json:"updatedAt"`
	Code               Script
	Graph              interface{}
	Tags               []struct {
		Name      string `json:"name"`
		Color     string `json:"color"`
		IsDefault bool   `json:"isDefault,omitempty"`
		Id        string `json:"id,omitempty"`
	} `json:"tags"`
	IconColor       string `json:"iconColor,omitempty"`
	BackgroundColor string `json:"backgroundColor,omitempty"`
	JourneyStepType string `json:"journeyStepType,omitempty"`
	Outcome         string `json:"outcome,omitempty"`
}

type JavaScript struct {
	Name         string   `json:"name,omitempty"`
	Id           string   `json:"id,omitempty"`
	ProjectId    string   `json:"projectId,omitempty"`
	Code         string   `json:"code,omitempty"`
	ArgumentList []string `json:"argumentList,omitempty"`
	// ArgumentList string `json:"argumentList,omitempty"`
	Tags []string `json:"tags,omitempty"`
	// Tags           string `json:"tags,omitempty"`
	SourceScriptId string `json:"sourceScriptId,omitempty"`
}

type SimpleGraph struct {
	Name              string           `json:"name,omitempty"`
	Id                string           `json:"id,omitempty"`
	ProjectId         string           `json:"projectId,omitempty"`
	YamlSourceGraphId string           `json:"yamlSourceGraphId,omitempty"`
	GraphId           string           `json:"graphId,omitempty"`
	GraphParameters   []GraphParameter `json:"graphParameters"`
}

type GraphParameter struct {
	ParameterDataKeys          []string `json:"parameterDataKeys"`
	ParameterDataReferenceType string   `json:"parameterDataReferenceType"`
	ParameterDataValue         string   `json:"parameterDataValue"`
	ParameterName              string   `json:"parameterName"`
	GraphId                    string   `json:"graphId"`
}

type GraphParameters []GraphParameter

type SimpleNode struct {
	Name                         string         `json:"name,omitempty"`
	Id                           string         `json:"id,omitempty"`
	ProjectId                    string         `json:"projectId,omitempty"`
	NodeType                     string         `json:"nodeType,omitempty"`
	XPosition                    string         `json:"xPosition,omitempty"`
	YPosition                    string         `json:"yPosition,omitempty"`
	GraphId                      string         `json:"graphId,omitempty"`
	SourceGraphId                string         `json:"sourceGraphId,omitempty"`
	SourcePrototypeId            string         `json:"sourcePrototypeId,omitempty"`
	TargetPrototypeId            string         `json:"targetPrototypeId,omitempty"`
	DataSourceReferenceType      string         `json:"dataSourceReferenceType,omitempty"`
	DataSourceKeys               []string       `json:"dataSourceKeys,omitempty"`
	DataDestinationReferenceType string         `json:"dataDestinationReferenceType,omitempty"`
	DataDestinationKeys          []string       `json:"dataDestinationKeys,omitempty"`
	DataSourceValue              string         `json:"dataSourceValue,omitempty"`
	DataDestinationValue         string         `json:"dataDestinationValue,omitempty"`
	StartNode                    bool           `json:"startNode,omitempty"`
	NodeArguments                []NodeArgument `json:"nodeArguments"`
}

type NodeArgument struct {
	ArgDataKeys          []string `json:"argDataKeys"`
	ArgDataReferenceType string   `json:"argDataReferenceType"`
	ArgDataValue         string   `json:"argDataValue"`
	ArgName              string   `json:"argName"`
	NodeId               string   `json:"nodeId"`
}

type GraphLink struct {
	Name                             string `json:"name,omitempty"`
	Id                               string `json:"id,omitempty"`
	YamlSourceLinkId                 string `json:"yamlSourceLinkId"`
	LinkType                         string `json:"linkType,omitempty"`
	GraphId                          string `json:"graphId,omitempty"`
	SourceGraphId                    string `json:"sourceGraphId,omitempty"`
	LinkFromNodeInstanceId           string `json:"linkFromNodeInstanceId,omitempty"`
	YamlSourceLinkFromNodeInstanceId string `json:"yamlSourceLinkFromNodeInstanceId,omitempty"`
	NextNodeInstanceId               string `json:"nextNodeInstanceId,omitempty"`
	YamlSourceNextNodeInstanceId     string `json:"yamlSourceNextNodeInstanceId,omitempty"`
	Condition                        string `json:"condition,omitempty"`
	Description                      string `json:"description,omitempty"`
	Mode                             string `json:"mode,omitempty"`
}

type QueueAdapter struct {
	Name                         string   `json:"name,omitempty"`
	Id                           string   `json:"id,omitempty"`
	ProjectId                    string   `json:"projectId,omitempty"`
	ConnectionId                 string   `json:"connectionId,omitempty"`
	Action                       string   `json:"action,omitempty"`
	DataSourceReferenceType      string   `json:"dataSourceReferenceType,omitempty"`
	DataSourceKeys               []string `json:"dataSourceKeys,omitempty"`
	DataDestinationReferenceType string   `json:"dataDestinationReferenceType,omitempty"`
	DataDestinationKeys          []string `json:"dataDestinationKeys,omitempty"`
	GraphId                      string   `json:"graphId,omitempty"`
	SubGraphId                   string   `json:"subGraphId,omitempty"`
	YamlSourcePrototypeId        string   `json:"yamlSourcePrototypeId,omitempty"`
}

type FileAdapter struct {
	Name                         string   `json:"name,omitempty"`
	Id                           string   `json:"id,omitempty"`
	ProjectId                    string   `json:"projectId,omitempty"`
	ConnectionId                 string   `json:"connectionId,omitempty"`
	Action                       string   `json:"action,omitempty"`
	DataSourceReferenceType      string   `json:"dataSourceReferenceType,omitempty"`
	DataSourceKeys               []string `json:"dataSourceKeys,omitempty"`
	DataDestinationReferenceType string   `json:"dataDestinationReferenceType,omitempty"`
	DataDestinationKeys          []string `json:"dataDestinationKeys,omitempty"`
	FileDataReferenceType        string   `json:"fileDataReferenceType,omitempty"`
	FilenameKeys                 []string `json:"filenameKeys,omitempty"`
	GraphId                      string   `json:"graphId,omitempty"`
	SubGraphId                   string   `json:"subGraphId,omitempty"`
	YamlSourcePrototypeId        string   `json:"yamlSourcePrototypeId,omitempty"`
}

type DatabaseAdapter struct {
	Name                         string   `json:"name,omitempty"`
	Id                           string   `json:"id,omitempty"`
	ProjectId                    string   `json:"projectId,omitempty"`
	ConnectionId                 string   `json:"connectionId,omitempty"`
	Action                       string   `json:"action,omitempty"`
	ActionParameters             []string `json:"actionParameters,omitempty"`
	DataDestinationReferenceType string   `json:"dataDestinationReferenceType,omitempty"`
	DataDestinationKeys          []string `json:"dataDestinationKeys,omitempty"`
	Query                        string   `json:"query,omitempty"`
	QueryParameters              []string `json:"queryParameters,omitempty"`
	GraphId                      string   `json:"graphId,omitempty"`
	SubGraphId                   string   `json:"subGraphId,omitempty"`
	YamlSourcePrototypeId        string   `json:"yamlSourcePrototypeId,omitempty"`
}

type WebServiceAdapter struct {
	Name                            string            `json:"name,omitempty"`
	Id                              string            `json:"id,omitempty"`
	ProjectId                       string            `json:"projectId,omitempty"`
	ConnectionId                    string            `json:"connectionId,omitempty"`
	Method                          string            `json:"method,omitempty"`
	Headers                         map[string]string `json:"headers,omitempty"`
	Parameters                      []GraphParameter  `json:"parameters,omitempty"`
	DataSourceReferenceType         string            `json:"dataSourceReferenceType,omitempty"`
	DataSourceKeys                  []string          `json:"dataSourceKeys,omitempty"`
	DataDestinationReferenceType    string            `json:"dataDestinationReferenceType,omitempty"`
	DataDestinationKeys             []string          `json:"dataDestinationKeys,omitempty"`
	ResponseHeaderDataReferenceType string            `json:"responseHeaderDataReferenceType,omitempty"`
	ResponseHeaderDataKeys          []string          `json:"responseHeaderDataKeys,omitempty"`
	HttpStatusCodeDataReferenceType string            `json:"httpStatusCodeDataReferenceType,omitempty"`
	HttpStatusCodeDataKeys          []string          `json:"httpStatusCodeDataKeys,omitempty"`
	UrlSegment                      string            `json:"urlSegment,omitempty"`
	RawRequest                      bool              `json:"rawRequest"`
	RawResponse                     bool              `json:"rawResponse"`
	GraphId                         string            `json:"graphId,omitempty"`
	SubGraphId                      string            `json:"subGraphId,omitempty"`
	YamlSourcePrototypeId           string            `json:"yamlSourcePrototypeId,omitempty"`
}
