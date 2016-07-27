// +build libvirt.1.2.14

package libvirt

import (
	"testing"
)

func TestDomainListAllInterfaceAddresses(t *testing.T) {
	dom, conn := buildTestQEMUDomain()
	defer func() {
		dom.Free()
		conn.CloseConnection()
	}()
	if err := dom.Create(); err != nil {
		t.Error(err)
		return
	}
	defer dom.Destroy()
	ifaces, err := dom.ListAllInterfaceAddresses(0)
	if err != nil {
		t.Fatal(err)
	}

	if len(ifaces) != 0 {
		t.Fatal("should have 0 interfaces", len(ifaces))
	}
}
