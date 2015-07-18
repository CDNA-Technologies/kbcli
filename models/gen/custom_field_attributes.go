/*
 * Copyright 2014 Groupon, Inc.
 * Copyright 2014 The Billing Project, LLC
 *
 * The Billing Project licenses this file to you under the Apache License, version 2.0
 * (the "License"); you may not use this file except in compliance with the
 * License.  You may obtain a copy of the License at:
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */


/*
 *                       DO NOT EDIT!!!
 *    File automatically generated by killbill-java-parser (git@github.com:killbill/killbill-java-parser.git)
 */


package gen

import "encoding/json"


type CustomFieldAttributes struct {
  CustomFieldId string `json:"customFieldId"`
  ObjectId string `json:"objectId"`
  ObjectType string `json:"objectType"`
  Name string `json:"name"`
  Value string `json:"value"`
  AuditLogs []AuditLogAttributes `json:"auditLogs"`
}


func (data * CustomFieldAttributes) FromJson(raw []byte) error {
  return json.Unmarshal(raw, data)
}
