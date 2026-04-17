package graph

// ─────────────────────────────────────────────
// Enum types
// ─────────────────────────────────────────────

// NodeStatus is the lifecycle status shared by most nodes.
type NodeStatus string

const (
	NodeStatusActive   NodeStatus = "active"
	NodeStatusArchived NodeStatus = "archived"
)

// PersonType distinguishes full person records from lightweight stubs.
type PersonType string

const (
	PersonTypeFull PersonType = "full"
	PersonTypeStub PersonType = "stub"
)

// ContributorType classifies a person's contribution mode.
type ContributorType string

const (
	ContributorTypeManager               ContributorType = "manager"
	ContributorTypeIndividualContributor ContributorType = "individual_contributor"
	ContributorTypeExternal              ContributorType = "external"
)

// LegalEntityType categorises a legal entity.
type LegalEntityType string

const (
	LegalEntityTypeGroupEntity        LegalEntityType = "group_entity"
	LegalEntityTypeThirdPartySupplier LegalEntityType = "third_party_supplier"
	LegalEntityTypeClientEntity       LegalEntityType = "client_entity"
)

// TeamType classifies a permanent team by its Topology type.
type TeamType string

const (
	TeamTypeStreamAligned        TeamType = "stream_aligned"
	TeamTypePlatform             TeamType = "platform"
	TeamTypeEnabling             TeamType = "enabling"
	TeamTypeComplicatedSubsystem TeamType = "complicated_subsystem"
)

// TemporaryTeamType classifies a temporary or ad-hoc team.
type TemporaryTeamType string

const (
	TemporaryTeamTypeTaskForce    TemporaryTeamType = "task_force"
	TemporaryTeamTypeProjectTeam  TemporaryTeamType = "project_team"
	TemporaryTeamTypeProgramTeam  TemporaryTeamType = "program_team"
	TemporaryTeamTypeWorkingGroup TemporaryTeamType = "working_group"
)

// GovernanceBodyType classifies a governance body.
type GovernanceBodyType string

const (
	GovernanceBodyTypeBoard       GovernanceBodyType = "board"
	GovernanceBodyTypeCommittee   GovernanceBodyType = "committee"
	GovernanceBodyTypeCouncil     GovernanceBodyType = "council"
	GovernanceBodyTypeForum       GovernanceBodyType = "forum"
	GovernanceBodyTypeReviewBoard GovernanceBodyType = "review_board"
)

// Cadence represents how often a governance body convenes.
type Cadence string

const (
	CadenceWeekly      Cadence = "weekly"
	CadenceFortnightly Cadence = "fortnightly"
	CadenceMonthly     Cadence = "monthly"
	CadenceBimonthly   Cadence = "bimonthly"
	CadenceSemiannual  Cadence = "semiannual"
	CadenceAnnual      Cadence = "annual"
	CadenceOnDemand    Cadence = "on_demand"
)

// ServiceType classifies a service.
type ServiceType string

const (
	ServiceTypeTechnical      ServiceType = "technical"
	ServiceTypeBusiness       ServiceType = "business"
	ServiceTypeCustomerFacing ServiceType = "customer_facing"
)

// ServiceStatus is the lifecycle status of a service.
type ServiceStatus string

const (
	ServiceStatusPipeline ServiceStatus = "pipeline"
	ServiceStatusActive   ServiceStatus = "active"
	ServiceStatusArchived ServiceStatus = "archived"
)

// CostCenterType classifies a cost center.
type CostCenterType string

const (
	CostCenterTypeOperational   CostCenterType = "operational"
	CostCenterTypeSharedService CostCenterType = "shared_service"
	CostCenterTypeProject       CostCenterType = "project"
	CostCenterTypeOverhead      CostCenterType = "overhead"
	CostCenterTypeInvestment    CostCenterType = "investment"
)

// ChargebackModel represents the financial recovery model for a cost center.
type ChargebackModel string

const (
	ChargebackModelNone       ChargebackModel = "none"
	ChargebackModelShowback   ChargebackModel = "showback"
	ChargebackModelChargeback ChargebackModel = "chargeback"
	ChargebackModelAllocation ChargebackModel = "allocation"
)

// KPIType distinguishes lagging indicators from leading ones.
type KPIType string

const (
	KPITypeLagging KPIType = "lagging"
	KPITypeLeading KPIType = "leading"
)

// KPICategory classifies KPIs by business domain.
type KPICategory string

const (
	KPICategoryFinancial   KPICategory = "financial"
	KPICategoryOperational KPICategory = "operational"
	KPICategoryCustomer    KPICategory = "customer"
	KPICategoryPeople      KPICategory = "people"
	KPICategoryQuality     KPICategory = "quality"
	KPICategoryRisk        KPICategory = "risk"
)

// KPIDirection indicates how to interpret KPI movement.
type KPIDirection string

const (
	KPIDirectionHigherIsBetter KPIDirection = "higher_is_better"
	KPIDirectionLowerIsBetter  KPIDirection = "lower_is_better"
	KPIDirectionTargetIsBetter KPIDirection = "target_is_better"
)

// TargetOperator qualifies the target comparison when direction is TargetIsBetter.
type TargetOperator string

const (
	TargetOperatorGTE TargetOperator = "gte"
	TargetOperatorLTE TargetOperator = "lte"
	TargetOperatorEQ  TargetOperator = "eq"
)

// MeasurementFrequency defines how often a KPI is measured.
type MeasurementFrequency string

const (
	MeasurementFrequencyDaily     MeasurementFrequency = "daily"
	MeasurementFrequencyWeekly    MeasurementFrequency = "weekly"
	MeasurementFrequencyMonthly   MeasurementFrequency = "monthly"
	MeasurementFrequencyQuarterly MeasurementFrequency = "quarterly"
	MeasurementFrequencyAnnual    MeasurementFrequency = "annual"
)

// OKRNodeType distinguishes an Objective from a Key Result.
type OKRNodeType string

const (
	OKRNodeTypeObjective OKRNodeType = "objective"
	OKRNodeTypeKeyResult OKRNodeType = "key_result"
)

// OKRProgressStatus tracks the delivery status of an OKR.
type OKRProgressStatus string

const (
	OKRProgressStatusNotStarted OKRProgressStatus = "not_started"
	OKRProgressStatusOnTrack    OKRProgressStatus = "on_track"
	OKRProgressStatusAtRisk     OKRProgressStatus = "at_risk"
	OKRProgressStatusOffTrack   OKRProgressStatus = "off_track"
	OKRProgressStatusAchieved   OKRProgressStatus = "achieved"
	OKRProgressStatusMissed     OKRProgressStatus = "missed"
)

// OKRStatus is the lifecycle status of an OKR.
type OKRStatus string

const (
	OKRStatusDraft  OKRStatus = "draft"
	OKRStatusActive OKRStatus = "active"
	OKRStatusClosed OKRStatus = "closed"
)

// ProcessStatus is the lifecycle status of a process.
type ProcessStatus string

const (
	ProcessStatusPipeline ProcessStatus = "pipeline"
	ProcessStatusActive   ProcessStatus = "active"
	ProcessStatusArchived ProcessStatus = "archived"
)

// AuthorityType classifies a decision authority.
type AuthorityType string

const (
	AuthorityTypeDecides  AuthorityType = "decides"
	AuthorityTypeApproves AuthorityType = "approves"
	AuthorityTypeAdvises  AuthorityType = "advises"
	AuthorityTypeRatifies AuthorityType = "ratifies"
)

// MemberType classifies a governance body membership role.
type MemberType string

const (
	MemberTypeChair       MemberType = "chair"
	MemberTypeFacilitator MemberType = "facilitator"
	MemberTypeStanding    MemberType = "standing"
	MemberTypeAdvisory    MemberType = "advisory"
)

// ContributionType classifies how a node contributes to a KPI.
type ContributionType string

const (
	ContributionTypeOwns        ContributionType = "owns"
	ContributionTypeContributes ContributionType = "contributes"
	ContributionTypeInfluences  ContributionType = "influences"
)

// AllocationKey defines the basis for cost allocation in chargeback relationships.
type AllocationKey string

const (
	AllocationKeyHeadcount AllocationKey = "headcount"
	AllocationKeyUsage     AllocationKey = "usage"
	AllocationKeyFixed     AllocationKey = "fixed"
	AllocationKeyCustom    AllocationKey = "custom"
)

// CapabilityLevel classifies a business capability by its granularity level in the hierarchy.
type CapabilityLevel string

const (
	CapabilityLevelL1 CapabilityLevel = "l1"
	CapabilityLevelL2 CapabilityLevel = "l2"
	CapabilityLevelL3 CapabilityLevel = "l3"
)

// ─────────────────────────────────────────────
// Shared / helper types
// ─────────────────────────────────────────────

// GeographicalLocation represents a geographical location with city, country, and timezone.
type GeographicalLocation struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	Timezone string `json:"timezone"` // IANA timezone e.g. "Europe/Amsterdam"
}

// HistoryEntry is a single entry in a node's timestamped change history.
type HistoryEntry struct {
	Timestamp     string   `json:"timestamp"`     // ISO 8601
	UpdatedFields []string `json:"updatedFields"`
}

// SocialLink is a link to a personal social or professional profile.
type SocialLink struct {
	Platform string `json:"platform"` // e.g. "LinkedIn", "GitHub", "X"
	URL      string `json:"url"`
}

// PersonalityProfile is a link to a personality or strengths profile.
type PersonalityProfile struct {
	Type string `json:"type"` // e.g. "CliftonStrengths", "MBTI", "DiSC"
	URL  string `json:"url"`
}

// CommunicationChannel is a named communication channel link.
type CommunicationChannel struct {
	Label string `json:"label"` // e.g. "Slack Channel", "MS Teams Channel"
	URL   string `json:"url"`   // mailto:team@example.com or https://slack.com/...
}

// WorkspaceLink is a named URL link to an external workspace or tool.
type WorkspaceLink struct {
	Label string `json:"label"` // e.g. "Jira Board", "Confluence Space"
	URL   string `json:"url"`
}

// DecisionAuthority is an inline decision authority attribute on a node.
type DecisionAuthority struct {
	AuthorityType AuthorityType `json:"authorityType"`
	Description   string        `json:"description"`
	DomainID      string        `json:"domainId"`
}

// ─────────────────────────────────────────────
// Node types
// ─────────────────────────────────────────────

// DomainNode represents a domain that governance bodies and decision authorities operate within.
type DomainNode struct {
	ID        string         `json:"id"` // format: "domain-[slug]"
	Name      string         `json:"name"`
	Status    NodeStatus     `json:"status"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	History   []HistoryEntry `json:"history"`
}

// LegalEntityNode represents a legal entity within the organisational structure.
type LegalEntityNode struct {
	ID        string               `json:"id"` // format: "legal-entity-[slug]"
	Name      string               `json:"name"`
	Type      LegalEntityType      `json:"type"`
	Address   string               `json:"address"`
	Location  GeographicalLocation `json:"location"`
	Status    NodeStatus           `json:"status"`
	CreatedAt string               `json:"createdAt"`
	UpdatedAt string               `json:"updatedAt"`
	History   []HistoryEntry       `json:"history"`
}

// TeamNode represents a permanent team node in the organisational hierarchy.
type TeamNode struct {
	ID                        string                 `json:"id"` // format: "team-[slug]"
	Name                      string                 `json:"name"`
	Type                      TeamType               `json:"type"`
	Purpose                   string                 `json:"purpose"`
	Responsibilities          []string               `json:"responsibilities"`
	DecisionAuthorities       []DecisionAuthority    `json:"decisionAuthorities"`
	DirectMemberCount         int                    `json:"directMemberCount"`
	TotalMemberCount          int                    `json:"totalMemberCount"`
	DirectInternalMemberCount int                    `json:"directInternalMemberCount"`
	DirectExternalMemberCount int                    `json:"directExternalMemberCount"`
	WorkspaceLinks            []WorkspaceLink        `json:"workspaceLinks"`
	CommunicationChannels     []CommunicationChannel `json:"communicationChannels"`
	IsManagedService          bool                   `json:"isManagedService"`
	Status                    NodeStatus             `json:"status"`
	CreatedAt                 string                 `json:"createdAt"`
	UpdatedAt                 string                 `json:"updatedAt"`
	History                   []HistoryEntry         `json:"history"`
}

// TemporaryTeamNode represents a temporary or ad-hoc team within the hierarchy.
type TemporaryTeamNode struct {
	ID                        string                 `json:"id"` // format: "temp-team-[slug]"
	Name                      string                 `json:"name"`
	Type                      TemporaryTeamType      `json:"type"`
	Purpose                   string                 `json:"purpose"`
	Responsibilities          []string               `json:"responsibilities"`
	DecisionAuthorities       []DecisionAuthority    `json:"decisionAuthorities"`
	WorkspaceLinks            []WorkspaceLink        `json:"workspaceLinks"`
	CommunicationChannels     []CommunicationChannel `json:"communicationChannels"`
	StartDate                 string                 `json:"startDate"`     // ISO 8601 date
	TargetEndDate             *string                `json:"targetEndDate"` // null if open-ended
	DirectMemberCount         int                    `json:"directMemberCount"`
	DirectInternalMemberCount int                    `json:"directInternalMemberCount"`
	DirectExternalMemberCount int                    `json:"directExternalMemberCount"`
	Status                    NodeStatus             `json:"status"`
	CreatedAt                 string                 `json:"createdAt"`
	UpdatedAt                 string                 `json:"updatedAt"`
	History                   []HistoryEntry         `json:"history"`
}

// MasterRoleNode represents a canonical role template shared across teams.
type MasterRoleNode struct {
	ID               string         `json:"id"` // format: "master-role-[slug]"
	Name             string         `json:"name"`
	Purpose          string         `json:"purpose"`
	Responsibilities []string       `json:"responsibilities"`
	Status           NodeStatus     `json:"status"`
	CreatedAt        string         `json:"createdAt"`
	UpdatedAt        string         `json:"updatedAt"`
	History          []HistoryEntry `json:"history"`
}

// RoleNode represents a role instance within a specific team.
type RoleNode struct {
	ID                    string                 `json:"id"` // format: "role-[master-role-slug]-[team-slug]"
	Name                  string                 `json:"name"`
	Purpose               string                 `json:"purpose"`
	WorkspaceLinks        []WorkspaceLink        `json:"workspaceLinks"`
	CommunicationChannels []CommunicationChannel `json:"communicationChannels"`
	Responsibilities      []string               `json:"responsibilities"`
	DecisionAuthorities   []DecisionAuthority    `json:"decisionAuthorities"`
	Status                NodeStatus             `json:"status"`
	CreatedAt             string                 `json:"createdAt"`
	UpdatedAt             string                 `json:"updatedAt"`
	History               []HistoryEntry         `json:"history"`
}

// PersonNode represents a person within the organisation.
type PersonNode struct {
	ID                    string                 `json:"id"` // format: "person-[slug]"
	PersonType            PersonType             `json:"personType"`
	FirstName             string                 `json:"firstName"`
	Surname               string                 `json:"surname"`
	Email                 *string                `json:"email"`    // null for stub persons
	IsExternal            bool                   `json:"isExternal"`
	CommunicationChannels []CommunicationChannel `json:"communicationChannels"`
	ContributorType       ContributorType        `json:"contributorType"`
	Pronouns              *string                `json:"pronouns"`
	AboutMe               *string                `json:"aboutMe"`
	PersonalityProfiles   []PersonalityProfile   `json:"personalityProfiles"`
	SocialLinks           []SocialLink           `json:"socialLinks"`
	AvatarURL             *string                `json:"avatarUrl"`
	Location              GeographicalLocation   `json:"location"`
	Status                NodeStatus             `json:"status"`
	CreatedAt             string                 `json:"createdAt"`
	UpdatedAt             string                 `json:"updatedAt"`
	History               []HistoryEntry         `json:"history"`
}

// ValueStreamNode represents an end-to-end value stream crossing team boundaries.
type ValueStreamNode struct {
	ID               string         `json:"id"` // format: "value-stream-[slug]"
	Name             string         `json:"name"`
	BusinessOutcomes []string       `json:"businessOutcomes"`
	Status           NodeStatus     `json:"status"`
	CreatedAt        string         `json:"createdAt"`
	UpdatedAt        string         `json:"updatedAt"`
	History          []HistoryEntry `json:"history"`
}

// GovernanceBodyNode represents a governance body.
type GovernanceBodyNode struct {
	ID                    string                 `json:"id"` // format: "governance-body-[slug]"
	Name                  string                 `json:"name"`
	Purpose               string                 `json:"purpose"`
	Type                  GovernanceBodyType     `json:"type"`
	Cadence               Cadence                `json:"cadence"`
	Inputs                []string               `json:"inputs"`
	Outputs               []string               `json:"outputs"`
	DecisionRecordURL     *string                `json:"decisionRecordUrl"`
	WorkspaceLinks        []WorkspaceLink        `json:"workspaceLinks"`
	CommunicationChannels []CommunicationChannel `json:"communicationChannels"`
	DecisionAuthorities   []DecisionAuthority    `json:"decisionAuthorities"`
	Status                NodeStatus             `json:"status"`
	CreatedAt             string                 `json:"createdAt"`
	UpdatedAt             string                 `json:"updatedAt"`
	History               []HistoryEntry         `json:"history"`
}

// ServiceNode represents a service provided or consumed within the organisational operating model.
type ServiceNode struct {
	ID                    string                 `json:"id"` // format: "service-[slug]"
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	Type                  ServiceType            `json:"type"`
	ServiceDomain         string                 `json:"serviceDomain"`
	SLA                   *string                `json:"sla"`
	RequestURL            *string                `json:"requestUrl"`
	CommunicationChannels []CommunicationChannel `json:"communicationChannels"`
	Status                ServiceStatus          `json:"status"`
	CreatedAt             string                 `json:"createdAt"`
	CreatedBy             string                 `json:"createdBy"` // PersonNode ID
	UpdatedAt             string                 `json:"updatedAt"`
	UpdatedBy             string                 `json:"updatedBy"` // PersonNode ID
	History               []HistoryEntry         `json:"history"`
}

// CostCenterNode represents a cost center used for financial tracking and budget allocation.
type CostCenterNode struct {
	ID              string          `json:"id"` // format: "cost-center-[slug]"
	Name            string          `json:"name"`
	Code            string          `json:"code"`            // ERP/GL code e.g. "CC-4720"
	ExternalID      *string         `json:"externalId"`      // for SAP/Oracle bi-directional sync
	Type            CostCenterType  `json:"type"`
	Currency        string          `json:"currency"`        // ISO 4217 e.g. "EUR"
	AnnualBudget    *float64        `json:"annualBudget"`
	FiscalYearStart *string         `json:"fiscalYearStart"` // ISO 8601 date
	ChargebackModel ChargebackModel `json:"chargebackModel"`
	Status          NodeStatus      `json:"status"`
	CreatedAt       string          `json:"createdAt"`
	UpdatedAt       string          `json:"updatedAt"`
	History         []HistoryEntry  `json:"history"`
}

// KPINode represents a measurable key performance indicator.
type KPINode struct {
	ID                   string               `json:"id"` // format: "kpi-[slug]"
	Name                 string               `json:"name"`
	Abbreviation         *string              `json:"abbreviation"` // e.g. "MTTR", "NPS"
	Description          string               `json:"description"`
	Type                 KPIType              `json:"type"`
	Category             KPICategory          `json:"category"`
	Unit                 string               `json:"unit"` // e.g. "hours", "EUR", "%"
	Direction            KPIDirection         `json:"direction"`
	TargetValue          *float64             `json:"targetValue"`
	TargetOperator       *TargetOperator      `json:"targetOperator"` // set when direction is target_is_better
	MeasurementFrequency MeasurementFrequency `json:"measurementFrequency"`
	DashboardURL         *string              `json:"dashboardUrl"`
	Status               NodeStatus           `json:"status"`
	CreatedAt            string               `json:"createdAt"`
	UpdatedAt            string               `json:"updatedAt"`
	History              []HistoryEntry       `json:"history"`
}

// OKRNode represents an Objective or Key Result within an OKR cycle.
type OKRNode struct {
	ID              string            `json:"id"` // format: "okr-[slug]"
	NodeType        OKRNodeType       `json:"nodeType"`
	Title           string            `json:"title"`
	Description     *string           `json:"description"`
	Period          string            `json:"period"`       // e.g. "Q2-2025", "FY2025"
	StartDate       string            `json:"startDate"`    // ISO 8601
	EndDate         string            `json:"endDate"`      // ISO 8601
	TargetValue     *float64          `json:"targetValue"`  // null for objectives
	CurrentValue    *float64          `json:"currentValue"` // null for objectives
	Unit            *string           `json:"unit"`         // null for objectives
	ProgressStatus  OKRProgressStatus `json:"progressStatus"`
	ConfidenceScore *float64          `json:"confidenceScore"` // 0–100
	Status          OKRStatus         `json:"status"`
	CreatedAt       string            `json:"createdAt"`
	UpdatedAt       string            `json:"updatedAt"`
	History         []HistoryEntry    `json:"history"`
}

// ProcessNode represents a standardised repeatable workflow.
type ProcessNode struct {
	ID               string         `json:"id"` // format: "process-[slug]"
	Name             string         `json:"name"`
	Type             string         `json:"type"`
	Description      string         `json:"description"`
	Steps            []string       `json:"steps"`
	DocumentationURL *string        `json:"documentationUrl"`
	Status           ProcessStatus  `json:"status"`
	CreatedAt        string         `json:"createdAt"`
	UpdatedAt        string         `json:"updatedAt"`
	History          []HistoryEntry `json:"history"`
}

// BusinessCapabilityNode represents a business capability —
// what the organisation can do, independent of how it is structured.
// L1/L2/L3 granularity is expressed via ParentCapabilityID hierarchy,
// not separate node types. Maximum depth is L3.
type BusinessCapabilityNode struct {
	ID                 string         `json:"id"`
	// format: "capability-[slug]"
	// e.g. "capability-customer-onboarding"

	Name               string         `json:"name"`
	Description        string         `json:"description"`

	ParentCapabilityID *string        `json:"parentCapabilityId"`
	// nil  = L1 top-level capability
	// one level deep  = L2
	// two levels deep = L3
	// do not model deeper than L3

	Status             NodeStatus     `json:"status"`
	CreatedAt          string         `json:"createdAt"`
	UpdatedAt          string         `json:"updatedAt"`
	History            []HistoryEntry `json:"history"`
}

// CommunityNode represents a community of practice or interest
// that cuts across team boundaries. Membership is expressed via
// edges — either from specific persons (explicit) or from master
// roles whose role instances and assigned persons automatically
// become members (implicit).
type CommunityNode struct {
	ID        string         `json:"id"`
	// format: "community-[slug]"
	// e.g. "community-frontend-engineers"

	Name      string         `json:"name"`
	Purpose   string         `json:"purpose"`
	Status    NodeStatus     `json:"status"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	History   []HistoryEntry `json:"history"`
}

// ─────────────────────────────────────────────
// Edge types
// ─────────────────────────────────────────────

// LegalEntityParentOfEdge — LegalEntity → LegalEntity: subsidiary relationship.
type LegalEntityParentOfEdge struct {
	ID               string `json:"id"` // format: "edge-[source-slug]-legal_entity_parent_of-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "legal_entity_parent_of"
	Source           string `json:"source"`
	Target           string `json:"target"`
}

// LegalRepresentativeEdge — Person → LegalEntity: person is the legal representative.
type LegalRepresentativeEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "is_legal_representative_of"
	Source           string `json:"source"`           // PersonNode ID
	Target           string `json:"target"`           // LegalEntityNode ID
}

// BelongsToLegalEntityEdge — Person | GovernanceBody | CostCenter → LegalEntity.
type BelongsToLegalEntityEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "belongs_to_legal_entity"
	Source           string `json:"source"`           // PersonNode | GovernanceBodyNode | CostCenterNode ID
	Target           string `json:"target"`           // LegalEntityNode ID
}

// ManagedByLegalEntityEdge — Team → LegalEntity: team is operated as a managed service.
type ManagedByLegalEntityEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "managed_by_legal_entity"
	Source           string `json:"source"`           // TeamNode ID
	Target           string `json:"target"`           // LegalEntityNode ID
}

// SponsorsEdge — Role → TemporaryTeam: role is the sponsor.
type SponsorsEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "sponsors"
	Source           string `json:"source"`           // RoleNode ID
	Target           string `json:"target"`           // TemporaryTeamNode ID
}

// ParentOfEdge — Team → Team: permanent team hierarchy.
type ParentOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "parent_of"
	Source           string `json:"source"`           // TeamNode ID
	Target           string `json:"target"`           // TeamNode ID
}

// TempTeamParentOfEdge — TemporaryTeam → Team | TemporaryTeam: temporary team parent.
type TempTeamParentOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "temp_team_parent_of"
	Source           string `json:"source"`           // TemporaryTeamNode ID
	Target           string `json:"target"`           // TeamNode | TemporaryTeamNode ID
}

// HasRoleEdge — Team | TemporaryTeam → Role: role exists within a team.
type HasRoleEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "has_role"
	Source           string `json:"source"`           // TeamNode | TemporaryTeamNode ID
	Target           string `json:"target"`           // RoleNode ID
}

// InstanceOfEdge — Role → MasterRole: role is an instance of a master role template.
type InstanceOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "instance_of"
	Source           string `json:"source"`           // RoleNode ID
	Target           string `json:"target"`           // MasterRoleNode ID
}

// SupervisesEdge — Person → Person: direct supervision relationship.
type SupervisesEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "supervises"
	Source           string `json:"source"`           // supervisor PersonNode ID
	Target           string `json:"target"`           // report PersonNode ID
}

// FillsRoleEdge — Person → Role: person currently fills this role.
type FillsRoleEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "fills_role"
	Source           string `json:"source"`           // PersonNode ID
	Target           string `json:"target"`           // RoleNode ID
}

// LeadOfEdge — Person → Team | TemporaryTeam: person is the lead.
type LeadOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "lead_of"
	Source           string `json:"source"`           // PersonNode ID
	Target           string `json:"target"`           // TeamNode | TemporaryTeamNode ID
}

// MemberOfEdge — Person → Team | TemporaryTeam: person is a member.
type MemberOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "member_of"
	Source           string `json:"source"`           // PersonNode ID
	Target           string `json:"target"`           // TeamNode | TemporaryTeamNode ID
}

// ContributesToEdge — Team | TemporaryTeam | Role | Service | Process → ValueStream.
type ContributesToEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "contributes_to"
	Source           string `json:"source"`           // TeamNode | TemporaryTeamNode | RoleNode | ServiceNode | ProcessNode ID
	Target           string `json:"target"`           // ValueStreamNode ID
}

// OwnsServiceEdge — Role → Service: role is the owner.
type OwnsServiceEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "owns_service"
	Source           string `json:"source"`           // RoleNode ID
	Target           string `json:"target"`           // ServiceNode ID
}

// ProvidesServiceEdge — Team | TemporaryTeam → Service: team is the provider.
type ProvidesServiceEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "provides_service"
	Source           string `json:"source"`           // TeamNode | TemporaryTeamNode ID
	Target           string `json:"target"`           // ServiceNode ID
}

// ConsumesServiceEdge — Team | TemporaryTeam → Service: team consumes this service.
type ConsumesServiceEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "consumes_service"
	Source           string `json:"source"`           // TeamNode | TemporaryTeamNode ID
	Target           string `json:"target"`           // ServiceNode ID
}

// LegalEntityProvidesServiceEdge — LegalEntity → Service: legal entity is the provider.
type LegalEntityProvidesServiceEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "legal_entity_provides_service"
	Source           string `json:"source"`           // LegalEntityNode ID
	Target           string `json:"target"`           // ServiceNode ID
}

// ServiceDependsOnEdge — Service → Service: this service depends on another.
type ServiceDependsOnEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "depends_on"
	Source           string `json:"source"`           // ServiceNode ID
	Target           string `json:"target"`           // ServiceNode ID
}

// GovernanceMemberEdge — Person | Role → GovernanceBody: membership with a role type.
type GovernanceMemberEdge struct {
	ID               string     `json:"id"`
	RelationshipType string     `json:"relationshipType"` // "governance_member"
	MemberType       MemberType `json:"memberType"`
	Source           string     `json:"source"` // PersonNode | RoleNode ID
	Target           string     `json:"target"` // GovernanceBodyNode ID
}

// GovernsEdge — GovernanceBody → Team | TemporaryTeam: governs this team.
type GovernsEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "governs"
	Source           string `json:"source"`           // GovernanceBodyNode ID
	Target           string `json:"target"`           // TeamNode | TemporaryTeamNode ID
}

// GovernsServiceEdge — GovernanceBody → Service: governs this service.
type GovernsServiceEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "governs_service"
	Source           string `json:"source"`           // GovernanceBodyNode ID
	Target           string `json:"target"`           // ServiceNode ID
}

// EscalatesToEdge — GovernanceBody → GovernanceBody: escalates unresolved decisions.
type EscalatesToEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "escalates_to"
	Source           string `json:"source"`           // GovernanceBodyNode ID
	Target           string `json:"target"`           // GovernanceBodyNode ID
}

// DelegatesToEdge — GovernanceBody → GovernanceBody | Role: delegates authority.
type DelegatesToEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "delegates_to"
	Source           string `json:"source"`           // GovernanceBodyNode ID
	Target           string `json:"target"`           // GovernanceBodyNode | RoleNode ID
}

// ContributingTeamEdge — Team | TemporaryTeam → TemporaryTeam: contributes to a temporary team.
type ContributingTeamEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "contributing_team"
	Source           string `json:"source"`           // TeamNode | TemporaryTeamNode ID
	Target           string `json:"target"`           // TemporaryTeamNode ID
}

// ValueStreamParentOfProcessEdge — ValueStream → Process: value stream is the top-level parent.
type ValueStreamParentOfProcessEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "value_stream_parent_of_process"
	Source           string `json:"source"`           // ValueStreamNode ID
	Target           string `json:"target"`           // ProcessNode ID
}

// BusinessCapabilityParentOfProcessEdge — BusinessCapability → Process:
// a business capability decomposes into this process.
// Distinct from ValueStreamParentOfProcessEdge — expresses how a
// capability is executed, not how a value stream is delivered.
type BusinessCapabilityParentOfProcessEdge struct {
	ID               string `json:"id"`
	// format: "edge-[source-slug]-capability_parent_of_process-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "capability_parent_of_process"
	Source           string `json:"source"`           // BusinessCapabilityNode ID
	Target           string `json:"target"`           // ProcessNode ID
}

// ProcessParentOfEdge — Process → Process: parent process decomposes into this child.
type ProcessParentOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "process_parent_of"
	Source           string `json:"source"`           // ProcessNode ID
	Target           string `json:"target"`           // ProcessNode ID
}

// ExecutesProcessEdge — Role | Team | GovernanceBody → Process: node executes this process.
type ExecutesProcessEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "executes_process"
	Source           string `json:"source"`           // RoleNode | TeamNode | GovernanceBodyNode ID
	Target           string `json:"target"`           // ProcessNode ID
}

// OwnsProcessEdge — Role → Process: role is the owner.
type OwnsProcessEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "owns_process"
	Source           string `json:"source"`           // RoleNode ID
	Target           string `json:"target"`           // ProcessNode ID
}

// OwnsCostCenterEdge — Role → CostCenter: role is the owner.
type OwnsCostCenterEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "owns_cost_center"
	Source           string `json:"source"`           // RoleNode ID
	Target           string `json:"target"`           // CostCenterNode ID
}

// BookedToCostCenterEdge — Team | TemporaryTeam | Role | Service → CostCenter: expenditure booking.
type BookedToCostCenterEdge struct {
	ID                string   `json:"id"`
	RelationshipType  string   `json:"relationshipType"`  // "booked_to_cost_center"
	AllocationPercent *float64 `json:"allocationPercent"` // 0–100; null implies 100%
	Source            string   `json:"source"`            // TeamNode | TemporaryTeamNode | RoleNode | ServiceNode ID
	Target            string   `json:"target"`            // CostCenterNode ID
}

// CostCenterParentOfEdge — CostCenter → CostCenter: hierarchical parent relationship.
type CostCenterParentOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "cost_center_parent_of"
	Source           string `json:"source"`           // CostCenterNode ID
	Target           string `json:"target"`           // CostCenterNode ID
}

// ChargesBackToEdge — CostCenter → CostCenter: charges back costs to another cost center.
type ChargesBackToEdge struct {
	ID               string         `json:"id"`
	RelationshipType string         `json:"relationshipType"` // "charges_back_to"
	AllocationKey    *AllocationKey `json:"allocationKey"`
	Source           string         `json:"source"` // CostCenterNode ID
	Target           string         `json:"target"` // CostCenterNode ID
}

// KPIParentOfEdge — KPI → KPI: parent KPI decomposes into this child KPI.
type KPIParentOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "kpi_parent_of"
	Source           string `json:"source"`           // KPINode ID
	Target           string `json:"target"`           // KPINode ID
}

// ContributesToKPIEdge — Team | TemporaryTeam | Role | Service | ValueStream | Person | Process | BusinessCapability → KPI.
type ContributesToKPIEdge struct {
	ID               string           `json:"id"`
	RelationshipType string           `json:"relationshipType"` // "contributes_to_kpi"
	ContributionType ContributionType `json:"contributionType"`
	Source           string           `json:"source"` // TeamNode | TemporaryTeamNode | RoleNode | ServiceNode | ValueStreamNode | PersonNode | ProcessNode | BusinessCapabilityNode ID
	Target           string           `json:"target"` // KPINode ID
}

// KPIDependsOnEdge — KPI → KPI: causal dependency (leading indicator relationship).
type KPIDependsOnEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "kpi_depends_on"
	Source           string `json:"source"`           // KPINode ID
	Target           string `json:"target"`           // KPINode ID
}

// OKRParentOfEdge — Objective → KeyResult: objective decomposes into this key result.
type OKRParentOfEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "okr_parent_of"
	Source           string `json:"source"`           // OKRNode ID (objective)
	Target           string `json:"target"`           // OKRNode ID (key_result)
}

// AlignedToOKREdge — OKR → OKR: this OKR supports a higher-level OKR (cross-ownership cascade).
type AlignedToOKREdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "aligned_to_okr"
	Source           string `json:"source"`           // child OKRNode ID
	Target           string `json:"target"`           // parent OKRNode ID
}

// OwnsOKREdge — Team | TemporaryTeam | Role | Person → OKR: accountable owner.
type OwnsOKREdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "owns_okr"
	Source           string `json:"source"`           // TeamNode | TemporaryTeamNode | RoleNode | PersonNode ID
	Target           string `json:"target"`           // OKRNode ID
}

// MeasuredByKPIEdge — KeyResult → KPI: key result is measured by this KPI.
type MeasuredByKPIEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "measured_by_kpi"
	Source           string `json:"source"`           // OKRNode ID (must be key_result)
	Target           string `json:"target"`           // KPINode ID
}

// BelongsToDomainEdge — Node → Domain: node belongs to a domain.
type BelongsToDomainEdge struct {
	ID               string `json:"id"`
	RelationshipType string `json:"relationshipType"` // "belongs_to_domain"
	Source           string `json:"source"`           // RoleNode | MasterRoleNode | TeamNode | TemporaryTeamNode | GovernanceBodyNode | ServiceNode | CostCenterNode | ProcessNode | KPINode | OKRNode ID
	Target           string `json:"target"`           // DomainNode ID
}

// CapabilityParentOfEdge — BusinessCapability → BusinessCapability:
// parent capability decomposes into this child capability.
// Mirrors ParentCapabilityID for graph traversal consistency.
type CapabilityParentOfEdge struct {
	ID               string `json:"id"`
	// format: "edge-[source-slug]-capability_parent_of-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "capability_parent_of"
	Source           string `json:"source"`           // BusinessCapabilityNode ID
	Target           string `json:"target"`           // BusinessCapabilityNode ID
}

// EnablesValueStreamEdge — BusinessCapability → ValueStream:
// this capability enables one or more value streams.
type EnablesValueStreamEdge struct {
	ID               string `json:"id"`
	// format: "edge-[source-slug]-enables_value_stream-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "enables_value_stream"
	Source           string `json:"source"`           // BusinessCapabilityNode ID
	Target           string `json:"target"`           // ValueStreamNode ID
}

// OwnsCapabilityEdge — Role | Team → BusinessCapability:
// node is the accountable owner of this capability.
type OwnsCapabilityEdge struct {
	ID               string `json:"id"`
	// format: "edge-[source-slug]-owns_capability-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "owns_capability"
	Source           string `json:"source"`           // RoleNode | TeamNode ID
	Target           string `json:"target"`           // BusinessCapabilityNode ID
}

// CommunityLeadEdge — Person → Community: person is the lead
// of this community.
type CommunityLeadEdge struct {
	ID               string `json:"id"`
	// format: "edge-[source-slug]-community_lead-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "community_lead"
	Source           string `json:"source"`           // PersonNode ID
	Target           string `json:"target"`           // CommunityNode ID
}

// PersonCommunityMemberEdge — Person → Community: person is an
// explicit member of this community.
type PersonCommunityMemberEdge struct {
	ID               string `json:"id"`
	// format: "edge-[source-slug]-person_community_member-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "person_community_member"
	Source           string `json:"source"`           // PersonNode ID
	Target           string `json:"target"`           // CommunityNode ID
}

// MasterRoleCommunityMemberEdge — MasterRole → Community: all persons
// filling any role that is an instance of this master role
// automatically become members of this community.
// Resolved at query time via three-hop traversal:
// MasterRole → (InstanceOfEdge) → Role → (FillsRoleEdge) → Person.
type MasterRoleCommunityMemberEdge struct {
	ID               string `json:"id"`
	// format: "edge-[source-slug]-master_role_community_member-[target-slug]"
	RelationshipType string `json:"relationshipType"` // "master_role_community_member"
	Source           string `json:"source"`           // MasterRoleNode ID
	Target           string `json:"target"`           // CommunityNode ID
}

// ─────────────────────────────────────────────
// Graph union and root types
// ─────────────────────────────────────────────

// GraphEdge is the union of all directed edge types in the graph.
// Use a type switch to discriminate between concrete edge types at runtime.
type GraphEdge interface{}

// GraphData is the complete graph data model.
type GraphData struct {
	Teams            []TeamNode           `json:"teams"`
	TemporaryTeams   []TemporaryTeamNode  `json:"temporaryTeams"`
	LegalEntities    []LegalEntityNode    `json:"legalEntities"`
	Persons          []PersonNode         `json:"persons"`
	MasterRoles      []MasterRoleNode     `json:"masterRoles"`
	Roles            []RoleNode           `json:"roles"`
	ValueStreams      []ValueStreamNode    `json:"valueStreams"`
	Domains          []DomainNode         `json:"domains"`
	GovernanceBodies []GovernanceBodyNode `json:"governanceBodies"`
	Services         []ServiceNode        `json:"services"`
	CostCenters      []CostCenterNode     `json:"costCenters"`
	KPIs             []KPINode            `json:"kpis"`
	OKRs             []OKRNode            `json:"okrs"`
	Processes            []ProcessNode            `json:"processes"`
	BusinessCapabilities []BusinessCapabilityNode `json:"businessCapabilities"`
	Communities          []CommunityNode          `json:"communities"`
	Edges                []GraphEdge              `json:"edges"`
}
