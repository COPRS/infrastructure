#
# Copyright 2019- CS Group
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

require "fluent/plugin/filter"
require "date"
require "json"

def is_formatted(record)
    begin
        json_record = JSON.parse(record["log"])
        if ! json_record.has_key?("header") || ! json_record["header"].has_key?("type") || ! json_record["header"].has_key?("timestamp") || ! json_record["header"].has_key?("level") || ! json_record.has_key?("message") 
            return false
        end
    rescue Exception => e  
        return false
    end 
    return true
end

class FilenamePropertiesFilter < Fluent::Plugin::Filter
    Fluent::Plugin.register_filter("icd_formatter", self)

    def filter(tag, time, record)
        formated_record = nil
        #Do not format logs who already match the ICD !!
        if ! is_formatted(record)
            level = "INFO"
            if record["stream"] == "stderr"
                level = "ERROR"
            end
            timestamp = Time.at(time).to_datetime.strftime("%FT%T.%6NZ")
            if record.has_key?("time")
                timestamp = record["time"]
            end    
            formated_record = { "header" => {"type" => "LOG", "timestamp" => timestamp, "level" => level}, "message" => {"content" => record["log"]} }
            if record.has_key?("kubernetes")
                formated_record["kubernetes"] = record["kubernetes"]
            elsif record.has_key?("hostname") && record.has_key?("processus")
                formated_record["kubernetes"] = {"host" => record["hostname"], "processus" => record["processus"]}
            end
        else 
            formated_record = JSON.parse(record["log"])
            formated_record["kubernetes"] = record["kubernetes"]
        end
        return formated_record
    end
end
