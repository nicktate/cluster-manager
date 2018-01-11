package resources

import (
	"encoding/json"
	"fmt"

	"github.com/containership/cloud-agent/internal/log"
	"github.com/containership/cloud-agent/internal/resources/registry"

	containershipv3 "github.com/containership/cloud-agent/pkg/apis/containership.io/v3"
)

// CsRegistries defines the Containership Cloud CsRegistries resource
type CsRegistries struct {
	cloudResource
	cache []containershipv3.RegistrySpec
}

// NewCsRegistries constructs a new CsRegistries
func NewCsRegistries() *CsRegistries {
	return &CsRegistries{
		cloudResource: cloudResource{
			endpoint: "/organizations/{{.OrganizationID}}/registries",
		},
		cache: make([]containershipv3.RegistrySpec, 0),
	}
}

// UnmarshalToCache take the json returned from containership api
// and writes it to CsRegistries cache
func (rs *CsRegistries) UnmarshalToCache(bytes []byte) error {
	log.Debug("CsRegistries UnmarshallToCache...")

	err := json.Unmarshal(bytes, &rs.cache)
	if err != nil {
		log.Error("Cloud returned registries response:", string(bytes))
	}

	log.Debugf("CsRegistries cache updated: %+v", rs.cache)
	return err
}

// Cache returns CsRegistries cache
func (rs *CsRegistries) Cache() []containershipv3.RegistrySpec {
	return rs.cache
}

// GetAuthToken return the AuthToken Generated by the registry generator
func (rs *CsRegistries) GetAuthToken(spec containershipv3.RegistrySpec) (containershipv3.AuthTokenDef, error) {
	generator := registry.New(spec.Provider, spec.Serveraddress, spec.Credentials)
	return generator.CreateAuthToken()
}

// IsEqual take a Registry Spec and compares it to a Registry to see if they are
// the same, returns an error if the objects are of the inforect type
func (rs *CsRegistries) IsEqual(specObj interface{}, parentSpecObj interface{}) (bool, error) {
	spec, ok := specObj.(containershipv3.RegistrySpec)
	if !ok {
		return false, fmt.Errorf("The object is not of type RegistrySpec")
	}

	user, ok := parentSpecObj.(*containershipv3.Registry)
	if !ok {
		return false, fmt.Errorf("The object is not of type Registry")
	}

	equal := spec.Description == user.Spec.Description &&
		spec.Organization == user.Spec.Organization &&
		spec.Email == user.Spec.Email &&
		spec.Serveraddress == user.Spec.Serveraddress &&
		spec.Provider == user.Spec.Provider &&
		spec.Owner == user.Spec.Owner

	if !equal {
		return false, nil
	}

	for i, k := range spec.Credentials {
		if user.Spec.Credentials[i] != k {
			return false, nil
		}
	}

	return true, nil
}
