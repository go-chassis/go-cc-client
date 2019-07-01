/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package configcenter

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
)

const (
	maxValue = 256
)

//errors
var (
	ErrAppEmpty       = errors.New("app can not be empty")
	ErrServiceTooLong = errors.New("exceeded max value for service name")
)

func GenerateDimension(serviceName, version, appName string) (string, error) {
	if appName != "" {
		serviceName = serviceName + "@" + appName
	} else {
		return "", ErrAppEmpty
	}

	if version != "" {
		serviceName = serviceName + "#" + version
	}

	if len(serviceName) > maxValue {
		return "", ErrServiceTooLong
	}

	dimeExp := `\A([^\$\%\&\+\(/)\[\]\" "\"])*\z`
	dimRegexVar, err := regexp.Compile(dimeExp)
	if err != nil {
		return "", errors.New("not a valid regular expression" + err.Error())
	}

	if !dimRegexVar.Match([]byte(serviceName)) {
		return "", fmt.Errorf("invalid value for dimension info, does not satisfy the regular expression for dimInfo:%s", serviceName)
	}

	return serviceName, nil
}
