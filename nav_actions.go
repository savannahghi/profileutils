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
	Title:          RequestsNavActionTitle,
	OnTapRoute:     KYCRequestsRoute,
	Icon:           RequestNavActionIcon,
	IsParent:       true,
	IsHighPriority: true,
}

// PartnerNavAction ...
var PartnerNavAction NavigationAction = NavigationAction{
	Title:          PartnerNavActionTitle,
	Icon:           PartnerNavActionIcon,
	IsParent:       true,
	IsHighPriority: true,
}

// ConsumerNavAction ...
var ConsumerNavAction NavigationAction = NavigationAction{
	Title:          PartnerNavActionTitle,
	Icon:           PartnerNavActionIcon,
	IsParent:       true,
	IsHighPriority: true,
}

// AgentNavAction ...
var AgentNavAction NavigationAction = NavigationAction{
	Group:    ActionGoupAgent,
	Title:    AgentNavActionTitle,
	Icon:     AgentNavActionIcon,
	IsParent: true,
}

// AgentRegistrationAction ...
var AgentRegistrationAction NavigationAction = NavigationAction{
	Group:      ActionGoupAgent,
	Title:      AgentRegistrationActionTitle,
	OnTapRoute: AgentRegistrationRoute,
}

// AgentIdentificationAction ...
var AgentIdentificationAction NavigationAction = NavigationAction{
	Group:      ActionGoupAgent,
	Title:      AgentIdentificationActionTitle,
	OnTapRoute: AgentIdentificationRoute,
}

// PatientNavAction ...
var PatientNavAction NavigationAction = NavigationAction{
	Group:    ActionGoupPatient,
	Title:    PatientNavActionTitle,
	Icon:     PatientNavActionIcon,
	IsParent: true,
}

// PatientRegistrationAction ...
var PatientRegistrationAction NavigationAction = NavigationAction{
	Group:      ActionGoupPatient,
	Title:      PatientRegistrationActionTitle,
	OnTapRoute: PatientRegistrationRoute,
}

// PatientIdentificationAction ...
var PatientIdentificationAction NavigationAction = NavigationAction{
	Group:      ActionGoupPatient,
	Title:      PatientIdentificationActionTitle,
	OnTapRoute: PatientIdentificationRoute,
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
