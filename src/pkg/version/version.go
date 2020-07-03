/*
 * Copyright © 2019 – 2020 Red Hat Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package version

import "fmt"

// Version is the version of Toolbox
type Version struct {
	Major int
	Minor int
	Micro int
}

// currentVersion holds the information about current build version
var currentVersion = Version{
	Major: 0,
	Minor: 0,
	Micro: 90,
}

// GetVersion returns string with the version of Toolbox
func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", currentVersion.Major, currentVersion.Minor, currentVersion.Micro)
}
