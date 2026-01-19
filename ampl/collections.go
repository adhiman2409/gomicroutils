package ampl

import (
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

const ORG_EMPLOYEE_COLLECTION = "org-employee-collection"
const ORG_EMPLOYEE_ID_COUNTER_COLLECTION = "org-employee-id-counter"
const ORG_SALES_EMPLOYEE_DATA_COLLECTION = "org-sales-employee-data"

const CUSTOMER_CUSTOMER_COLLECTION = "customer-customer-collection"
const CUSTOMER_FINANCIAL_DATA_COLLECTION = "customer-financial-data"
const CUSTOMER_ID_COUNTER_COLLECTION = "customer-id-counter"
const AUTH_OEM_SHOWROOM_COLLECTION = "auth-oem-showrooms"
const AUTH_SHOWROOM_ID_COUNTER_COLLECTION = "auth-showroom-id-counter"

func GetOrgEmployeeCollection(client *mongo.Client, domain string) *mongo.Collection {
	domain = strings.Replace(domain, ".", "_", -1)
	return client.Database(domain).Collection(ORG_EMPLOYEE_COLLECTION)
}

func GetOrgEmployeeIdCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_EMPLOYEE_ID_COUNTER_COLLECTION)
}

func GetCustomerCustomerCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(CUSTOMER_CUSTOMER_COLLECTION)
}

func GetCustomerFinancialDataCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(CUSTOMER_FINANCIAL_DATA_COLLECTION)
}

func GetCustomerIdCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(CUSTOMER_ID_COUNTER_COLLECTION)
}

func GetAuthOEMShowroomCollection(client *mongo.Client, domain string) *mongo.Collection {
	domain = strings.Replace(domain, ".", "_", -1)
	return client.Database(domain).Collection(AUTH_OEM_SHOWROOM_COLLECTION)
}

func GetOrgSalesEmployeeDataCollection(client *mongo.Client, domain string) *mongo.Collection {
	db := client.Database(strings.Replace(domain, ".", "_", -1))
	return db.Collection(ORG_SALES_EMPLOYEE_DATA_COLLECTION)
}

func GetAuthShowroomCounterCollection(client *mongo.Client, domain string) *mongo.Collection {
	domain = strings.Replace(domain, ".", "_", -1)
	return client.Database(domain).Collection(AUTH_SHOWROOM_ID_COUNTER_COLLECTION)
}
