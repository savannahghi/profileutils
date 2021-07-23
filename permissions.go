package profileutils

var AllPermisions []Permission = []Permission{
	{Group: PermissionGroupAdministration, Operation: "*.*", Description: "Super administrator"},
	{Group: PermissionGroupAdministration, Operation: "role.create", Description: "Can create role"},
	{Group: PermissionGroupAdministration, Operation: "role.edit", Description: "Can edit role"},

	//employee management
	{Group: PermissionGroupEmployee, Operation: "employee.create", Description: "Can create employee"},
	{Group: PermissionGroupEmployee, Operation: "employee.remove", Description: "Can remove employee"},

	// agent management
	{Group: PermissionGroupAgent, Operation: "agent.view", Description: "Can view agents"},
	{Group: PermissionGroupAgent, Operation: "agent.register", Description: "Can register agent"},
	{Group: PermissionGroupAgent, Operation: "agent.suspend", Description: "Can suspend agent"},
	{Group: PermissionGroupAgent, Operation: "agent.unsuspend", Description: "Can unsuspebd agent"},

	// partner management
	{Group: PermissionGroupPartner, Operation: "partner.view", Description: "Can view partners"},
	{Group: PermissionGroupPartner, Operation: "partner.create", Description: "Can create partner"},
	{Group: PermissionGroupPartner, Operation: "kyc.process", Description: "Can process KYC"},

	// consumer management
	{Group: PermissionGroupConsumer, Operation: "consumer.view", Description: "Can view consumers"},
	{Group: PermissionGroupConsumer, Operation: "consumer.create", Description: "Can add consumer"},

	// patient management
	{Group: PermissionGroupPatient, Operation: "patient.view", Description: "Can view patient"},
	{Group: PermissionGroupPatient, Operation: "patient.create", Description: "Can add patient"},
}
