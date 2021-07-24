package profileutils

var CanViewRole = Permission{Group: PermissionGroupAdministration, Operation: "role.view", Description: "Can view role"}
var CanCreateRole = Permission{Group: PermissionGroupAdministration, Operation: "role.create", Description: "Can create role"}
var CanEditRole = Permission{Group: PermissionGroupAdministration, Operation: "role.create", Description: "Can create role"}

var CanViewEmployee = Permission{Group: PermissionGroupEmployee, Operation: "employee.view", Description: "Can create employee"}
var CanCreateEmployee = Permission{Group: PermissionGroupEmployee, Operation: "employee.create", Description: "Can create employee"}
var CanRemoveEmploye = Permission{Group: PermissionGroupEmployee, Operation: "employee.remove", Description: "Can remove employee"}

var CanViewAgent = Permission{Group: PermissionGroupAgent, Operation: "agent.view", Description: "Can view agents"}
var CanRegisterAgent = Permission{Group: PermissionGroupAgent, Operation: "agent.register", Description: "Can register agent"}
var CanIdentifyAgent = Permission{Group: PermissionGroupAgent, Operation: "agent.register", Description: "Can identify agent"}
var CanSuspendAgent = Permission{Group: PermissionGroupAgent, Operation: "agent.suspend", Description: "Can suspend agent"}
var CanUnsuspendAgent = Permission{Group: PermissionGroupAgent, Operation: "agent.unsuspend", Description: "Can unsuspebd agent"}

var CanViewPartner = Permission{Group: PermissionGroupPartner, Operation: "partner.view", Description: "Can view partners"}
var CanCreatePartner = Permission{Group: PermissionGroupPartner, Operation: "partner.create", Description: "Can create partner"}
var CanProcessKYC = Permission{Group: PermissionGroupPartner, Operation: "kyc.process", Description: "Can process KYC"}

var CanViewConsumers = Permission{Group: PermissionGroupConsumer, Operation: "consumer.view", Description: "Can view consumers"}
var CanCreateConsumer = Permission{Group: PermissionGroupConsumer, Operation: "consumer.create", Description: "Can add consumer"}

var CanViewPatient = Permission{Group: PermissionGroupPatient, Operation: "patient.view", Description: "Can view patient"}
var CanCreatePatient = Permission{Group: PermissionGroupPatient, Operation: "patient.create", Description: "Can add patient"}
var CanIdentifyPatient = Permission{Group: PermissionGroupPatient, Operation: "patient.identify", Description: "Can identify patient"}

var AllPermisions []Permission = []Permission{
	//role management
	CanViewRole,
	CanCreateRole,
	CanEditRole,

	//employee management
	CanViewEmployee,
	CanCreateEmployee,
	CanRemoveEmploye,

	// agent management
	CanViewAgent,
	CanRegisterAgent,
	CanSuspendAgent,
	CanUnsuspendAgent,

	// partner management
	CanViewPartner,
	CanCreatePartner,
	CanProcessKYC,

	// consumer management
	CanViewConsumers,
	CanCreateConsumer,

	// patient management
	CanViewPatient,
	CanCreatePatient,
	CanIdentifyPatient,
}

func GetPermissionByOperation(operation string) *Permission {
	for _, permission := range AllPermisions {
		if permission.Operation == operation {
			return &permission
		}
	}
	return nil
}
