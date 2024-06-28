package main

import (
	"fmt"
	"go.seq.tf/oiia/selfservice-data-generator/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	//dsn := "postgres:p@tcp(127.0.0.1:5432)/db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=localhost user=postgres password=p dbname=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Generate new topics
	for i := 1; i < 900; i++ {
		newTopic := model.GenerateNewTopic("cloudengineering-xxx")
		db.Exec(`INSERT into "KafkaTopic" ("Id", "Name", "Description", "CreatedAt", "CreatedBy","Status", "ModifiedAt", "ModifiedBy", "Partitions", "KafkaClusterId", "CapabilityId", "Retention") values (?, ?, ?, ?, ?, ?, ?, ?, ? ,?, ?, ?)`, newTopic.Id, newTopic.Name, newTopic.Description, newTopic.CreatedAt, newTopic.CreatedBy, newTopic.Status, newTopic.ModifiedAt, newTopic.ModifiedBy, newTopic.Partitions, newTopic.KafkaClusterId, newTopic.CapabilityId, newTopic.Retention)
		fmt.Println(newTopic)
	}

	// Generate new members
	//for i := 1; i < 1000; i++ {
	//	newMember := model.GenerateNewMember()
	//	db.Exec(`INSERT into "Member" ("Id", "Email", "DisplayName") values (?, ?, ?)`, newMember.Id, newMember.Email, newMember.DisplayName)
	//	fmt.Println(newMember)
	//
	//	// Generate new memberships
	//	newMembership := model.GenerateNewMembership(newMember.Id, "cloudengineering-xxx")
	//	db.Exec(`INSERT into "Membership" ("Id", "UserId", "CapabilityId", "CreatedAt") values (?, ?, ?, ?)`, newMembership.Id, newMembership.UserId, newMembership.CapabilityId, newMembership.CreatedAt)
	//
	//	fmt.Println(newMembership)
	//}

	// Generate new Capabilities
	//for i := 1; i < 4000; i++ {
	//	newCap := model.GenerateNewCapability()
	//	db.Exec(`INSERT into "Capability" ("Id","Name","Description","CreatedAt","CreatedBy","Status","ModifiedAt","ModifiedBy","JsonMetadata","JsonMetadataSchemaVersion") values (?, ?, ?, ?, ?, ?, ?, ?, ? ,?)`, newCap.Id, newCap.Name, newCap.Description, newCap.CreatedAt, newCap.CreatedBy, newCap.Status, newCap.ModifiedAt, newCap.ModifiedBy, newCap.JsonMetadata, newCap.JsonMetadataSchemaVersion)
	//	fmt.Println(newCap)
	//}

	// Generate new ECR repos

	//for i := 1; i < 100; i++ {
	//	newEcr := model.GenerateNewEcr()
	//	db.Exec(`INSERT into "ECRRepository" ("Id","Name","Description","CreatedBy","RequestedAt") values (?, ?, ?, ?, ?)`, newEcr.Id, newEcr.Name, newEcr.Description, newEcr.CreatedBy, newEcr.RequestedAt)
	//	fmt.Println(newEcr)
	//}

}
