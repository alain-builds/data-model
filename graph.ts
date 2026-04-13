/** Represents a geographical location with city, country, and timezone. */
export interface GeographicalLocation {
  city: string;
  country: string;
  timezone: string; // IANA timezone e.g. "Europe/Amsterdam"
}

/** A single entry in a node's timestamped change history. */
export interface HistoryEntry {
  timestamp: string; // ISO 8601
  updatedFields: string[];
}

/** A link to a personal social or professional profile (e.g. LinkedIn, Workday, GitHub, X). */
export interface SocialLink {
  platform: string; // e.g. "LinkedIn", "Workday", "GitHub", "X", "Instagram"
  url: string;
}

/** A link to a personality or strengths profile (e.g. CliftonStrengths, MBTI, DiSC). */
export interface PersonalityProfile {
  type: string; // e.g. "CliftonStrengths", "MBTI", "DiSC", "Enneagram"
  url: string;
}

/** A named communication channel link (e.g. Slack, MS Teams, email). Use mailto: prefix for email addresses. */
export interface CommunicationChannel {
  label: string; // e.g. "Slack Channel", "Team Email", "MS Teams Channel"
  url: string; // mailto:team@example.com or https://slack.com/...
}

/** A named URL link to an external workspace, tool, or space (e.g. Jira, Confluence, Asana). */
export interface WorkspaceLink {
  label: string; // e.g. "Jira Board", "Confluence Space", "Asana Project"
  url: string;
}

/** Represents a domain that governance bodies and decision authorities operate within. */
export interface DomainNode {
  id: string; // format: "domain-[slug]"
  name: string;
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents an inline decision authority attribute on a node. */
export interface DecisionAuthority {
  authorityType: 'decides' | 'approves' | 'advises' | 'ratifies';
  description: string;
  domainId: DomainNode['id'];
}

/** Represents a legal entity within the organizational structure. */
export interface LegalEntityNode {
  id: string; // format: "legal-entity-[slug]"
  name: string;
  type: 'group_entity' | 'third_party_supplier' | 'client_entity';
  address: string;
  location: GeographicalLocation;
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a permanent team node in the organizational hierarchy. */
export interface TeamNode {
  id: string; // format: "team-[slug]"
  name: string;
  type: 'stream_aligned' | 'platform' | 'enabling' | 'complicated_subsystem';
  purpose: string;
  responsibilities: string[];
  decisionAuthorities: DecisionAuthority[];
  directMemberCount: number;
  totalMemberCount: number;
  directInternalMemberCount: number;
  directExternalMemberCount: number;
  workspaceLinks: WorkspaceLink[];
  communicationChannels: CommunicationChannel[];
  isManagedService: boolean;
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a temporary or ad-hoc team within the hierarchy. */
export interface TemporaryTeamNode {
  id: string; // format: "temp-team-[slug]"
  name: string;
  type: 'task_force' | 'project_team' | 'program_team' | 'working_group';
  purpose: string;
  responsibilities: string[];
  decisionAuthorities: DecisionAuthority[];
  workspaceLinks: WorkspaceLink[];
  communicationChannels: CommunicationChannel[];
  startDate: string; // ISO 8601 date
  targetEndDate: string | null; // ISO 8601 date; null if open-ended
  directMemberCount: number;
  directInternalMemberCount: number;
  directExternalMemberCount: number;
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a canonical role template shared across teams. Contains no relationship arrays — all relationships expressed as edges. */
export interface MasterRoleNode {
  id: string; // format: "master-role-[slug]"
  name: string;
  purpose: string;
  responsibilities: string[];
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a role instance within a specific team. Contains no relationship arrays — all relationships expressed as edges. */
export interface RoleNode {
  id: string; // format: "role-[master-role-slug]-[team-slug]"
  name: string;
  purpose: string;
  workspaceLinks: WorkspaceLink[];
  communicationChannels: CommunicationChannel[];
  responsibilities: string[];
  decisionAuthorities: DecisionAuthority[];
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a person within the organization. Contains no relationship arrays — all relationships expressed as edges. */
export interface PersonNode {
  id: string; // format: "person-[slug]"
  personType: 'full' | 'stub';
  firstName: string;
  surname: string;
  email: string | null; // null for stub persons
  isExternal: boolean;
  communicationChannels: CommunicationChannel[];
  contributorType: 'manager' | 'individual_contributor' | 'external';
  pronouns: string | null;
  aboutMe: string | null;
  personalityProfiles: PersonalityProfile[];
  socialLinks: SocialLink[];
  avatarUrl: string | null;
  location: GeographicalLocation;
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents an end-to-end value stream crossing team boundaries. Contains no relationship arrays — all relationships expressed as edges. */
export interface ValueStreamNode {
  id: string; // format: "value-stream-[slug]"
  name: string;
  businessOutcomes: string[];
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a governance body. Contains no membership arrays — all membership relationships expressed as edges. */
export interface GovernanceBodyNode {
  id: string; // format: "governance-body-[slug]"
  name: string;
  purpose: string;
  type: 'board' | 'committee' | 'council' | 'forum' | 'review_board';
  cadence: 'weekly' | 'fortnightly' | 'monthly' | 'bimonthly' | 'semiannual' | 'annual' | 'on_demand';
  inputs: string[];
  outputs: string[];
  decisionRecordUrl: string | null;
  workspaceLinks: WorkspaceLink[];
  communicationChannels: CommunicationChannel[];
  decisionAuthorities: DecisionAuthority[];
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a service provided or consumed within the organizational operating model. */
export interface ServiceNode {
  id: string; // format: "service-[slug]"
  name: string;
  description: string;
  type: 'technical' | 'business' | 'customer_facing';
  serviceDomain: string;
  sla: string | null;
  requestUrl: string | null;
  communicationChannels: CommunicationChannel[];
  status: 'pipeline' | 'active' | 'archived';
  createdAt: string;
  createdBy: PersonNode['id'];
  updatedAt: string;
  updatedBy: PersonNode['id'];
  history: HistoryEntry[];
}

/** Represents a cost center used for financial tracking and budget allocation. */
export interface CostCenterNode {
  id: string; // format: "cost-center-[slug]"
  name: string;
  code: string; // ERP/GL code e.g. "CC-4720"
  externalId: string | null; // for SAP/Oracle bi-directional sync
  type: 'operational' | 'shared_service' | 'project' | 'overhead' | 'investment';
  currency: string; // ISO 4217 e.g. "EUR"
  annualBudget: number | null;
  fiscalYearStart: string | null; // ISO 8601 date
  chargebackModel: 'none' | 'showback' | 'chargeback' | 'allocation';
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a measurable key performance indicator used to track outcomes, performance, or health. */
export interface KPINode {
  id: string; // format: "kpi-[slug]"
  name: string;
  abbreviation: string | null; // e.g. "MTTR", "NPS"
  description: string;
  type: 'lagging' | 'leading';
  // lagging = measures outcomes after the fact (revenue, churn, incidents)
  // leading = predicts future performance (pipeline coverage, test coverage)
  category: 'financial' | 'operational' | 'customer' | 'people' | 'quality' | 'risk';
  unit: string; // e.g. "hours", "EUR", "%", "count", "NPS score"
  direction: 'higher_is_better' | 'lower_is_better' | 'target_is_better';
  targetValue: number | null;
  targetOperator: 'gte' | 'lte' | 'eq' | null; // applies when direction is target_is_better
  measurementFrequency: 'daily' | 'weekly' | 'monthly' | 'quarterly' | 'annual';
  dashboardUrl: string | null;
  status: 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents an Objective or Key Result within an OKR cycle. */
export interface OKRNode {
  id: string; // format: "okr-[slug]"
  nodeType: 'objective' | 'key_result';
  // objective  = qualitative goal; has no numeric target itself
  // key_result = quantitative outcome; must reference at least one KPI via MeasuredByKPIEdge
  title: string;
  description: string | null;
  period: string; // e.g. "Q2-2025", "FY2025", "H1-2025"
  startDate: string; // ISO 8601
  endDate: string; // ISO 8601
  targetValue: number | null; // null for objectives; set on key_results
  currentValue: number | null; // null for objectives; set on key_results
  unit: string | null; // mirrors linked KPI unit; null for objectives
  progressStatus: 'not_started' | 'on_track' | 'at_risk' | 'off_track' | 'achieved' | 'missed';
  confidenceScore: number | null; // 0–100; team's subjective confidence rating
  status: 'draft' | 'active' | 'closed';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** Represents a standardised repeatable workflow for delivering services or managing operations. */
export interface ProcessNode {
  id: string; // format: "process-[slug]"
  name: string;
  type: string;
  description: string;
  steps: string[];
  documentationUrl: string | null;
  status: 'pipeline' | 'active' | 'archived';
  createdAt: string;
  updatedAt: string;
  history: HistoryEntry[];
}

/** LegalEntity → LegalEntity: subsidiary relationship. */
export interface LegalEntityParentOfEdge {
  id: string; // format: "edge-[source-slug]-legal_entity_parent_of-[target-slug]"
  relationshipType: 'legal_entity_parent_of';
  source: LegalEntityNode['id'];
  target: LegalEntityNode['id'];
}

/** Person → LegalEntity: person is the legal representative of this legal entity. */
export interface LegalRepresentativeEdge {
  id: string; // format: "edge-[source-slug]-is_legal_representative_of-[target-slug]"
  relationshipType: 'is_legal_representative_of';
  source: PersonNode['id'];
  target: LegalEntityNode['id'];
}

/** Person | GovernanceBody | CostCenter → LegalEntity: node belongs to or is registered under this legal entity. */
export interface BelongsToLegalEntityEdge {
  id: string; // format: "edge-[source-slug]-belongs_to_legal_entity-[target-slug]"
  relationshipType: 'belongs_to_legal_entity';
  source: PersonNode['id'] | GovernanceBodyNode['id'] | CostCenterNode['id'];
  target: LegalEntityNode['id'];
}

/** Team → LegalEntity: team is operated as a managed service by this legal entity. */
export interface ManagedByLegalEntityEdge {
  id: string; // format: "edge-[source-slug]-managed_by_legal_entity-[target-slug]"
  relationshipType: 'managed_by_legal_entity';
  source: TeamNode['id'];
  target: LegalEntityNode['id'];
}

/** Role → TemporaryTeam: role is the sponsor of this temporary team or project. */
export interface SponsorsEdge {
  id: string; // format: "edge-[source-slug]-sponsors-[target-slug]"
  relationshipType: 'sponsors';
  source: RoleNode['id'];
  target: TemporaryTeamNode['id'];
}

/** Team → Team: permanent team hierarchy. */
export interface ParentOfEdge {
  id: string; // format: "edge-[source-slug]-parent_of-[target-slug]"
  relationshipType: 'parent_of';
  source: TeamNode['id'];
  target: TeamNode['id'];
}

/** TemporaryTeam → Team | TemporaryTeam: temporary team parent relationship. Uses distinct literal to preserve discriminated union. */
export interface TempTeamParentOfEdge {
  id: string; // format: "edge-[source-slug]-temp_team_parent_of-[target-slug]"
  relationshipType: 'temp_team_parent_of';
  source: TemporaryTeamNode['id'];
  target: TeamNode['id'] | TemporaryTeamNode['id'];
}

/** Team | TemporaryTeam → Role: role exists within a team. */
export interface HasRoleEdge {
  id: string; // format: "edge-[source-slug]-has_role-[target-slug]"
  relationshipType: 'has_role';
  source: TeamNode['id'] | TemporaryTeamNode['id'];
  target: RoleNode['id'];
}

/** Role → MasterRole: role is an instance of a master role template. */
export interface InstanceOfEdge {
  id: string; // format: "edge-[source-slug]-instance_of-[target-slug]"
  relationshipType: 'instance_of';
  source: RoleNode['id'];
  target: MasterRoleNode['id'];
}

/** Person → Person: person directly supervises another person. */
export interface SupervisesEdge {
  id: string; // format: "edge-[source-slug]-supervises-[target-slug]"
  relationshipType: 'supervises';
  source: PersonNode['id']; // supervisor
  target: PersonNode['id']; // report
}

/** Person → Role: person currently fills this role. */
export interface FillsRoleEdge {
  id: string; // format: "edge-[source-slug]-fills_role-[target-slug]"
  relationshipType: 'fills_role';
  source: PersonNode['id'];
  target: RoleNode['id'];
}

/** Person → Team | TemporaryTeam: person is the lead of this team. */
export interface LeadOfEdge {
  id: string; // format: "edge-[source-slug]-lead_of-[target-slug]"
  relationshipType: 'lead_of';
  source: PersonNode['id'];
  target: TeamNode['id'] | TemporaryTeamNode['id'];
}

/** Person → Team | TemporaryTeam: person is a member of this team. */
export interface MemberOfEdge {
  id: string; // format: "edge-[source-slug]-member_of-[target-slug]"
  relationshipType: 'member_of';
  source: PersonNode['id'];
  target: TeamNode['id'] | TemporaryTeamNode['id'];
}

/** Team | TemporaryTeam | Role | Service | Process → ValueStream: contributes to this value stream. */
export interface ContributesToEdge {
  id: string; // format: "edge-[source-slug]-contributes_to-[target-slug]"
  relationshipType: 'contributes_to';
  source: TeamNode['id'] | TemporaryTeamNode['id'] | RoleNode['id'] | ServiceNode['id'] | ProcessNode['id'];
  target: ValueStreamNode['id'];
}

/** Role → Service: role is the owner of this service. */
export interface OwnsServiceEdge {
  id: string; // format: "edge-[source-slug]-owns_service-[target-slug]"
  relationshipType: 'owns_service';
  source: RoleNode['id'];
  target: ServiceNode['id'];
}

/** Team | TemporaryTeam → Service: team is the provider of this service. */
export interface ProvidesServiceEdge {
  id: string; // format: "edge-[source-slug]-provides_service-[target-slug]"
  relationshipType: 'provides_service';
  source: TeamNode['id'] | TemporaryTeamNode['id'];
  target: ServiceNode['id'];
}

/** Team | TemporaryTeam → Service: team consumes this service. */
export interface ConsumesServiceEdge {
  id: string; // format: "edge-[source-slug]-consumes_service-[target-slug]"
  relationshipType: 'consumes_service';
  source: TeamNode['id'] | TemporaryTeamNode['id'];
  target: ServiceNode['id'];
}

/** LegalEntity → Service: legal entity is the provider of this service. */
export interface LegalEntityProvidesServiceEdge {
  id: string; // format: "edge-[source-slug]-legal_entity_provides_service-[target-slug]"
  relationshipType: 'legal_entity_provides_service';
  source: LegalEntityNode['id'];
  target: ServiceNode['id'];
}

/** Service → Service: this service depends on another service. */
export interface ServiceDependsOnEdge {
  id: string; // format: "edge-[source-slug]-depends_on-[target-slug]"
  relationshipType: 'depends_on';
  source: ServiceNode['id'];
  target: ServiceNode['id'];
}

/** Person | Role → GovernanceBody: membership in a governance body with a specific role type. */
export interface GovernanceMemberEdge {
  id: string; // format: "edge-[source-slug]-governance_member-[target-slug]"
  relationshipType: 'governance_member';
  memberType: 'chair' | 'facilitator' | 'standing' | 'advisory';
  source: PersonNode['id'] | RoleNode['id'];
  target: GovernanceBodyNode['id'];
}

/** GovernanceBody → Team | TemporaryTeam: governs this team. */
export interface GovernsEdge {
  id: string; // format: "edge-[source-slug]-governs-[target-slug]"
  relationshipType: 'governs';
  source: GovernanceBodyNode['id'];
  target: TeamNode['id'] | TemporaryTeamNode['id'];
}

/** GovernanceBody → Service: governs this service. */
export interface GovernsServiceEdge {
  id: string; // format: "edge-[source-slug]-governs_service-[target-slug]"
  relationshipType: 'governs_service';
  source: GovernanceBodyNode['id'];
  target: ServiceNode['id'];
}

/** GovernanceBody → GovernanceBody: escalates unresolved decisions. */
export interface EscalatesToEdge {
  id: string; // format: "edge-[source-slug]-escalates_to-[target-slug]"
  relationshipType: 'escalates_to';
  source: GovernanceBodyNode['id'];
  target: GovernanceBodyNode['id'];
}

/** GovernanceBody → GovernanceBody | Role: delegates authority. */
export interface DelegatesToEdge {
  id: string; // format: "edge-[source-slug]-delegates_to-[target-slug]"
  relationshipType: 'delegates_to';
  source: GovernanceBodyNode['id'];
  target: GovernanceBodyNode['id'] | RoleNode['id'];
}

/** Team | TemporaryTeam → TemporaryTeam: team contributes to or is linked to a temporary team or project. */
export interface ContributingTeamEdge {
  id: string; // format: "edge-[source-slug]-contributing_team-[target-slug]"
  relationshipType: 'contributing_team';
  source: TeamNode['id'] | TemporaryTeamNode['id'];
  target: TemporaryTeamNode['id'];
}

/** ValueStream → Process: value stream sits at the top of the process hierarchy; a process belongs under a value stream. A process cannot be the parent of a value stream. */
export interface ValueStreamParentOfProcessEdge {
  id: string; // format: "edge-[source-slug]-value_stream_parent_of_process-[target-slug]"
  relationshipType: 'value_stream_parent_of_process';
  source: ValueStreamNode['id'];
  target: ProcessNode['id'];
}

/** Process → Process: parent process decomposes into this child process. */
export interface ProcessParentOfEdge {
  id: string; // format: "edge-[source-slug]-process_parent_of-[target-slug]"
  relationshipType: 'process_parent_of';
  source: ProcessNode['id'];
  target: ProcessNode['id'];
}

/** Role | Team | GovernanceBody → Process: node is integral to executing this process. */
export interface ExecutesProcessEdge {
  id: string; // format: "edge-[source-slug]-executes_process-[target-slug]"
  relationshipType: 'executes_process';
  source: RoleNode['id'] | TeamNode['id'] | GovernanceBodyNode['id'];
  target: ProcessNode['id'];
}

/** Role → Process: role is the owner of this process. */
export interface OwnsProcessEdge {
  id: string; // format: "edge-[source-slug]-owns_process-[target-slug]"
  relationshipType: 'owns_process';
  source: RoleNode['id'];
  target: ProcessNode['id'];
}

/** Role → CostCenter: role is the owner of this cost center. */
export interface OwnsCostCenterEdge {
  id: string; // format: "edge-[source-slug]-owns_cost_center-[target-slug]"
  relationshipType: 'owns_cost_center';
  source: RoleNode['id'];
  target: CostCenterNode['id'];
}

/** Team | TemporaryTeam | Role | Service → CostCenter: expenditure is booked to this cost center. */
export interface BookedToCostCenterEdge {
  id: string; // format: "edge-[source-slug]-booked_to_cost_center-[target-slug]"
  relationshipType: 'booked_to_cost_center';
  allocationPercent: number | null; // 0–100; null implies 100%
  source: TeamNode['id'] | TemporaryTeamNode['id'] | RoleNode['id'] | ServiceNode['id'];
  target: CostCenterNode['id'];
}

/** CostCenter → CostCenter: hierarchical parent relationship between cost centers. */
export interface CostCenterParentOfEdge {
  id: string; // format: "edge-[source-slug]-cost_center_parent_of-[target-slug]"
  relationshipType: 'cost_center_parent_of';
  source: CostCenterNode['id'];
  target: CostCenterNode['id'];
}

/** CostCenter → CostCenter: cost center charges back costs to another cost center. */
export interface ChargesBackToEdge {
  id: string; // format: "edge-[source-slug]-charges_back_to-[target-slug]"
  relationshipType: 'charges_back_to';
  allocationKey: 'headcount' | 'usage' | 'fixed' | 'custom' | null;
  source: CostCenterNode['id'];
  target: CostCenterNode['id'];
}

/** KPI → KPI: parent KPI decomposes into this child KPI. Allows strategic KPIs to cascade into operational metrics. */
export interface KPIParentOfEdge {
  id: string; // format: "edge-[source-slug]-kpi_parent_of-[target-slug]"
  relationshipType: 'kpi_parent_of';
  source: KPINode['id'];
  target: KPINode['id'];
}

/** Team | TemporaryTeam | Role | Service | ValueStream | Person | Process → KPI: node contributes to this KPI with a typed contribution role. */
export interface ContributesToKPIEdge {
  id: string; // format: "edge-[source-slug]-contributes_to_kpi-[target-slug]"
  relationshipType: 'contributes_to_kpi';
  contributionType: 'owns' | 'contributes' | 'influences';
  // owns      = single accountable node; should be exactly one per KPI
  // contributes = directly moves the metric
  // influences = upstream dependency, not directly measured
  source: TeamNode['id'] | TemporaryTeamNode['id'] | RoleNode['id'] | ServiceNode['id'] | ValueStreamNode['id'] | PersonNode['id'] | ProcessNode['id'];
  target: KPINode['id'];
}

/** KPI → KPI: this KPI is causally dependent on another KPI moving first. Distinct from parent-child hierarchy — models leading indicator relationships. */
export interface KPIDependsOnEdge {
  id: string; // format: "edge-[source-slug]-kpi_depends_on-[target-slug]"
  relationshipType: 'kpi_depends_on';
  source: KPINode['id'];
  target: KPINode['id'];
}

/** Objective → KeyResult: objective decomposes into this key result. Source must be nodeType 'objective'; target must be nodeType 'key_result'. Always exactly two levels deep — do not chain objectives to objectives. */
export interface OKRParentOfEdge {
  id: string; // format: "edge-[source-slug]-okr_parent_of-[target-slug]"
  relationshipType: 'okr_parent_of';
  source: OKRNode['id'];
  target: OKRNode['id'];
}

/** OKR → OKR: this OKR supports a higher-level OKR. Models the company → team → individual cascade. Distinct from okr_parent_of — this crosses ownership boundaries between different teams or levels. */
export interface AlignedToOKREdge {
  id: string; // format: "edge-[source-slug]-aligned_to_okr-[target-slug]"
  relationshipType: 'aligned_to_okr';
  source: OKRNode['id']; // child (team or individual level)
  target: OKRNode['id']; // parent (company or leadership level)
}

/** Team | TemporaryTeam | Role | Person → OKR: node is the accountable owner of this OKR for its stated period. */
export interface OwnsOKREdge {
  id: string; // format: "edge-[source-slug]-owns_okr-[target-slug]"
  relationshipType: 'owns_okr';
  source: TeamNode['id'] | TemporaryTeamNode['id'] | RoleNode['id'] | PersonNode['id'];
  target: OKRNode['id'];
}

/** KeyResult → KPI: this key result is measured by this KPI. The KPI is the persistent measurement instrument; the key result is the time-boxed commitment to move it. Source must be nodeType 'key_result'. A key result may reference more than one KPI. */
export interface MeasuredByKPIEdge {
  id: string; // format: "edge-[source-slug]-measured_by_kpi-[target-slug]"
  relationshipType: 'measured_by_kpi';
  source: OKRNode['id']; // must be nodeType 'key_result'
  target: KPINode['id'];
}

/** Node → Domain: this node belongs to a domain. */
export interface BelongsToDomainEdge {
  id: string; // format: "edge-[source-slug]-belongs_to_domain-[target-slug]"
  relationshipType: 'belongs_to_domain';
  source: RoleNode['id'] | MasterRoleNode['id'] | TeamNode['id'] | TemporaryTeamNode['id'] | GovernanceBodyNode['id'] | ServiceNode['id'] | CostCenterNode['id'] | ProcessNode['id'] | KPINode['id'] | OKRNode['id'];
  target: DomainNode['id'];
}

/** Union of all valid directed edge types in the graph. */
export type GraphEdge =
  | LegalEntityParentOfEdge
  | LegalRepresentativeEdge
  | BelongsToLegalEntityEdge
  | ManagedByLegalEntityEdge
  | SponsorsEdge
  | ParentOfEdge
  | TempTeamParentOfEdge
  | HasRoleEdge
  | InstanceOfEdge
  | SupervisesEdge
  | FillsRoleEdge
  | LeadOfEdge
  | MemberOfEdge
  | ContributesToEdge
  | ValueStreamParentOfProcessEdge
  | ProcessParentOfEdge
  | ExecutesProcessEdge
  | OwnsProcessEdge
  | OwnsServiceEdge
  | ProvidesServiceEdge
  | ConsumesServiceEdge
  | LegalEntityProvidesServiceEdge
  | ServiceDependsOnEdge
  | GovernanceMemberEdge
  | GovernsEdge
  | GovernsServiceEdge
  | EscalatesToEdge
  | DelegatesToEdge
  | ContributingTeamEdge
  | OwnsCostCenterEdge
  | BookedToCostCenterEdge
  | CostCenterParentOfEdge
  | ChargesBackToEdge
  | KPIParentOfEdge
  | ContributesToKPIEdge
  | KPIDependsOnEdge
  | OKRParentOfEdge
  | AlignedToOKREdge
  | OwnsOKREdge
  | MeasuredByKPIEdge
  | BelongsToDomainEdge;

/** The complete graph data model for the Plait Org Canvas. */
export interface GraphData {
  teams: TeamNode[];
  temporaryTeams: TemporaryTeamNode[];
  legalEntities: LegalEntityNode[];
  persons: PersonNode[];
  masterRoles: MasterRoleNode[];
  roles: RoleNode[];
  valueStreams: ValueStreamNode[];
  domains: DomainNode[];
  governanceBodies: GovernanceBodyNode[];
  services: ServiceNode[];
  costCenters: CostCenterNode[];
  kpis: KPINode[];
  okrs: OKRNode[];
  processes: ProcessNode[];
  edges: GraphEdge[];
}
