package profileutils

// HomeNavAction ...
var HomeNavAction NavigationAction = NavigationAction{
	Title:          HomeNavActionTitle,
	OnTapRoute:     HomeRoute,
	Icon:           HomeNavActionIcon,
	IsHighPriority: true,
}

// HelpNavAction ...
var HelpNavAction NavigationAction = NavigationAction{
	Title:          HelpNavActionTitle,
	OnTapRoute:     HelpNavActionActionRoute,
	Icon:           HelpNavActionIcon,
	IsHighPriority: true,
}

// RequestsNavAction ...
var RequestsNavAction NavigationAction = NavigationAction{
	Title:              RequestsNavActionTitle,
	OnTapRoute:         KYCRequestsRoute,
	Icon:               RequestNavActionIcon,
	IsParent:           true,
	IsHighPriority:     true,
	RequiredPermission: &CanProcessKYC,
}

// PartnerNavAction ...
var PartnerNavAction NavigationAction = NavigationAction{
	Title:              PartnerNavActionTitle,
	Icon:               PartnerNavActionIcon,
	IsParent:           true,
	IsHighPriority:     true,
	RequiredPermission: &CanViewPartner,
}

// ConsumerNavAction ...
var ConsumerNavAction NavigationAction = NavigationAction{
	Title:              PartnerNavActionTitle,
	Icon:               PartnerNavActionIcon,
	IsParent:           true,
	IsHighPriority:     true,
	RequiredPermission: &CanViewConsumers,
}

// AgentNavAction ...
var AgentNavAction NavigationAction = NavigationAction{
	Group:              ActionGoupAgent,
	Title:              AgentNavActionTitle,
	Icon:               AgentNavActionIcon,
	IsParent:           true,
	RequiredPermission: &CanViewAgent,
}

// AgentRegistrationAction ...
var AgentRegistrationAction NavigationAction = NavigationAction{
	Group:              ActionGoupAgent,
	Title:              AgentRegistrationActionTitle,
	OnTapRoute:         AgentRegistrationRoute,
	RequiredPermission: &CanRegisterAgent,
}

// AgentIdentificationAction ...
var AgentIdentificationAction NavigationAction = NavigationAction{
	Group:              ActionGoupAgent,
	Title:              AgentIdentificationActionTitle,
	OnTapRoute:         AgentIdentificationRoute,
	RequiredPermission: &CanIdentifyAgent,
}

// PatientNavAction ...
var PatientNavAction NavigationAction = NavigationAction{
	Group:              ActionGoupPatient,
	Title:              PatientNavActionTitle,
	Icon:               PatientNavActionIcon,
	IsParent:           true,
	RequiredPermission: &CanViewPatient,
}

// PatientRegistrationAction ...
var PatientRegistrationAction NavigationAction = NavigationAction{
	Group:              ActionGoupPatient,
	Title:              PatientRegistrationActionTitle,
	OnTapRoute:         PatientRegistrationRoute,
	RequiredPermission: &CanCreatePatient,
}

// PatientIdentificationAction ...
var PatientIdentificationAction NavigationAction = NavigationAction{
	Group:              ActionGoupPatient,
	Title:              PatientIdentificationActionTitle,
	OnTapRoute:         PatientIdentificationRoute,
	RequiredPermission: &CanIdentifyPatient,
}

// AllNavActions is the list of navgigation actions combined
// MAKESURE you add all the Navigation actions ABOVE to this list
var AllNavActions []NavigationAction = []NavigationAction{
	HomeNavAction,
	HelpNavAction,
	RequestsNavAction,
	PartnerNavAction,
	ConsumerNavAction,
	AgentNavAction,
	AgentRegistrationAction,
	AgentIdentificationAction,
	PatientNavAction,
	PatientRegistrationAction,
	PatientIdentificationAction,
}
