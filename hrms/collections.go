package hrms

import (
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

const ORG_COLLECTION_ORGANIZATION = "organization"
const ORG_COLLECTION_DEPARTMENTS = "org-departments"
const ORG_COLLECTION_DOCUMENTS = "org-documents"
const ORG_COLLECTION_EMPLOYEE_TYPE = "org-employee-types"
const ORG_COLLECTION_PERMISSIONS = "org-permissions"
const ORG_COLLECTION_AUTH = "auth"
const ORG_COLLECTION_ATTENDANCE_CONF = "org-attendance-conf"
const ORG_COLLECTION_HOLIDAYS = "org-holidays"
const ORG_COLLECTION_LEAVES_CONF = "org-leaves-conf"
const ORG_COLLECTION_MARQUE = "org-marque"
const ORG_COLLECTION_POST = "org-post"
const ORG_COLLECTION_POLL = "org-poll"
const ORG_COLLECTION_ROLES = "org-roles"
const ORG_COLLECTION_INVENTORY = "org-inventory"
const ORG_COLLECTION_ASSIGN_ASSET = "org-assign-asset"
const ORG_COLLECTION_INVENTORY_DEPARTMENT = "org-inventory-department"
const ORG_COLLECTION_INVENTORY_STORE = "org-inventory-store"
const ORG_COLLECTION_INVENTORY_BUCKET = "org-inventory-bucket"
const ORG_COLLECTION_CLIENT_ID_COUNTER = "org-client-id-counter"
const ORG_COLLECTION_CLIENT_INFO = "org-client-info"
const ORG_COLLECTION_CLIENT_CONTACT_INFO = "org-client-contact-info"
const ORG_COLLECTION_CLIENT_MEETING_INFO = "org-client-meeting-info"
const ORG_COLLECTION_CLIENT_HOLIDAYS = "org-client_holidays"
const ORG_COLLECTION_POLICY = "org-policy"
const ORG_COLLECTION_POLICY_STATS = "org-policy-stats"
const ORG_COLLECTION_SALARY_COMPONENTS = "org-salary_components"
const ORG_COLLECTION_SALARY_GROUPS = "org-salary_groups"
const ORG_COLLECTION_SALARY_STRUCTURES = "org-salary_structures"
const ORG_COLLECTION_OFFER_LETTERS = "org-offer-letters"
const ORG_COLLECTION_EMPLOYEE_SHIFTS = "org-employee-shifts"

const DEPARTMENT_COLLECTION_DESIGNATIONS = "department-designations"
const DEPARTMENT_COLLECTION_ATTENDANCE_CONF = "department-attendance-conf"

const EMPLOYEE_COLLECTION_PERSONAL = "employee-personal"
const EMPLOYEE_COLLECTION_PROFESSIONAL = "employee-professional"
const EMPLOYEE_COLLECTION_DOCUMENTS = "employee-documents"
const EMPLOYEE_COLLECTION_SUMMARY = "employee-summary"
const EMPLOYEE_COLLECTION_ATTENDANCE_CONF = "employee-attendance-conf"
const EMPLOYEE_COLLECTION_ATTENDANCE = "employee-attendance"
const EMPLOYEE_COLLECTION_DAILY_ATN_OBJECT = "employee-daily-atn-object"
const EMPLOYEE_COLLECTION_LEAVES = "employee-leaves"
const EMPLOYEE_COLLECTION_LEAVE_STATS = "employee-leave-stats"
const EMPLOYEE_COLLECTION_QUALIFICATION = "employee-qualification"
const EMPLOYEE_COLLECTION_CERTIFICATE_AND_TRAINING = "employee-cert-and-training"
const EMPLOYEE_COLLECTION_ATTENDANCE_REGULARIZATION = "employee-attendance-regularization"
const EMPLOYEE_COLLECTION_REWARD = "employee-rewards"
const EMPLOYEE_COLLECTION_TECH_INFO = "employee-tech-info"
const EMPLOYEE_COLLECTION_ASSET = "employee-assets"
const EMPLOYEE_COLLECTION_BANK = "employee-banks"
const EMPLOYEE_COLLECTION_ASSET_COUNTER = "employee-asset-counter"
const EMPLOYEE_COLLECTION_LEAVES_CARRY_FORWARD = "employee-leaves-carry-forward"

const REIMBURSEMENT_COLLECTION_EXPENSE_COUNTER = "reimbursement-expense-counter"
const REIMBURSEMENT_COLLECTION_EXPENSE = "reimbursement-expenses"

const JOBMGR_COLLECTION_JOBS = "jobmgr-jobs"
const JOBMGR_COLLECTION_APPLICATIONS = "jobmgr-applications"
const JOBMGR_COLLECTION_JOB_COUNTER = "jobmgr-job-counter"
const JOBMGR_COLLECTION_CANDIDATES = "jobmgr-candidates"
const JOBMGR_COLLECTION_APPLICATION_COUNTER = "jobmgr-application-counter"
const JOBMGR_COLLECTION_CANDIDATE_COUNTER = "jobmgr-candidate-counter"
const JOBMGR_COLLECTION_ALLOCATION_TABLE = "jobmgr-allocation-table"
const JOBMGR_COLLECTION_DAILY_STATS = "jobmgr-daily-stats"
const JOBMGR_COLLECTION_CALCULATED_STATS = "jobmgr-calculated-stats"
const JOBMGR_COLLECTION_ERROR_RECORDS = "jobmgr-error-records"

const WORKER_COLLECTION_DOMAINS = "worker-domains"

const TICKET_COLLECTION_TICKET_CONFIGS = "ticket-ticket-configs"
const TICKET_COLLECTION_TICKETS = "ticket-tickets"
const TICKET_COLLECTION_TICKET_COUNTER = "ticket-ticket-counter"
const TICKET_COLLECTION_TICKET_LEADS = "ticket-ticket-leads"

const VOLUNTEER_COLLECTION_VOLUNTEER_COUNTER = "volunteer-volunteer-counter"
const VOLUNTEER_COLLECTION_REQUESTS = "volunteer-requests"
const VOLUNTEER_COLLECTION_VOLUNTEERS = "volunteer-volunteers"

const GSORG_COLLECTION_SCHOLARSHIP_REQUESTS = "gsorg-scholarship-requests"
const GSORG_COLLECTION_STUDENT_COUNTER = "gsorg-student-counter"
const GSORG_COLLECTION_DONATION_COUNTER = "gsorg-donation-counter"
const GSORG_COLLECTION_DONOR_COUNTER = "gsorg-donor-counter"
const GSORG_COLLECTION_SUBSCRIPTION_COUNTER = "gsorg-subscription-counter"
const GSORG_COLLECTION_DONATIONS = "gsorg-donations"
const GSORG_COLLECTION_DONATION_PAYMENTS = "gsorg-donation-payments"
const GSORG_COLLECTION_DONORS = "gsorg-donors"
const GSORG_COLLECTION_SUBSCRIPTIONS = "gsorg-subscriptions"
const GSORG_COLLECTION_PAYMENT_ORDERS = "gsorg-payment-orders"
const GSORG_COLLECTION_PAYMENT_PAYMENTS = "gsorg-payment-payments"
const GSORG_COLLECTION_PROCESSED_EVENTS = "gsorg-processed-events"
const NOTIFICATION_COLLECTION_DEVICES = "notification-devices"

const PAYMENT_COLLECTION_ORDERS = "payment-orders"
const PAYMENT_COLLECTION_PAYMENTS = "payment-payments"
const PAYMENT_COLLECTION_PROCESSED_EVENTS = "payment-processed-events"

const ANALYTIC_COLLECTION_ANALYTICS = "analytic-analytics"

func GetPaymentOrdersCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(PAYMENT_COLLECTION_ORDERS)
}

func GetPaymentPaymentsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(PAYMENT_COLLECTION_PAYMENTS)
}

func GetPaymentProcessedEventsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(PAYMENT_COLLECTION_PROCESSED_EVENTS)
}

func GetAnalyticsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ANALYTIC_COLLECTION_ANALYTICS)
}

func GetGSOrgProcessedEventsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_PROCESSED_EVENTS)
}

func GetNotificationDevicesCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(NOTIFICATION_COLLECTION_DEVICES)
}

func GetOrgCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ORGANIZATION)
}

func GetDailyAttendanceObjectCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_DAILY_ATN_OBJECT)
}

func GetOrgAuthCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_AUTH)
}

func GetOrgLeaveConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_LEAVES_CONF)
}

func GetOrgDepartmentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_DEPARTMENTS)
}

func GetOrgDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_DOCUMENTS)
}

func GetOrgHolidayCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_HOLIDAYS)
}

func GetOrgRolesCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ROLES)
}

func GetOrgPermissionsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_PERMISSIONS)
}

func GetOrgEmployeeTypeCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_EMPLOYEE_TYPE)
}

func GetOrgAttendanceConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ATTENDANCE_CONF)
}

func GetEmpPersonalCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_PERSONAL)
}

func GetEmpProfessionalCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_PROFESSIONAL)
}

func GetEmpDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_DOCUMENTS)
}

func GetEmpSummaryCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_SUMMARY)
}

func GetEmpAttendanceConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_ATTENDANCE_CONF)
}

func GetEmpAttendanceCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_ATTENDANCE)
}

func GetEmpLeavesCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_LEAVES)
}

func GetDocumentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_DOCUMENTS)
}

func GetDepDesignationCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(DEPARTMENT_COLLECTION_DESIGNATIONS)
}

func GetDepAttendanceConfCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(DEPARTMENT_COLLECTION_ATTENDANCE_CONF)
}

func GetOrgHolidaysCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_HOLIDAYS)
}

func GetEmpLeaveStatsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_LEAVE_STATS)
}

func GetOrgMarqueCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_MARQUE)
}

func GetOrgPostCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_POST)
}

func GetOrgPollCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_POLL)
}

func GetEmpQualificationCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_QUALIFICATION)
}

func GetOrgInventoryDepartmentCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_INVENTORY_DEPARTMENT)
}

func GetOrgInventoryStoreCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_INVENTORY_STORE)
}

func GetOrgInventoryBucketCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_INVENTORY_BUCKET)
}

func GetOrgInventoryCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_INVENTORY)
}

func GetOrgAssignAssetCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_ASSIGN_ASSET)
}

func GetEmpCertAndTrainingCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_CERTIFICATE_AND_TRAINING)
}

func GetOrgClientInfoCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_CLIENT_INFO)
}

func GetOrgClientContactInfoCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_CLIENT_CONTACT_INFO)
}

func GetOrgClientIdCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_CLIENT_ID_COUNTER)
}

func GetOrgClientMeetingInfoCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_CLIENT_MEETING_INFO)
}

func GetOrgClientHolidaysCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_CLIENT_HOLIDAYS)
}

func GetOrgPolicyCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_POLICY)
}

func GetOrgSalaryComponentsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_SALARY_COMPONENTS)
}

func GetOrgSalaryStructuresCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_SALARY_STRUCTURES)
}

func GetOrgOfferLettersCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_OFFER_LETTERS)
}

func GetOrgEmployeeShiftsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_EMPLOYEE_SHIFTS)
}

func GetOrgSalaryGroupsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_SALARY_GROUPS)
}

func GetOrgPolicyStatsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_COLLECTION_POLICY_STATS)
}

func GetEmpAttendanceRegularizationCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_ATTENDANCE_REGULARIZATION)
}

func GetEmpRewardCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_REWARD)
}

func GetEmpTechInfoCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_TECH_INFO)
}

func GetEmpAssetCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_ASSET)
}

func GetEmpBankCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_BANK)
}

func GetEmpAssetCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_ASSET_COUNTER)
}

func GetEmpLeavesCarryForwardCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(EMPLOYEE_COLLECTION_LEAVES_CARRY_FORWARD)
}

func GetExpenseCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(REIMBURSEMENT_COLLECTION_EXPENSE_COUNTER)
}

func GetExpenseCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(REIMBURSEMENT_COLLECTION_EXPENSE)
}

func GetJobMgrJobsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_JOBS)
}

func GetJobMgrApplicationsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_APPLICATIONS)
}

func GetJobMgrJobCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_JOB_COUNTER)
}

func GetJobMgrApplicationCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_APPLICATION_COUNTER)
}

func GetJobMgrCandidatesCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_CANDIDATES)
}

func GetJobMgrCandidateCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_CANDIDATE_COUNTER)
}

func GetJobMgrAllocationTableCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_ALLOCATION_TABLE)
}

func GetJobMgrDailyStatsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_DAILY_STATS)
}

func GetJobMgrCalculatedStatsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_CALCULATED_STATS)
}

func GetJobMgrErrorRecordsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(JOBMGR_COLLECTION_ERROR_RECORDS)
}

func GetTicketTicketConfigCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(TICKET_COLLECTION_TICKET_CONFIGS)
}

func GetTicketTicketsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(TICKET_COLLECTION_TICKETS)
}

func GetTicketTicketCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(TICKET_COLLECTION_TICKET_COUNTER)
}

func GetTicketTicketLeadsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(TICKET_COLLECTION_TICKET_LEADS)
}

func GetVolunteerCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(VOLUNTEER_COLLECTION_VOLUNTEER_COUNTER)
}

func GetVolunteersCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(VOLUNTEER_COLLECTION_VOLUNTEERS)
}

func GetGSOrgScholarshipRequestCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_SCHOLARSHIP_REQUESTS)
}

func GetGSOrgStudentCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_STUDENT_COUNTER)
}

func GetGSOrgDonationCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_DONATION_COUNTER)
}

func GetGSOrgDonorCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_DONOR_COUNTER)
}

func GetGSOrgSubscriptionCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_SUBSCRIPTION_COUNTER)
}

func GetGSOrgDonationsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_DONATIONS)
}

func GetGSOrgDonationPaymentsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_DONATION_PAYMENTS)
}

func GetGSOrgDonorsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_DONORS)
}

func GetGSOrgSubscriptionCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_SUBSCRIPTIONS)
}

func GetGSOrgPaymentOrdersCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_PAYMENT_ORDERS)
}

func GetGSOrgPaymentPaymentsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(GSORG_COLLECTION_PAYMENT_PAYMENTS)
}

func GetVolunteerRequestsCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(VOLUNTEER_COLLECTION_REQUESTS)
}

func GetWorkerDomainCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(WORKER_COLLECTION_DOMAINS)
}
