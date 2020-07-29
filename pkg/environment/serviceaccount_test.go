package environment

import (
	"os"
	"testing"
)

func TestCreateServiceAccountFiles(t *testing.T) {
	svc := ServiceAccount{
		Name:      "testName",
		Namespace: "testNamespace",
	}

	fileName := "05-serviceaccount.yaml"

	err := createFile(&svc)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	fileContainsString(t, fileName, svc.Name)

	os.Remove(fileName)
}