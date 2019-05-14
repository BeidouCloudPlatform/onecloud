// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package system_service

import (
	"testing"
)

func TestParseSystemd(t *testing.T) {
	cases := []struct {
		in         string
		name       string
		wantLoaded bool
		wantActive bool
	}{
		{`openvswitch.service - LSB: Open vSwitch switch
   Loaded: loaded (/etc/rc.d/init.d/openvswitch; bad; vendor preset: disabled)
   Active: active (running) since Wed 2019-04-10 15:11:23 CST; 4 days ago
     Docs: man:systemd-sysv-generator(8)
  Process: 3686 ExecStart=/etc/rc.d/init.d/openvswitch start (code=exited, status=0/SUCCESS)
   CGroup: /system.slice/openvswitch.service
           ├─3722 ovsdb-server: monitoring pid 3723 (healthy)
           ├─3723 ovsdb-server /etc/openvswitch/conf.db -vconsole:emer -vsysl...
           ├─3774 ovs-vswitchd: monitoring pid 3775 (healthy)
           └─3775 ovs-vswitchd unix:/var/run/openvswitch/db.sock -vconsole:em...`, "openvswitch", true, true},
		{"Unit yunion-host.service could not be found.", "yunion-host", false, false},
	}
	for _, c := range cases {
		status := parseSystemdStatus(c.in, c.name)
		if status.Loaded != c.wantLoaded || status.Active != c.wantActive {
			t.Errorf("parseSystemdStatus %s Loaded: want %v got %v Active: want %v got %v", c.name, c.wantLoaded, status.Loaded, c.wantActive, status.Active)
		}
	}
}
