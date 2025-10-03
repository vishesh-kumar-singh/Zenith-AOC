package main

import (
	"database/sql"
	"time"
)

type Documents struct {
	DocumentID string    `json:"document_id"`
	Title      string    `json:"title"`
	Source     string    `json:"source"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	RawText    string    `json:"raw_text"`
}

type DocumentEmbedding struct {
	DocumentID      string    `json:"document_id"`
	ChunkID         string    `json:"chunk_id"`
	EmbeddingVector []float64 `json:"embedding_vector"`
	BM25Index       string    `json:"bm_25_index"`
}

type Customers struct {
	CustomerID     string    `json:"customerID"`
	CompanyName    string    `json:"companyName"`
	Industy        string    `json:"industry"`
	Tier           string    `json:"tier"`
	ContractValue  float64   `json:"contractValue"`
	Location       string    `json:"location"`
	AccountManager string    `json:"accountManager"`
	CreatedAt      time.Time `json:"timestamp"`
	LastModified   time.Time `json:"lastModified"`
}

// support Tickets schema
type SupportTickets struct {
	TicketID          string         `json:"ticket_id"`
	CustomerID        string         `json:"customer_id"`
	Title             string         `json:"title"`
	Description       sql.NullString `json:"description"`
	Category          string         `json:"category"`
	Priority          string         `json:"priority"`
	Status            string         `json:"status"`
	AssignedTo        string         `json:"assigned_to"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	ResolvedAt        sql.NullTime   `json:"resolved_at"`
	SLADeadline       time.Time      `json:"sla_deadline"`
	RelatedEntityType string         `json:"related_entity_type"`
	RelatedEntityID   string         `json:"related_entity_id"`
}

type TicketHistory struct {
	HistoryID    int            `json:"history_id"`
	TicketID     string         `json:"ticket_id"`
	StatusChange string         `json:"status_change"`
	Notes        sql.NullString `json:"notes"`
	ChangedBy    string         `json:"changed_by"`
	ChangedAt    time.Time      `json:"changed_at"`
}

type Machines struct {
	MachineID           string    `json:"machine_id"`
	FacilityID          string    `json:"facility_id"`
	MachineType         string    `json:"machine_type"`
	Model               string    `json:"model"`
	InstallationDate    time.Time `json:"installation_date"`
	LastMaintenanceDate time.Time `json:"last_maintenance_date"`
	Status              string    `json:"status"`
	Location            string    `json:"location"`
}

// IoT sensors Reading
type SensorReadings struct {
	ReadingID    int64     `json:"reading_id"`
	MachineID    string    `json:"machine_id"`
	SensorType   string    `json:"sensor_type"`
	Value        float64   `json:"value"`
	Unit         string    `json:"unit"`
	Timestamp    time.Time `json:"timestamp"`
	AnomalyScore float64   `json:"anomaly_score"`
}

// Shipment Data
type Shipments struct {
	ShipmentID        string    `json:"shipment_id"`
	CustomerID        string    `json:"customer_id"`
	TrackingNumber    string    `json:"tracking_number"`
	Origin            string    `json:"origin"`
	Destination       string    `json:"destination"`
	Status            string    `json:"status"`
	EstimatedDelivery time.Time `json:"estimated_delivery"`
	ActualDelivery    time.Time `json:"actual_delivery"`
	CreatedAt         time.Time `json:"created_at"`
}

type SlackMessages struct {
	MessageID string       `json:"message_id"`
	Channel   string       `json:"channel"`
	UserID    string       `json:"user_id"`
	Text      string       `json:"text"`
	Timestamp sql.NullTime `json:"timestamp"`
	MachineID string       `json:"machine_id"`
}

type SensorAggregates struct {
	MachineID    string    `json:"machine_id"`
	MetricName   string    `json:"metric_name"`
	WindowStart  time.Time `json:"window_start"`
	WindowEnd    time.Time `json:"window_end"`
	RollingAvg   float64   `json:"rolling_average"`
	RollingStd   float64   `json:"rolling_std"`
	AnomalyScore float64   `json:"anomaly_score"`
}

type Anomalies struct {
	Anomaly_ID  string    `json:"anomly_id"`
	MachineID   string    `json:"machine_id"`
	MetricName  string    `json:"metric_name"`
	Timestamp   time.Time `json:"timestamp"`
	Value       float64   `json:"value"`
	ZScore      float64   `json:"z_score"`
	Severity    string    `json:"severity"`
	Description string    `json:"description"`
}

type Incidents struct {
	IncidentID string    `json:"incident_id"`
	Source     string    `json:"source"`
	MachineID  string    `json:"machine_id"`
	Status     string    `json:"status"`
	Priority   string    `json:"priority"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type InvestigationReports struct {
	ReportID       string    `json:"report_id"`
	IncidentID     string    `json:"incident_id"`
	InvestigatorID string    `json:"investigator_id"`
	Hypothesis     string    `json:"hypothesis"`
	Confidence     float64   `json:"confidence"`
	CreatedAt      time.Time `json:"created_at"`
}

type Evidence struct {
	EvidenceID string                 `json:"evidence_id"`
	ReportID   string                 `json:"report_id"`
	Source     string                 `json:"source"`
	Type       string                 `json:"type"`
	Details    map[string]interface{} `json:"details"`
	Timestamp  time.Time              `json:"timestamp"`
}

type ResolutionPlans struct {
	PlanID          string                 `json:"plan_id"`
	IncidentID      string                 `json:"incident_id"`
	ResolutionSteps map[string]interface{} `json:"resolution_steps"`
	Executed        bool                   `json:"executed"`
	CreatedAt       time.Time              `json:"created_at"`
	ExecutedAt      time.Time              `json:"executed_at"`
}

type Audits struct {
	AuditID    string    `json:"audit_id"`
	IncidentID string    `json:"incident_id"`
	AuditorID  string    `json:"auditor_id"`
	Action     string    `json:"action"`
	Comments   string    `json:"comments"`
	Timestamp  time.Time `json:"timestamp"`
}

type PredictiveForecasts struct {
	ForecastID     string    `json:"forecast_id"`
	MachineID      string    `json:"machine_id"`
	MetricName     string    `json:"metric_name"`
	ForecastTime   time.Time `json:"forecast_time"`
	PredictedValue float64   `json:"predicted_value"`
	Confidence     float64   `json:"confidence"`
	CreatedAt      time.Time `json:"created_at"`
}

type ToolCalls struct {
	CallID     string                 `json:"call_id"`
	AgentID    string                 `json:"agent_id"`
	ToolName   string                 `json:"tool_name"`
	Request    map[string]interface{} `json:"request"`
	Response   map[string]interface{} `json:"response"`
	Latencyms  int64                  `json:"latency_ms"`
	TokensUsed int64                  `json:"tokens_used"`
	Timestamp  time.Time              `json:"timestamp"`
}
