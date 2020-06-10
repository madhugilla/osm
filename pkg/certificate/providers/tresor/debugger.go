package tresor

import "github.com/open-service-mesh/osm/pkg/certificate"

// ListIssuedCertificates implements CertificateDebugger interface and returns the list of issued certificates.
func (cm *CertManager) ListIssuedCertificates() []certificate.Certificater {
	var certs []certificate.Certificater
	cm.cacheLock.Lock()
	defer cm.cacheLock.Unlock()
	for _, cert := range *cm.cache {
		certs = append(certs, cert)
	}
	return certs
}
