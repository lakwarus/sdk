/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package runtime

import (
	"path/filepath"

	"github.com/cellery-io/sdk/components/cli/pkg/kubectl"
)

func addApim() error {
	for _, v := range buildApimYamlPaths() {
		err := kubectl.ApplyFileWithNamespace(v, "cellery-system")
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteApim() error {
	for _, v := range buildApimYamlPaths() {
		err := kubectl.DeleteFileWithNamespace(v, "cellery-system")
		if err != nil {
			return err
		}
	}
	return nil
}

func buildApimYamlPaths() []string {
	base := buildArtifactsPath(ApiManager)
	return []string{
		filepath.Join(base, "global-apim-volatile.yaml"),
	}
}
