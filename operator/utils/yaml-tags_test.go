package utils

import (
	"fmt"
	"testing"

	akmv1a1 "gitlab.com/GeorgeRaven/authentik-manager/operator/api/v1alpha1"
	"gopkg.in/yaml.v3"
)

type Taggable struct {
	value interface{}
	Tag   string
}

// TestYamlTags ensures blueprints marshal to and unmarshal from the same result.
func TestYamlTags(t *testing.T) {
	// get example blueprint raw marshalled data
	yamlData := getExampleBlueprint()
	fmt.Println(yamlData)
	// unmarshal
	blueprint := akmv1a1.BP{}
	err := yaml.Unmarshal([]byte(yamlData), &blueprint)
	if err != nil {
		t.Fatal(err)
	}
}

// getExampleBlueprint deterministically generates an example blueprint to use.
func getExampleBlueprint() string {
	return `
version: 1
metadata:
  name: Default - Authentication flow
entries:
- model: authentik_blueprints.metaapplyblueprint
  attrs:
    identifiers:
      name: Default - Password change flow
    required: false
- attrs:
    designation: authentication
    name: Welcome to authentik!
    title: Welcome to authentik!
    authentication: none
  identifiers:
    slug: default-authentication-flow
  model: authentik_flows.flow
  id: flow
- attrs:
    backends:
    - authentik.core.auth.InbuiltBackend
    - authentik.sources.ldap.auth.LDAPBackend
    - authentik.core.auth.TokenBackend
    configure_flow: !Find [authentik_flows.flow, [slug, default-password-change]]
  identifiers:
    name: default-authentication-password
  id: default-authentication-password
  model: authentik_stages_password.passwordstage
- identifiers:
    name: default-authentication-mfa-validation
  id: default-authentication-mfa-validation
  model: authentik_stages_authenticator_validate.authenticatorvalidatestage
- attrs:
    user_fields:
    - email
    - username
  identifiers:
    name: default-authentication-identification
  id: default-authentication-identification
  model: authentik_stages_identification.identificationstage
- identifiers:
    name: default-authentication-login
  id: default-authentication-login
  model: authentik_stages_user_login.userloginstage
- identifiers:
    order: 10
    stage: !KeyOf default-authentication-identification
    target: !KeyOf flow
  model: authentik_flows.flowstagebinding
- identifiers:
    order: 20
    stage: !KeyOf default-authentication-password
    target: !KeyOf flow
  model: authentik_flows.flowstagebinding
- identifiers:
    order: 30
    stage: !KeyOf default-authentication-mfa-validation
    target: !KeyOf flow
  model: authentik_flows.flowstagebinding
- identifiers:
    order: 100
    stage: !KeyOf default-authentication-login
    target: !KeyOf flow
  model: authentik_flows.flowstagebinding
`
}
