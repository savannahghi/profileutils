package profileutils

const (
	// the available groups for nested actions
	ActionGroupAgent   ActionGroup = "AGENT"
	ActionGroupPatient ActionGroup = "PATIENT"

	// the available groups for nested actions
	PermissionGroupAdministration PermissionGroup = "Administration"
	PermissionGroupEmployee       PermissionGroup = "Employee"
	PermissionGroupAgent          PermissionGroup = "Agent"
	PermissionGroupPartner        PermissionGroup = "Partner"
	PermissionGroupConsumer       PermissionGroup = "Consumer"
	PermissionGroupPatient        PermissionGroup = "Patient"

	// Icon links for navactions
	// StaticBase is the default path at which static assets are hosted
	StaticBase = "https://assets.healthcloud.co.ke"

	AgentNavActionIcon    = StaticBase + "/actions/agent_navaction.png"
	ConsumerNavActionIcon = StaticBase + "/actions/consumer_navaction.png"
	HelpNavActionIcon     = StaticBase + "/actions/help_navaction.png"
	HomeNavActionIcon     = StaticBase + "/actions/home_navaction.png"
	KYCNavActionIcon      = StaticBase + "/actions/kyc_navaction.png"
	PartnerNavActionIcon  = StaticBase + "/actions/partner_navaction.png"
	PatientNavActionIcon  = StaticBase + "/actions/patient_navaction.png"
	RequestNavActionIcon  = StaticBase + "/actions/request_navaction.png"

	// On Tap Routes
	HomeRoute                  = "/home"
	PatientRegistrationRoute   = "/addPatient"
	PatientIdentificationRoute = "/patients"
	HelpNavActionActionRoute   = "/helpCenter"
	KYCRequestsRoute           = "/admin"
	AgentRegistrationRoute     = "/agentRegistration"
	AgentIdentificationRoute   = "/agentIdentification"

	// Navigation actions
	HomeNavActionTitle               = "Home"
	HelpNavActionTitle               = "Help"
	PatientNavActionTitle            = "Patient"
	PatientRegistrationActionTitle   = "Patient Registration"
	PatientIdentificationActionTitle = "Patient Identification"
	RequestsNavActionTitle           = "Requests"
	AgentNavActionTitle              = "Agent"
	AgentRegistrationActionTitle     = "Agent Registration"
	AgentIdentificationActionTitle   = "Agent Identification"
	AgentRegistrationTitle           = "Agent Registration"
	ConsumerNavActionTitle           = "Consumer"
	PartnerNavActionTitle            = "Partner"
)
