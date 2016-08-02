/* Copyright 2015 Palantir Technologies, Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package launchlib

import (
	"os"
	"testing"
)

func TestGetJavaHome(t *testing.T) {
	originalJavaHome := os.Getenv("JAVA_HOME")
	setEnvOrFail("JAVA_HOME", "foo")

	javaHome := getJavaHome("")
	if javaHome != "foo" {
		t.Error("Expected JAVA_HOME='foo', found", javaHome)
	}
	javaHome = getJavaHome("explicit javahome")
	if javaHome != "explicit javahome" {
		t.Error("Expected JAVA_HOME='explicit javahome', found", javaHome)
	}

	setEnvOrFail("JAVA_HOME", originalJavaHome)
}

func setEnvOrFail(key string, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		panic("Failed to set env var: " + key)
	}
}
